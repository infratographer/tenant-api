package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/metal-toolbox/iam-runtime-contrib/iamruntime"
	"github.com/metal-toolbox/iam-runtime-contrib/middleware/echo/iamruntimemiddleware"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"go.infratographer.com/x/crdbx"
	"go.infratographer.com/x/echox"
	"go.infratographer.com/x/events"
	"go.infratographer.com/x/otelx"
	"go.infratographer.com/x/versionx"

	"go.infratographer.com/tenant-api/internal/config"
	ent "go.infratographer.com/tenant-api/internal/ent/generated"
	"go.infratographer.com/tenant-api/internal/ent/generated/eventhooks"
	"go.infratographer.com/tenant-api/internal/graphapi"
)

const (
	// APIDefaultListen defines the default listening address for the tenant-api.
	APIDefaultListen = ":7902"

	shutdownTimeout = 10 * time.Second
)

var (
	enablePlayground bool
	serveDevMode     bool
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start Tenant API",
	Run: func(cmd *cobra.Command, _ []string) {
		serve(cmd.Context())
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	echox.MustViperFlags(viper.GetViper(), serveCmd.Flags(), APIDefaultListen)
	events.MustViperFlags(viper.GetViper(), serveCmd.Flags(), appName)
	config.MustViperFlags(viper.GetViper(), serveCmd.Flags())

	// only available as a CLI arg because it shouldn't be something that could accidentially end up in a config file or env var
	serveCmd.Flags().BoolVar(&serveDevMode, "dev", false, "dev mode: enables playground, disables all auth checks, sets CORS to allow all, pretty logging, etc.")
	serveCmd.Flags().BoolVar(&enablePlayground, "playground", false, "enable the graph playground")
}

func serve(ctx context.Context) {
	iamMiddlewareConfig := iamruntimemiddleware.NewConfig()

	if serveDevMode {
		enablePlayground = true
		config.AppConfig.Logging.Debug = true
		config.AppConfig.Logging.Pretty = true
		config.AppConfig.Server.WithMiddleware(middleware.CORS())

		iamMiddlewareConfig.Skipper = func(_ echo.Context) bool { return true }
	}

	events, err := events.NewConnection(config.AppConfig.Events, events.WithLogger(logger))
	if err != nil {
		logger.Fatal("unable to initialize events", zap.Error(err))
	}

	err = otelx.InitTracer(config.AppConfig.Tracing, appName, logger)
	if err != nil {
		logger.Fatal("unable to initialize tracing system", zap.Error(err))
	}

	db, err := crdbx.NewDB(config.AppConfig.CRDB, config.AppConfig.Tracing.Enabled)
	if err != nil {
		logger.Fatal("unable to initialize crdb client", zap.Error(err))
	}

	defer db.Close() //nolint:errcheck

	entDB := entsql.OpenDB(dialect.Postgres, db)

	cOpts := []ent.Option{ent.Driver(entDB), ent.EventsPublisher(events)}

	if config.AppConfig.Logging.Debug {
		cOpts = append(cOpts,
			ent.Log(logger.Named("ent").Debugln),
			ent.Debug(),
		)
	}

	client := ent.NewClient(cOpts...)
	defer client.Close() //nolint:errcheck

	eventhooks.EventHooks(client)

	srv, err := echox.NewServer(logger.Desugar(), echox.ConfigFromViper(viper.GetViper()), versionx.BuildDetails())
	if err != nil {
		logger.Fatal("failed to initialize new server", zap.Error(err))
	}

	var middleware []echo.MiddlewareFunc

	runtime, err := iamruntime.NewClient(config.AppConfig.RuntimeSocket)
	if err != nil {
		logger.Fatal("failed to initialize IAM runtime", zap.Error(err))
	}

	iamMiddleware, err := iamMiddlewareConfig.WithRuntime(runtime).ToMiddleware()
	if err != nil {
		logger.Fatal("failed to initialize IAM runtime middleware", zap.Error(err))
	}

	middleware = append(middleware, iamMiddleware)

	r := graphapi.NewResolver(client, logger.Named("resolvers"))
	handler := r.Handler(enablePlayground, middleware)

	srv.AddHandler(handler)

	// TODO: we should have a database check
	// srv.AddReadinessCheck("database", r.DatabaseCheck)

	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.Run(); err != nil {
			logger.Fatal("failed to run server", zap.Error(err))
		}

		cancel()
	}()

	select {
	case <-sig:
		logger.Info("signal caught, shutting down")

		ctx, cancel = context.WithTimeout(ctx, shutdownTimeout)
	case <-ctx.Done():
		logger.Info("context done, shutting down")

		ctx, cancel = context.WithTimeout(context.Background(), shutdownTimeout)
	}

	defer cancel()

	if err := events.Shutdown(ctx); err != nil {
		logger.Fatalw("failed to shutdown events gracefully", "error", err)
	}
}
