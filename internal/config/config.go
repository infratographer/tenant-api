// Package config defines the application config used through tenant-api.
package config

import (
	"go.infratographer.com/x/crdbx"
	"go.infratographer.com/x/echox"
	"go.infratographer.com/x/events"
	"go.infratographer.com/x/loggingx"
	"go.infratographer.com/x/otelx"
)

// AppConfig contains the application configuration structure.
var AppConfig struct {
	CRDB    crdbx.Config
	Logging loggingx.Config
	Events  EventsConfig
	Server  echox.Config
	Tracing otelx.Config
}

type EventsConfig struct {
	Publisher events.Publisher
}
