package chi

import (
	httpServer "github.com/Compogo/http_server"
	"github.com/go-chi/chi/v5"
)

// Decorator — обёртка над chi.Router, реализующая интерфейс http_server.Router.
// Позволяет использовать chi в качестве роутера в Compogo HTTP-сервере.
type Decorator struct {
	chi.Router
}

// NewDecorator создаёт новый декоратор для chi.Router.
func NewDecorator(router chi.Router) *Decorator {
	return &Decorator{Router: router}
}

// Use добавляет middleware к роутеру.
// Реализует интерфейс http_server.Router.
func (d *Decorator) Use(middlewares ...httpServer.Middleware) {
	for _, middleware := range middlewares {
		d.Router.Use(middleware.Middleware)
	}
}

// Group создаёт группу маршрутов с общим префиксом и middleware.
// Реализует интерфейс http_server.Router.
func (d *Decorator) Group(fn func(r httpServer.Router)) {
	d.Router.Group(func(r chi.Router) {
		fn(NewDecorator(r))
	})
}

// Route создаёт под-роутер для организации маршрутов.
// Реализует интерфейс http_server.Router.
func (d *Decorator) Route(pattern string, fn func(r httpServer.Router)) {
	d.Router.Route(pattern, func(r chi.Router) {
		fn(NewDecorator(r))
	})
}
