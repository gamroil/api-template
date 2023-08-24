package api

import (
	"api-template/api/handlers"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	routeFoo     = "/foo"
	routeHello   = "/hello"
	routeMetrics = "/metrics"
)

func (a *App) createRoutes(router *gin.Engine) {
	prometheusHandler := func(c *gin.Context) {
		h := promhttp.HandlerFor(a.metrics.Registry, promhttp.HandlerOpts{Registry: a.metrics.Registry})
		h.ServeHTTP(c.Writer, c.Request)
	}

	router.GET(routeMetrics, prometheusHandler)

	router.GET(routeHello, handlers.HandleHello())
	router.GET(routeFoo, handlers.HandleFoo(a.metrics))
}
