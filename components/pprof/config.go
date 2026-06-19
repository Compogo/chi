package pprof

import (
	"github.com/Compogo/compogo"
)

const (
	// UseProfileFieldName использовать ли pprof
	UseProfileFieldName = "trace.pprof"

	// EndpointFieldName путь для pprof
	EndpointFieldName = "server.http.routes.pprof"
)

var (
	// UseProfileDefault включить pprof
	UseProfileDefault = false

	// EndpointDefault путь для pprof
	EndpointDefault = "/debug"
)

// Config содержит конфигурацию pprof эндпоинтов.
type Config struct {
	UseProfile bool
	Endpoint   string
}

// NewConfig создаёт новую конфигурацию.
func NewConfig() *Config {
	return &Config{}
}

// Configuration загружает конфигурацию из Configurator.
func Configuration(config *Config, configurator compogo.Configurator) *Config {
	if config.UseProfile == UseProfileDefault {
		configurator.SetDefault(UseProfileFieldName, UseProfileDefault)
		config.UseProfile = configurator.GetBool(UseProfileFieldName)
	}

	if config.Endpoint == "" || config.Endpoint == EndpointDefault {
		configurator.SetDefault(EndpointFieldName, EndpointDefault)
		config.Endpoint = configurator.GetString(EndpointFieldName)
	}

	return config
}
