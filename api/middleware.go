package api

import (
	"api-template/api/middleware"
	"net/http"
)

func (a *App) NewMetricsMiddleware(next http.Handler) http.Handler {
	return middleware.NewPrometheusMiddleware(a.metrics.Registry, appName)(next)
}
