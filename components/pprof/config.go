package pprof

import "github.com/Compogo/compogo/configurator"

const (
	// UseProfileFieldName is the command-line flag to enable pprof endpoints.
	UseProfileFieldName = "trace.pprof"

	// EndpointFieldName is the command-line flag for pprof endpoint path.
	EndpointFieldName = "server.http.routes.pprof"

	// UseProfileDefault disables pprof by default.
	UseProfileDefault = false

	// EndpointDefault is the default path for pprof endpoints.
	EndpointDefault = "/debug"
)

// Config holds the pprof configuration.
// It can be populated from command-line flags or config files.
type Config struct {
	UseProfile bool
	Endpoint   string
}

// NewConfig creates a new Config instance with default values.
func NewConfig() *Config {
	return &Config{}
}

// Configuration applies configuration values to the Config struct.
// It reads from the provided configurator and sets defaults if values are not present.
func Configuration(config *Config, configurator configurator.Configurator) *Config {
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
