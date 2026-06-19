package pprof

import (
	"github.com/Compogo/chi"
	"github.com/Compogo/compogo"
	"github.com/Compogo/compogo/flag"
	httpServer "github.com/Compogo/http_server"
	"github.com/go-chi/chi/v5/middleware"
)

// Component — компонент pprof эндпоинтов для chi-роутера.
// Добавляет эндпоинты для профилирования (CPU, memory, goroutine и т.д.).
// Включается только если UseProfile = true.
var Component = &compogo.Component{
	Dependencies: compogo.Components{
		&chi.Component,
	},
	Init: compogo.StepFunc(func(container compogo.Container) error {
		return container.Provide(NewConfig)
	}),
	BindFlags: compogo.BindFlags(func(flagSet flag.FlagSet, container compogo.Container) error {
		return container.Invoke(func(config *Config) {
			flagSet.BoolVar(&config.UseProfile, UseProfileFieldName, UseProfileDefault, "if true, add debug path to routing")
			flagSet.StringVar(&config.Endpoint, EndpointFieldName, EndpointDefault, "path for debug endpoint")
		})
	}),
	Configuration: compogo.StepFunc(func(container compogo.Container) error {
		return container.Invoke(Configuration)
	}),
	PreExecute: compogo.StepFunc(func(container compogo.Container) error {
		return container.Invoke(func(config *Config, r httpServer.Router, logger compogo.Logger) {
			if config.UseProfile {
				logger.GetLogger("http").
					GetLogger("server").
					GetLogger("router").
					GetLogger("chi").
					Infof("add pprof endpoint - '%s'", config.Endpoint)

				r.Mount(config.Endpoint, middleware.Profiler())
			}
		})
	}),
}
