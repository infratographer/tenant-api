// Package config defines the application config used through tenant-api.
package config

import (
	"go.infratographer.com/x/crdbx"
	"go.infratographer.com/x/echojwtx"
	"go.infratographer.com/x/echox"
	"go.infratographer.com/x/events"
	"go.infratographer.com/x/loggingx"
	"go.infratographer.com/x/otelx"

	"go.infratographer.com/permissions-api/pkg/permissions"
)

// AppConfig contains the application configuration structure.
var AppConfig struct {
	CRDB        crdbx.Config
	Logging     loggingx.Config
	Events      events.Config
	Server      echox.Config
	OIDC        echojwtx.AuthConfig
	Tracing     otelx.Config
	Permissions permissions.Config
}
