package logger

import (
	"github.com/Compogo/compogo"
	httpServer "github.com/Compogo/http_server"
	"github.com/Compogo/http_server/middleware/logger"
)

var (
	// RequestComponent — компонент логирования запросов для chi-роутера.
	RequestComponent = compogo.Component{
		Dependencies: compogo.Components{
			&logger.RequestComponent,
		},
		PreExecute: compogo.StepFunc(func(container compogo.Container) error {
			return container.Invoke(func(r httpServer.Router, logger *logger.Request) {
				r.Use(logger)
			})
		}),
	}

	// ResponseComponent — компонент логирования ответов для chi-роутера.
	ResponseComponent = compogo.Component{
		Dependencies: compogo.Components{
			&logger.ResponseComponent,
		},
		PreExecute: compogo.StepFunc(func(container compogo.Container) error {
			return container.Invoke(func(r httpServer.Router, logger *logger.Response) {
				r.Use(logger)
			})
		}),
	}
)
