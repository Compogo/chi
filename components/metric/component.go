package metric

import (
	"github.com/Compogo/chi"
	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
	"github.com/Compogo/compogo/flag"
	"github.com/Compogo/compogo/logger"
	"github.com/Compogo/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Component adds a Prometheus metrics endpoint to the HTTP router.
// It depends on chi.Component for the router.
//
// Usage:
//
//	compogo.WithComponents(
//	    http.Component,
//	    chi.Component,
//	    metric.Component,
//	)
//
// Metrics are available at: /metrics (configurable)
var Component = &component.Component{
	Dependencies: component.Components{
		chi.Component,
	},
	Init: component.StepFunc(func(container container.Container) error {
		return container.Provide(NewConfig)
	}),
	BindFlags: component.BindFlags(func(flagSet flag.FlagSet, container container.Container) error {
		return container.Invoke(func(config *Config) {
			flagSet.StringVar(&config.Endpoint, EndpointFieldName, EndpointDefault, "path for metrics endpoint")
		})
	}),
	Configuration: component.StepFunc(func(container container.Container) error {
		return container.Invoke(Configuration)
	}),
	PreExecute: component.StepFunc(func(container container.Container) error {
		return container.Invoke(func(config *Config, r http.Router, logger logger.Logger) {
			logger.Infof("[chi.router] add metrics endpoint - '%s'", config.Endpoint)
			r.Mount(config.Endpoint, promhttp.Handler())
		})
	}),
}
