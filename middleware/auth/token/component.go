package token

import (
	"github.com/Compogo/chi"
	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
	"github.com/Compogo/http"
	"github.com/Compogo/http/middleware/auth/token"
)

// Component automatically adds token authentication middleware to the HTTP router.
// It depends on:
//   - token.Component (the authentication middleware itself)
//   - chi.Component (the router)
//
// Usage:
//
//	compogo.WithComponents(
//	    http.Component,
//	    chi.Component,
//	    token.Component,
//	)
//
// All routes will now require a valid token in the configured header.
var Component = &component.Component{
	Dependencies: component.Components{
		token.Component,
		chi.Component,
	},
	Run: component.StepFunc(func(container container.Container) error {
		return container.Invoke(func(r http.Router, auth *token.Auth) {
			r.Use(auth)
		})
	}),
}
