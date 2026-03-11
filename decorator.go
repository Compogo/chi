package chi

import (
	"github.com/Compogo/http"
	"github.com/go-chi/chi/v5"
)

// Decorator adapts a chi.Router to implement the http.Router interface.
// This allows any chi router to be used with Compogo's HTTP server component.
type Decorator struct {
	chi.Router
}

// NewDecorator creates a new Decorator wrapping the provided chi router.
func NewDecorator(router chi.Router) *Decorator {
	return &Decorator{Router: router}
}

// Use implements http.Router.Use by converting Compogo middlewares to chi middlewares.
// It iterates through the provided middlewares and adds them to the underlying chi router.
func (d *Decorator) Use(middlewares ...http.Middleware) {
	for _, middleware := range middlewares {
		d.Router.Use(middleware.Middleware)
	}
}

// Group implements http.Router.Group by creating a sub-router with inherited middleware.
// The provided function receives a wrapped http.Router that operates on the chi sub-router.
func (d *Decorator) Group(fn func(r http.Router)) {
	d.Router.Group(func(r chi.Router) {
		fn(NewDecorator(r))
	})
}

// Route implements http.Router.Route by creating a new sub-router with the given prefix.
// The provided function receives a wrapped http.Router that operates on the chi sub-router.
func (d *Decorator) Route(pattern string, fn func(r http.Router)) {
	d.Router.Route(pattern, func(r chi.Router) {
		fn(NewDecorator(r))
	})
}
