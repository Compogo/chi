package chi

import (
	"compress/gzip"

	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
	"github.com/Compogo/compogo/logger"
	"github.com/Compogo/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Component is a ready-to-use Compogo component that provides a chi router.
// It automatically:
//   - Creates a chi router with standard middleware (recoverer, compress, request logger)
//   - Provides the router as both chi.Router and http.Router
//   - Attaches the router to the HTTP server during Run phase
//
// Usage:
//
//	compogo.WithComponents(
//	    http.Component,
//	    chi.Component,
//	    // ... other components that need the router
//	)
var Component = &component.Component{
	Dependencies: component.Components{
		http.Component,
	},
	Init: component.StepFunc(func(container container.Container) error {
		return container.Provides(
			func(logger logger.Logger) chi.Router {
				router := chi.NewRouter()

				router.Use(
					middleware.Recoverer,
					middleware.Compress(gzip.BestSpeed),
					middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: logger, NoColor: true}),
				)

				return router
			},
			func(router chi.Router) http.Router { return NewDecorator(router) },
		)
	}),
	Execute: component.StepFunc(func(container container.Container) error {
		return container.Invoke(func(r http.Router, server http.Server) {
			server.SetRouter(r)
		})
	}),
}
