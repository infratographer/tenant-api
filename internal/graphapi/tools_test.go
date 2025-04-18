package graphapi_test

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/Yamashou/gqlgenc/clientv2"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vektah/gqlparser/v2/ast"
	"go.uber.org/zap"

	"go.infratographer.com/x/echojwtx"
	"go.infratographer.com/x/events"
	"go.infratographer.com/x/goosex"
	"go.infratographer.com/x/testing/containersx"
	"go.infratographer.com/x/testing/eventtools"

	"go.infratographer.com/tenant-api/db"
	ent "go.infratographer.com/tenant-api/internal/ent/generated"
	"go.infratographer.com/tenant-api/internal/ent/generated/eventhooks"
	"go.infratographer.com/tenant-api/internal/graphapi"
	"go.infratographer.com/tenant-api/internal/testclient"
)

var TestDBURI = os.Getenv("TENANTAPI_TESTDB_URI")

var testTools struct {
	entClient   *ent.Client
	dbContainer *containersx.DBContainer

	pubsubEntClient *ent.Client
	eventsConfig    events.Config
}

func TestMain(m *testing.M) {
	// setup the database if needed
	setupDB()
	// run the tests
	code := m.Run()
	// teardown the database
	teardownDB()
	// return the test response code
	os.Exit(code)
}

func parseDBURI(ctx context.Context) (string, string, *containersx.DBContainer) {
	switch {
	// if you don't pass in a database we default to an in memory sqlite
	case TestDBURI == "":
		return dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1", nil
	case strings.HasPrefix(TestDBURI, "sqlite://"):
		return dialect.SQLite, strings.TrimPrefix(TestDBURI, "sqlite://"), nil
	case strings.HasPrefix(TestDBURI, "postgres://"), strings.HasPrefix(TestDBURI, "postgresql://"):
		return dialect.Postgres, TestDBURI, nil
	case strings.HasPrefix(TestDBURI, "docker://"):
		dbImage := strings.TrimPrefix(TestDBURI, "docker://")

		switch {
		case strings.HasPrefix(dbImage, "cockroach"), strings.HasPrefix(dbImage, "cockroachdb"), strings.HasPrefix(dbImage, "crdb"):
			cntr, err := containersx.NewCockroachDB(ctx, dbImage)
			if err != nil {
				log.Panicf("error starting db test container: %s", err.Error())
			}

			return dialect.Postgres, cntr.URI, cntr
		case strings.HasPrefix(dbImage, "postgres"):
			cntr, err := containersx.NewPostgresDB(ctx, dbImage)
			if err != nil {
				log.Panicf("error starting db test container: %s", err.Error())
			}

			return dialect.Postgres, cntr.URI, cntr
		default:
			panic("invalid testcontainer URI, uri: " + TestDBURI)
		}

	default:
		panic("invalid DB URI, uri: " + TestDBURI)
	}
}

func setupDB() {
	// don't setup the datastore if we already have one
	if testTools.entClient != nil {
		return
	}

	ctx := context.Background()

	nats, err := eventtools.NewNatsServer()
	if err != nil {
		log.Panicf("error creating nats server: %s", err.Error())
	}

	testTools.eventsConfig = nats.Config
	testTools.eventsConfig.NATS.Source = "tenant-api-test"

	dia, uri, cntr := parseDBURI(ctx)

	conn, err := events.NewConnection(testTools.eventsConfig)
	if err != nil {
		log.Panicf("error creating pubsubx publisher: %s", err.Error())
	}

	c, err := ent.Open(dia, uri, ent.Debug(), ent.EventsPublisher(conn))
	if err != nil {
		if err := cntr.Container.Terminate(ctx); err != nil {
			log.Printf("error terminating test db container: %s", err.Error())
		}

		log.Panicf("error opening connection to database: %s", err)
	}

	switch dia {
	case dialect.SQLite:
		// Run automatic migrations for SQLite
		if err := c.Schema.Create(ctx); err != nil {
			log.Panicf("error creating db schema: %s", err.Error())
		}
	case dialect.Postgres:
		log.Println("Running database migrations")
		goosex.MigrateUpContext(context.Background(), uri, db.Migrations)
	}

	testTools.dbContainer = cntr
	testTools.entClient = c
	testTools.pubsubEntClient = c
	eventhooks.EventHooks(testTools.pubsubEntClient)
}

func teardownDB() {
	ctx := context.Background()

	if testTools.entClient != nil {
		if err := testTools.entClient.Close(); err != nil {
			log.Panicf("teardown failed to close database connection: %s", err.Error())
		}
	}

	if testTools.dbContainer != nil {
		if err := testTools.dbContainer.Container.Terminate(ctx); err != nil {
			log.Panicf("teardown failed to terminate test db container: %s", err.Error())
		}
	}
}

func graphTestClient(entClient *ent.Client) testclient.TestClient {
	return testclient.NewClient(&http.Client{Transport: localRoundTripper{handler: newDefaultServer(
		graphapi.NewExecutableSchema(
			graphapi.Config{Resolvers: graphapi.NewResolver(entClient, zap.NewNop().Sugar())},
		))}}, "graph", &clientv2.Options{})
}

// localRoundTripper is an http.RoundTripper that executes HTTP transactions
// by using handler directly, instead of going over an HTTP connection.
type localRoundTripper struct {
	handler http.Handler
}

func (l localRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	w := httptest.NewRecorder()
	// set the actor to "testing-roundtrip-actor"
	req := r.WithContext(context.WithValue(r.Context(), echojwtx.ActorCtxKey, "testing-roundtrip-actor"))
	l.handler.ServeHTTP(w, req)

	return w.Result(), nil
}

func newDefaultServer(es graphql.ExecutableSchema) *handler.Server {
	srv := handler.New(es)

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	return srv
}
