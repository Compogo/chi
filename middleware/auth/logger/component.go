package logger

import (
	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
	"github.com/Compogo/http"
	"github.com/Compogo/http/middleware/logger"
)

var (
	// RequestComponent automatically adds request logging middleware.
	// It logs request bodies at DEBUG level.
	RequestComponent = &component.Component{
		Dependencies: component.Components{
			logger.RequestComponent,
		},
		Run: component.StepFunc(func(container container.Container) error {
			return container.Invoke(func(r http.Router, logger *logger.Request) {
				r.Use(logger)
			})
		}),
	}

	// ResponseComponent automatically adds response logging middleware.
	// It logs response bodies at DEBUG level.
	ResponseComponent = &component.Component{
		Dependencies: component.Components{
			logger.ResponseComponent,
		},
		Run: component.StepFunc(func(container container.Container) error {
			return container.Invoke(func(r http.Router, logger *logger.Response) {
				r.Use(logger)
			})
		}),
	}
)
