package metric

import (
	"github.com/Compogo/compogo"
	httpServer "github.com/Compogo/http_server"
	"github.com/Compogo/http_server/middleware/metric"
)

var (
	// RequestCountComponent — компонент метрики количества запросов для chi-роутера.
	RequestCountComponent = compogo.Component{
		Dependencies: compogo.Components{
			metric.RequestCountComponent,
		},
		PreExecute: compogo.StepFunc(func(container compogo.Container) error {
			return container.Invoke(func(r httpServer.Router, metric *metric.RequestCount) {
				r.Use(metric)
			})
		}),
	}

	// DurationComponent — компонент метрики длительности запросов для chi-роутера.
	DurationComponent = compogo.Component{
		Dependencies: compogo.Components{
			metric.DurationComponent,
		},
		PreExecute: compogo.StepFunc(func(container compogo.Container) error {
			return container.Invoke(func(r httpServer.Router, metric *metric.Duration) {
				r.Use(metric)
			})
		}),
	}
)
