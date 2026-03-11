package param

import (
	"net/http"

	"github.com/Compogo/http/middleware/param"
	"github.com/go-chi/chi/v5"
)

// WithChiURLParam returns a param.Option that extracts a value from chi's URL parameters.
// This allows using chi's named path parameters (e.g., /users/{id}) with Compogo's param system.
//
// Example:
//
//	userID := param.NewParamInt("id", logger,
//	    param.WithChiURLParam("id"),  // extracts from /users/{id}
//	)
func WithChiURLParam(name string) param.Option {
	return param.AddGetter(func(request *http.Request) string {
		return chi.URLParam(request, name)
	})
}
