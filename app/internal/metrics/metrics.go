package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// TODO USE IT
var (
	TotalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint"},
	)

	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of response time for HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)
)

func Init() {
	prometheus.MustRegister(TotalRequests)
	prometheus.MustRegister(RequestDuration)
}
