package api

import (
	"api-template/api/handlers"
	"github.com/go-chi/chi/v5"
)

const (
	routeFoo     = "/foo"
	routeHello   = "/hello"
	routeMetrics = "/metrics"
)

func (a *App) createRoutes(router *chi.Mux) {
	router.Get(routeMetrics, handlers.HandleMetrics(a.metrics.Registry))
	router.Get(routeHello, handlers.HandleHello())
	router.Get(routeFoo, handlers.HandleFoo(a.metrics))
}
