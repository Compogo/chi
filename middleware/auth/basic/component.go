package basic

import (
	"github.com/Compogo/chi"
	"github.com/Compogo/compogo"
	httpServer "github.com/Compogo/http_server"
	"github.com/Compogo/http_server/middleware/auth/basic"
)

// Component — компонент Basic Auth для chi-роутера.
// Подключает Basic Auth middleware ко всем маршрутам роутера.
var Component = compogo.Component{
	Dependencies: compogo.Components{
		&basic.Component,
		&chi.Component,
	},
	PreExecute: compogo.StepFunc(func(container compogo.Container) error {
		return container.Invoke(func(r httpServer.Router, auth *basic.Auth) {
			r.Use(auth)
		})
	}),
}
