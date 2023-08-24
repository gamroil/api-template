package observability

import (
	prometheus "github.com/prometheus/client_golang/prometheus"

	"go.uber.org/zap"
)

const (
	fooCounterMetric = "foo_counter"
)

// Client allows the creation and invocation of metrics within the API. Instantiation should occur through the New
// function as it creates internal resources.
type Client struct {
	logger   *zap.Logger
	metrics  map[string]prometheus.Metric
	Registry *prometheus.Registry
}

// New is the intended way to instantiate an observability Client. This method should be used over direct instantiation
// because it creates internal resources.
func New(logger *zap.Logger) *Client {
	reg := prometheus.NewRegistry()

	metrics := map[string]prometheus.Metric{
		fooCounterMetric: newFooCounter(reg),
	}

	return &Client{
		logger:   logger,
		metrics:  metrics,
		Registry: reg,
	}
}

func (c *Client) IncrementFooCount() {
	c.metrics[fooCounterMetric].(prometheus.Counter).Add(1)
}
