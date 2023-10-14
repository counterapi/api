package controllers

import (
	"net/http"

	"github.com/counterapi/api/pkg/repositories"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	namespaceCounter = "namespace_counter"
	counterCounter   = "counter_counter"
	countCounter     = "count_counter"
)

// MetricController controls Health operations.
type MetricController struct {
	CounterRepository repositories.CounterRepository

	Metrics  map[string]prometheus.Collector
	Registry *prometheus.Registry
}

func (m *MetricController) setMetrics() error {
	m.Metrics = map[string]prometheus.Collector{
		namespaceCounter: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "counterapi",
			Name:      namespaceCounter,
			Help:      "Number of namespaces.",
		}),
		counterCounter: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "counterapi",
			Name:      counterCounter,
			Help:      "Number of counters.",
		}),
		countCounter: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "counterapi",
			Name:      countCounter,
			Help:      "Number of total count.",
		}),
	}

	err := m.Registry.Register(prometheus.NewRegistry())
	if err != nil {
		return err
	}

	for _, v := range m.Metrics {
		err = m.Registry.Register(v)
		if err != nil {
			return err
		}
	}

	return nil
}

// Serve returns all metrics.
func (m *MetricController) Serve(c *gin.Context) {
	err := m.SetMetrics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "could not get metrics")
	}

	handler := promhttp.HandlerFor(m.Registry, promhttp.HandlerOpts{})
	handler.ServeHTTP(c.Writer, c.Request)
}

func (m *MetricController) SetMetrics() error {
	counterCount, err := m.CounterRepository.CountCounters()
	if err != nil {
		return err
	}

	m.Metrics[counterCounter].(prometheus.Gauge).Set(float64(counterCount))

	countCount, err := m.CounterRepository.CountCounts()
	if err != nil {
		return err
	}

	m.Metrics[countCounter].(prometheus.Gauge).Set(float64(countCount))

	namespaceCount, err := m.CounterRepository.CountNamespaces()
	if err != nil {
		return err
	}

	m.Metrics[namespaceCounter].(prometheus.Gauge).Set(float64(namespaceCount))

	return nil
}

// NewMetricController returns a new MetricController.
func NewMetricController(counterRepository repositories.CounterRepository) MetricController {
	ctrl := MetricController{
		CounterRepository: counterRepository,
		Registry:          prometheus.NewRegistry(),
	}

	_ = ctrl.setMetrics()

	return ctrl
}
