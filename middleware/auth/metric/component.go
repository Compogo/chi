package metric

import (
	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
	"github.com/Compogo/http"
	"github.com/Compogo/http/middleware/metric"
)

var (
	// RequestCountComponent automatically adds request count metrics middleware.
	// It exports: compogo_http_server_requests_total{app, code, endpoint}
	RequestCountComponent = &component.Component{
		Dependencies: component.Components{
			metric.RequestCountComponent,
		},
		Run: component.StepFunc(func(container container.Container) error {
			return container.Invoke(func(r http.Router, metric *metric.RequestCount) {
				r.Use(metric)
			})
		}),
	}

	// DurationComponent automatically adds request duration metrics middleware.
	// It exports: compogo_http_server_duration_seconds{app, endpoint}
	DurationComponent = &component.Component{
		Dependencies: component.Components{
			metric.DurationComponent,
		},
		Run: component.StepFunc(func(container container.Container) error {
			return container.Invoke(func(r http.Router, metric *metric.Duration) {
				r.Use(metric)
			})
		}),
	}
)
