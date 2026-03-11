package health_check

import "github.com/Compogo/compogo/configurator"

const (
	// EndpointFieldName is the command-line flag for health check endpoint path.
	EndpointFieldName = "server.http.routes.health_check"

	// EndpointDefault is the default path for health checks.
	EndpointDefault = "/health-check"
)

// Config holds the health check endpoint configuration.
type Config struct {
	Endpoint string
}

// NewConfig creates a new Config instance with default values.
func NewConfig() *Config {
	return &Config{}
}

// Configuration applies configuration values to the Config struct.
// It reads from the provided configurator and sets defaults if values are not present.
func Configuration(config *Config, configurator configurator.Configurator) *Config {
	if config.Endpoint == "" || config.Endpoint == EndpointDefault {
		configurator.SetDefault(EndpointFieldName, EndpointDefault)
		config.Endpoint = configurator.GetString(EndpointFieldName)
	}

	return config
}
