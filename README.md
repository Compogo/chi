# Compogo Chi

[![Go Reference](https://pkg.go.dev/badge/github.com/Compogo/chi.svg)](https://pkg.go.dev/github.com/Compogo/chi)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Адаптер [chi](https://github.com/go-chi/chi) для [HTTP-сервера Compogo](https://github.com/Compogo/http_server).

Предоставляет:

* Реализацию `http_server.Router` через chi
* Готовые компоненты для middleware (auth, logger, metrics)
* Встроенные эндпоинты (health-check, metrics, pprof)
* Интеграцию с параметрами запроса через chi.URLParam

## Установка

```shell
go get github.com/Compogo/chi
```

## Быстрый старт

```go
package main

import (
    "net/http"

    "github.com/Compogo/chi"
    "github.com/Compogo/compogo"
    "github.com/Compogo/http_server"
    "github.com/Compogo/runner"
)

func main() {
    app := compogo.NewApp("myapp",
        // Базовые компоненты
        compogo.WithComponents(&runner.Component),
        compogo.WithComponents(&http_server.Component),

        // Chi роутер
        compogo.WithComponents(&chi.Component),
    )

    // Добавление маршрутов
    app.AddComponents(&compogo.Component{
        Name: "routes",
        PreExecute: compogo.StepFunc(func(container compogo.Container) error {
            return container.Invoke(func(r http_server.Router) error {
                r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
                    w.Write([]byte("Hello, World!"))
                })
                return nil
            })
        }),
    })

    if err := app.Serve(); err != nil {
        panic(err)
    }
}
```

## Компоненты

### Базовый роутер

```go
import "github.com/Compogo/chi"

app.AddComponents(&chi.Component)

// Роутер автоматически добавляет middleware:
// - Recoverer — восстановление после паник
// - RequestLogger — логирование запросов (через DefaultLogFormatter)
```

### Аутентификация

#### Basic Auth

```go
import "github.com/Compogo/chi/middleware/basic"

app.AddComponents(&basic.Component)

// Настройка через флаги
// --server.http.auth.basic.creds=admin:password
// --server.http.auth.basic.filepath=/path/to/creds.txt
```

#### Token Auth

```go
import "github.com/Compogo/chi/middleware/token"

app.AddComponents(&token.Component)

// Настройка через флаги
// --server.http.auth.token.tokens=abc123,def456
// --server.http.auth.token.header=X-Auth-Token
```

#### Логирование

```go
import "github.com/Compogo/chi/middleware/logger"

// Только запросы
app.AddComponents(&logger.RequestComponent)

// Только ответы
app.AddComponents(&logger.ResponseComponent)

// Оба
app.AddComponents(&logger.RequestComponent, &logger.ResponseComponent)
```

#### Метрики Prometheus

```go
import "github.com/Compogo/chi/middleware/metric"

// Количество запросов
app.AddComponents(&metric.RequestCountComponent)

// Длительность запросов
app.AddComponents(&metric.DurationComponent)

// Оба
app.AddComponents(&metric.RequestCountComponent, &metric.DurationComponent)
```

### Готовые эндпоинты

#### Health Check

```go
import "github.com/Compogo/chi/routes/health_check"

app.AddComponents(&health_check.Component)

// Настройка через флаг
// --server.http.routes.health_check=/health-check
```

#### Metrics

```go
import "github.com/Compogo/chi/routes/metric"

app.AddComponents(&metric.Component)

// Настройка через флаг
// --server.http.routes.metrics=/metrics
```

#### Pprof (профилирование)

```go
import "github.com/Compogo/chi/routes/pprof"

app.AddComponents(&pprof.Component)

// Настройка через флаги
// --trace.pprof=true
// --server.http.routes.pprof=/debug
```

## Параметры запроса

### Извлечение параметров из URL

```go
import (
    "github.com/Compogo/chi/param"
    "github.com/Compogo/http_server/middleware/param"
)

// Создание параметра с извлечением из chi.URLParam
userID := param.NewParamInt("userID", logger,
    param.WithChiURLParam("userID"),
    param.AddValidator(param.GteValidator(1)),
)

// Использование в роутере
router.Get("/users/{userID}", userID.Middleware(handler))

// Получение из контекста
userID := r.Context().Value("userID").(int)
```

## Зависимости

* [Compogo](https://github.com/Compogo/compogo) — основной фреймворк
* [Compogo HTTP Server](https://github.com/Compogo/http_server) — HTTP-сервер
* [go-chi/chi](https://github.com/go-chi/chi) — роутер
* [Prometheus](https://github.com/prometheus/client_golang) — метрики

## Лицензия

```plantuml
MIT License

Copyright (c) 2026 Compogo

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

```
