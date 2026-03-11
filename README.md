# Compogo Chi 🔀

**Compogo Chi** — это готовая интеграция [chi router](https://github.com/go-chi/chi) с фреймворком [Compogo](https://github.com/Compogo/compogo). Предоставляет базовый роутер со стандартными middleware и набор компонентов для подключения аутентификации, метрик, логирования, pprof и health checks.

## 🚀 Установка

```bash
go get github.com/Compogo/chi
```

### 📦 Быстрый старт

```go
package main

import (
    "github.com/Compogo/compogo"
    "github.com/Compogo/runner"
    "github.com/Compogo/http"
    "github.com/Compogo/chi"
    "github.com/Compogo/chi/metric"
    "github.com/Compogo/chi/health_check"
)

func main() {
    app := compogo.NewApp("myapp",
        compogo.WithOsSignalCloser(),
        runner.WithRunner(),
        http.WithServer(),
        chi.Component,                    // базовый chi-роутер
        metric.Component,                  // prometheus метрики на /metrics
        health_check.Component,            // health check на /health-check
        compogo.WithComponents(
            myHandlerComponent,
        ),
    )

    if err := app.Serve(); err != nil {
        panic(err)
    }
}

// Ваш компонент с обработчиками
var myHandlerComponent = &component.Component{
    Dependencies: component.Components{chi.Component},
    Run: component.StepFunc(func(c container.Container) error {
        return c.Invoke(func(r http.Router) {
            r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
                w.Write([]byte("Hello, World!"))
            })
        })
    }),
}
```

### ✨ Возможности

#### 🎯 Базовый роутер (chi.Component)

Создаёт `chi.Router` со стандартными middleware:

* Recoverer — защита от паник
* Compress — gzip-сжатие ответов
* RequestLogger — логирование запросов

Автоматически подключается к HTTP-серверу в Run-фазе.

#### 🔌 Интеграция с param

```go
import "github.com/Compogo/chi/param"

page := param.NewParamInt("page", logger,
    param.WithChiURLParam("page"),  // из /users/{page}
    param.WithDefault(1),
)
```

### 🧩 Примеры

#### Полный набор middleware

```go
app := compogo.NewApp("myapp",
    compogo.WithOsSignalCloser(),
    runner.WithRunner(),
    http.WithServer(),
    chi.Component,
    chi_metric.RequestCountComponent,
    chi_metric.DurationComponent,
    chi_logger.RequestComponent,
    chi_logger.ResponseComponent,
    chi_auth_basic.Component,
    chi_pprof.Component,
    chi_health_check.Component,
    compogo.WithComponents(
        yourComponent,
    ),
)
```

#### Группы маршрутов

```go
r.Group(func(r http.Router) {
    r.Use(authMiddleware)
    r.Get("/admin", adminHandler)
    r.Get("/profile", profileHandler)
})

r.Route("/api/v1", func(r http.Router) {
    r.Get("/users", listUsers)
    r.Post("/users", createUser)
})
```
