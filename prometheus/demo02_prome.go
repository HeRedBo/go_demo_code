package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// 1. 定义业务指标
var (
	// 请求总数（计数器）
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{Name: "http_requests_total", Help: "HTTP请求总数"},
		[]string{"path", "method"}, // 标签：路径、请求方式
	)

	// 请求耗时（直方图）
	httpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "请求耗时分布",
			Buckets: []float64{0.1, 0.2, 0.5, 1, 2}, // 耗时区间
		},
		[]string{"path"},
	)
)

// 模拟接口
func helloHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// 指标 +1
	httpRequestsTotal.WithLabelValues(r.URL.Path, r.Method).Inc()

	w.Write([]byte("hello prometheus!"))

	// 记录耗时
	httpRequestDuration.WithLabelValues(r.URL.Path).Observe(time.Since(start).Seconds())
}

func main002() {
	http.HandleFunc("/hello", helloHandler)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8082", nil)
}
