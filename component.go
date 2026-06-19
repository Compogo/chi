package chi

import (
	"github.com/Compogo/compogo"
	httpServer "github.com/Compogo/http_server"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Component — компонент роутера chi для Compogo.
// Регистрирует chi.Router и http_server.Router в DI-контейнере.
// Автоматически добавляет middleware Recoverer и RequestLogger.
var Component = compogo.Component{
	Name: "http.server.router.chi",
	Dependencies: compogo.Components{
		&httpServer.Component,
	},
	Init: compogo.StepFunc(func(container compogo.Container) error {
		return container.Provides(
			func(logger compogo.Logger) chi.Router {
				router := chi.NewRouter()

				router.Use(
					middleware.Recoverer,
					middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: logger, NoColor: true}),
				)

				return router
			},
			func(router chi.Router) httpServer.Router { return NewDecorator(router) },
		)
	}),
	PostExecute: compogo.StepFunc(func(container compogo.Container) error {
		return container.Invoke(func(r httpServer.Router, server httpServer.Server) {
			server.SetRouter(r)
		})
	}),
}
