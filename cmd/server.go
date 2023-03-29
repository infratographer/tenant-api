/*
Copyright © 2022 The Infratographer Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/metal-toolbox/auditevent/helpers"
	"github.com/metal-toolbox/auditevent/middleware/echoaudit"
	nats "github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.infratographer.com/tenant-api/internal/config"
	"go.infratographer.com/tenant-api/internal/pubsub"
	"go.infratographer.com/tenant-api/pkg/api/v1"
	"go.infratographer.com/tenant-api/pkg/echox"
	"go.infratographer.com/tenant-api/pkg/jwtauth"
	"go.infratographer.com/x/crdbx"
	"go.infratographer.com/x/otelx"
	"go.infratographer.com/x/viperx"
	"go.uber.org/zap"
)

var (
	APIDefaultListen = ":7601"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start Tenant API",
	Run: func(cmd *cobra.Command, args []string) {
		serve(cmd.Context())
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	echox.MustViperFlags(viper.GetViper(), serveCmd.Flags(), APIDefaultListen)
	jwtauth.MustViperFlags(viper.GetViper(), serveCmd.Flags())

	// audit log path
	serveCmd.Flags().String("audit-log-path", "/app-audit/audit.log", "Path to the audit log file")
	viperx.MustBindFlag(viper.GetViper(), "audit.log.path", serveCmd.Flags().Lookup("audit-log-path"))

}

func serve(ctx context.Context) {
	err := otelx.InitTracer(config.AppConfig.Tracing, appName, logger.Sugar())
	if err != nil {
		logger.Fatal("unable to initialize tracing system", zap.Error(err))
	}

	db, err := crdbx.NewDB(config.AppConfig.CRDB, config.AppConfig.Tracing.Enabled)
	if err != nil {
		logger.Fatal("unable to initialize crdb client", zap.Error(err))
	}

	js, natsClose, err := newJetstreamConnection()
	if err != nil {
		logger.Fatal("failed to create NATS jetstream connection", zap.Error(err))
	}

	defer natsClose()

	var auth *jwtauth.Auth

	if jwksurl := viper.GetString("jwks.url"); jwksurl != "" {
		auth, err = jwtauth.NewAuth(jwtauth.AuthConfig{
			JWKSURI: jwksurl,
		})
		if err != nil {
			logger.Fatal("failed to initialize jwt authentication", zap.Error(err))
		}
	}

	auditMiddleware, auditCloseFn, err := newAuditMiddleware(ctx)
	if err != nil {
		logger.Fatal("Failed to initialize audit middleware", zap.Error(err))
	}

	e := echox.NewServer()

	if auditMiddleware != nil {
		defer auditCloseFn()

		e.Use(auditMiddleware.Audit())
	}

	r := api.NewRouter(
		db,
		logger,
		auth,
		pubsub.NewClient(
			pubsub.WithJetreamContext(js),
			pubsub.WithLogger(logger),
			pubsub.WithStreamName(viper.GetString("nats.stream-name")),
			pubsub.WithSubjectPrefix(viper.GetString("nats.subject-prefix")),
		),
	)

	r.Routes(e)

	e.Logger.Fatal(e.Start(config.AppConfig.Server.Listen))
}

func newJetstreamConnection() (nats.JetStreamContext, func(), error) {
	opts := []nats.Option{nats.Name(appName)}

	if viper.GetBool("debug") {
		logger.Debug("enabling development settings")

		opts = append(opts, nats.Token(viper.GetString("nats.token")))
	} else {
		opts = append(opts, nats.UserCredentials(viper.GetString("nats.creds-file")))
	}

	nc, err := nats.Connect(viper.GetString("nats.url"), opts...)
	if err != nil {
		return nil, nil, err
	}

	js, err := nc.JetStream()
	if err != nil {
		return nil, nil, err
	}

	return js, nc.Close, nil
}

func newAuditMiddleware(ctx context.Context) (*echoaudit.Middleware, func() error, error) {
	auditFile := viper.GetString("audit.log.path")
	if auditFile == "" {
		logger.Warn("audit log path not provied, logging disabled.")

		return nil, nil, nil
	}

	auditLogPath := viper.GetViper().GetString("audit.log.path")
	fd, err := helpers.OpenAuditLogFileUntilSuccessWithContext(ctx, auditLogPath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open audit log file: %w", err)
	}

	return echoaudit.NewJSONMiddleware("tenant-api", fd), fd.Close, nil
}
