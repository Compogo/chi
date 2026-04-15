package pprof

import (
	"github.com/Compogo/chi"
	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
	"github.com/Compogo/compogo/flag"
	"github.com/Compogo/compogo/logger"
	"github.com/Compogo/http"
	"github.com/go-chi/chi/v5/middleware"
)

// Component adds pprof debugging endpoints to the HTTP router when enabled.
// It depends on chi.Component for the router.
//
// Usage:
//
//	compogo.WithComponents(
//	    http.Component,
//	    chi.Component,
//	    pprof.Component,
//	)
//
// Enable with: --trace.pprof=true --server.http.routes.pprof=/debug
var Component = &component.Component{
	Dependencies: component.Components{
		chi.Component,
	},
	Init: component.StepFunc(func(container container.Container) error {
		return container.Provide(NewConfig)
	}),
	BindFlags: component.BindFlags(func(flagSet flag.FlagSet, container container.Container) error {
		return container.Invoke(func(config *Config) {
			flagSet.BoolVar(&config.UseProfile, UseProfileFieldName, UseProfileDefault, "if true, add debug path to routing")
			flagSet.StringVar(&config.Endpoint, EndpointFieldName, EndpointDefault, "path for debug endpoint")
		})
	}),
	Configuration: component.StepFunc(func(container container.Container) error {
		return container.Invoke(Configuration)
	}),
	Execute: component.StepFunc(func(container container.Container) error {
		return container.Invoke(func(config *Config, r http.Router, logger logger.Logger) {
			if config.UseProfile {
				logger.Infof("[chi.router] add pprof endpoint - '%s'", config.Endpoint)
				r.Mount(config.Endpoint, middleware.Profiler())
			}
		})
	}),
}
