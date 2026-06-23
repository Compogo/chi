package metric

import (
	"github.com/Compogo/chi"
	"github.com/Compogo/compogo"
	"github.com/Compogo/compogo/flag"
	httpServer "github.com/Compogo/http_server"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Component — компонент metrics эндпоинта для chi-роутера.
// Добавляет эндпоинт для сбора метрик Prometheus.
var Component = compogo.Component{
	Dependencies: compogo.Components{
		&chi.Component,
	},
	Init: compogo.StepFunc(func(container compogo.Container) error {
		return container.Provide(NewConfig)
	}),
	BindFlags: compogo.BindFlags(func(flagSet flag.FlagSet, container compogo.Container) error {
		return container.Invoke(func(config *Config) {
			flagSet.StringVar(&config.Endpoint, EndpointFieldName, EndpointDefault, "path for metrics endpoint")
		})
	}),
	Configuration: compogo.StepFunc(func(container compogo.Container) error {
		return container.Invoke(Configuration)
	}),
	PreExecute: compogo.StepFunc(func(container compogo.Container) error {
		return container.Invoke(func(config *Config, r httpServer.Router, logger compogo.Logger) {
			logger.GetLogger("http").
				GetLogger("server").
				GetLogger("router").
				GetLogger("chi").
				Infof("add metrics endpoint - '%s'", config.Endpoint)

			r.Mount(config.Endpoint, promhttp.Handler())
		})
	}),
}
