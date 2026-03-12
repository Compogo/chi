package health_check

import (
	"github.com/Compogo/chi"
	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
	"github.com/Compogo/compogo/flag"
	"github.com/Compogo/compogo/logger"
	"github.com/Compogo/http"
	"github.com/go-chi/chi/v5/middleware"
)

// Component adds a health check endpoint to the HTTP router using chi's heartbeat middleware.
// It depends on chi.Component for the router.
//
// Usage:
//
//	compogo.WithComponents(
//	    http.Component,
//	    chi.Component,
//	    health_check.Component,
//	)
//
// Health check responds with 200 OK at: /health-check (configurable)
var Component = &component.Component{
	Dependencies: component.Components{
		chi.Component,
	},
	Init: component.StepFunc(func(container container.Container) error {
		return container.Provide(NewConfig)
	}),
	BindFlags: component.BindFlags(func(flagSet flag.FlagSet, container container.Container) error {
		return container.Invoke(func(config *Config) {
			flagSet.StringVar(&config.Endpoint, EndpointFieldName, EndpointDefault, "path for liveness test endpoint")
		})
	}),
	Configuration: component.StepFunc(func(container container.Container) error {
		return container.Invoke(Configuration)
	}),
	PreExecute: component.StepFunc(func(container container.Container) error {
		return container.Invoke(func(config *Config, r http.Router, logger logger.Logger) {
			logger.Infof("[chi.router] add health endpoint - '%s'", config.Endpoint)
			r.Use(http.MiddlewareFunc(middleware.Heartbeat(config.Endpoint)))
		})
	}),
}
