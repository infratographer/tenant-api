// Package config defines the application config used through tenant-api.
package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.infratographer.com/x/crdbx"
	"go.infratographer.com/x/echojwtx"
	"go.infratographer.com/x/echox"
	"go.infratographer.com/x/events"
	"go.infratographer.com/x/loggingx"
	"go.infratographer.com/x/otelx"
	"go.infratographer.com/x/viperx"
)

const defaultRuntimeSocket = "unix:///tmp/runtime.sock"

// AppConfig contains the application configuration structure.
var AppConfig struct {
	CRDB          crdbx.Config
	Logging       loggingx.Config
	Events        events.Config
	Server        echox.Config
	OIDC          echojwtx.AuthConfig
	Tracing       otelx.Config
	RuntimeSocket string
}

// MustViperFlags returns the cobra flags and wires them up with viper to prevent code duplication.
func MustViperFlags(v *viper.Viper, flags *pflag.FlagSet) {
	flags.String("runtime-socket", "", "set the IAM Runtime socket. default: "+defaultRuntimeSocket)
	viperx.MustBindFlag(v, "runtimeSocket", flags.Lookup("runtime-socket"))
	v.SetDefault("runtimeSocket", defaultRuntimeSocket)
}
