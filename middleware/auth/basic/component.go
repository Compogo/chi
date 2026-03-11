package basic

import (
	"github.com/Compogo/chi"
	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
	"github.com/Compogo/http"
	"github.com/Compogo/http/middleware/auth/basic"
)

// Component automatically adds basic authentication middleware to the HTTP router.
// It depends on:
//   - basic.Component (the authentication middleware itself)
//   - chi.Component (the router)
//
// Usage:
//
//	compogo.WithComponents(
//	    http.Component,
//	    chi.Component,
//	    basic.Component,
//	)
//
// All routes will now require valid basic authentication credentials.
var Component = &component.Component{
	Dependencies: component.Components{
		basic.Component,
		chi.Component,
	},
	Run: component.StepFunc(func(container container.Container) error {
		return container.Invoke(func(r http.Router, auth *basic.Auth) {
			r.Use(auth)
		})
	}),
}
