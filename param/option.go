package param

import (
	"net/http"

	"github.com/Compogo/http_server/middleware/param"
	"github.com/go-chi/chi/v5"
)

// WithChiURLParam возвращает Option для извлечения параметра из URL-пути chi.
// Используется совместно с пакетом param для извлечения параметров маршрута chi.
//
// Пример:
//
//	userID := param.NewParamInt("userID", logger,
//	    param.WithUriGetter(),
//	    param.WithChiURLParam("userID"), // извлекает из chi.URLParam
//	)
//	router.Get("/users/{userID}", userID.Middleware(handler))
func WithChiURLParam(name string) param.Option {
	return param.AddGetter(func(request *http.Request) string {
		return chi.URLParam(request, name)
	})
}
