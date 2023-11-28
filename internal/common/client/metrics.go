package client

import "github.com/prometheus/client_golang/prometheus"

var (
	MetricsSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "server",
		Subsystem: "requests",
		Name:      "duration_sec",
		Help:      "server requests duration(sec).",
		Buckets:   []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.250, 0.5, 1},
	}, []string{"kind", "operation"})

	MetricsRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "client",
		Subsystem: "requests",
		Name:      "code_total",
		Help:      "The total number of processed requests",
	}, []string{"kind", "operation", "code", "reason"})
	MetricsLoads = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "server",
		Subsystem: "system",
		Name:      "load_total",
		Help:      "The load of cpu & memory",
	}, []string{"kind", "operation", "code", "reason"})
)

func init() {
	prometheus.MustRegister(MetricsSeconds, MetricsRequests, MetricsLoads)
}
