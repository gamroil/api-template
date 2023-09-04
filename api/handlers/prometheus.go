package handlers

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func HandleMetrics(reg *prometheus.Registry) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		h := promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})
		h.ServeHTTP(w, r)
	}
}
