package token

import (
	"github.com/Compogo/chi"
	"github.com/Compogo/compogo"
	httpServer "github.com/Compogo/http_server"
	"github.com/Compogo/http_server/middleware/auth/token"
)

// Component — компонент Token Auth для chi-роутера.
// Подключает Token Auth middleware ко всем маршрутам роутера.
var Component = &compogo.Component{
	Dependencies: compogo.Components{
		&token.Component,
		&chi.Component,
	},
	PreExecute: compogo.StepFunc(func(container compogo.Container) error {
		return container.Invoke(func(r httpServer.Router, auth *token.Auth) {
			r.Use(auth)
		})
	}),
}
