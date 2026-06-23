package health_check

import (
	"github.com/Compogo/chi"
	"github.com/Compogo/compogo"
	"github.com/Compogo/compogo/flag"
	httpServer "github.com/Compogo/http_server"
	"github.com/go-chi/chi/v5/middleware"
)

// Component — компонент health-check эндпоинта для chi-роутера.
// Добавляет эндпоинт для проверки живости сервиса (liveness probe).
var Component = compogo.Component{
	Dependencies: compogo.Components{
		&chi.Component,
	},
	Init: compogo.StepFunc(func(container compogo.Container) error {
		return container.Provide(NewConfig)
	}),
	BindFlags: compogo.BindFlags(func(flagSet flag.FlagSet, container compogo.Container) error {
		return container.Invoke(func(config *Config) {
			flagSet.StringVar(&config.Endpoint, EndpointFieldName, EndpointDefault, "path for liveness test endpoint")
		})
	}),
	Configuration: compogo.StepFunc(func(container compogo.Container) (err error) {
		return container.Invoke(Configuration)
	}),
	PreExecute: compogo.StepFunc(func(container compogo.Container) error {
		return container.Invoke(func(config *Config, r httpServer.Router, logger compogo.Logger) {
			logger.GetLogger("http").
				GetLogger("server").
				GetLogger("router").
				GetLogger("chi").
				Infof("add health endpoint - '%s'", config.Endpoint)

			r.Use(httpServer.MiddlewareFunc(middleware.Heartbeat(config.Endpoint)))
		})
	}),
}
