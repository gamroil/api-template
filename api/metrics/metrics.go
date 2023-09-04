package metrics

import "github.com/prometheus/client_golang/prometheus"

func newFooCounter(reg *prometheus.Registry) prometheus.Counter {
	m := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "foo_request_total",
		Help: "Total number of requests to the foo endpoint",
	})
	reg.MustRegister(m)
	return m
}
