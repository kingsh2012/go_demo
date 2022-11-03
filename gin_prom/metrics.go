package main

import "github.com/prometheus/client_golang/prometheus"

var (
	// HttpHistogram prometheus 模型
	HttpHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace:   "gin_route",
		Subsystem:   "",
		Name:        "requests_milliseconds",
		Help:        "接口请求时长",
		ConstLabels: nil,
		Buckets:     []float64{10, 20, 30, 60, 80, 100, 600, 900, 1000, 1300, 1600, 1900, 2100, 5000, 10000, 20000, 30000},
	}, []string{"method", "code", "uri"})
)

func init() {
	prometheus.MustRegister(HttpHistogram)
}
