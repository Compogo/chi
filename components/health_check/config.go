package health_check

import (
	"github.com/Compogo/compogo"
)

// EndpointFieldName имя поля в конфигурации
const EndpointFieldName = "server.http.routes.health_check"

// EndpointDefault путь по умолчанию
var EndpointDefault = "/health-check"

// Config содержит конфигурацию health-check эндпоинта.
type Config struct {
	Endpoint string
}

// NewConfig создаёт новую конфигурацию.
func NewConfig() *Config {
	return &Config{}
}

// Configuration загружает конфигурацию из Configurator.
func Configuration(config *Config, configurator compogo.Configurator) *Config {
	if config.Endpoint == "" || config.Endpoint == EndpointDefault {
		configurator.SetDefault(EndpointFieldName, EndpointDefault)
		config.Endpoint = configurator.GetString(EndpointFieldName)
	}

	return config
}
