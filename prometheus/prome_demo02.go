package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/push"
)

// 定义指标
var (
	// 计数器：记录请求总数
	requestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "myapp_requests_total",
			Help: "Total number of requests",
		},
	)

	// 直方图：记录请求处理时间
	requestDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "myapp_request_duration_seconds",
			Help:    "Request duration in seconds",
			Buckets: prometheus.DefBuckets, // 默认分桶
		},
	)

	// 仪表盘：记录当前活跃请求数
	activeRequests = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "myapp_active_requests",
			Help: "Current number of active requests",
		},
	)

	//  summaries：记录请求大小
	requestSize = prometheus.NewSummary(
		prometheus.SummaryOpts{
			Name:       "myapp_request_size_bytes",
			Help:       "Request size in bytes",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001}, // 分位数
		},
	)
)

func init() {
	// 注册指标
	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(requestDuration)
	prometheus.MustRegister(activeRequests)
	prometheus.MustRegister(requestSize)
}

func main2() {
	// 启动一个 goroutine 定期向 Pushgateway 推送指标
	go pushMetricsToGateway()

	// 注册 HTTP 处理函数
	http.HandleFunc("/", handleRequest)
	// 暴露 Prometheus 指标的端点
	http.Handle("/metrics", promhttp.Handler())

	// 启动服务器
	port := "8080"
	log.Printf("Server starting on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// 处理 HTTP 请求
func handleRequest(w http.ResponseWriter, r *http.Request) {
	// 增加活跃请求数
	activeRequests.Inc()
	defer activeRequests.Dec()

	// 记录请求开始时间
	start := time.Now()
	defer func() {
		// 记录请求处理时间
		requestDuration.Observe(time.Since(start).Seconds())
		// 增加请求计数
		requestsTotal.Inc()
		// 记录请求大小
		requestSize.Observe(float64(r.ContentLength))
	}()

	// 模拟处理时间
	time.Sleep(100 * time.Millisecond)

	// 返回响应
	fmt.Fprintf(w, "Hello, Prometheus! Current time: %s", time.Now().Format(time.RFC3339))
}

// 向 Pushgateway 推送指标
func pushMetricsToGateway() {
	// 定义一个计数器，用于演示 Pushgateway
	jobCompletionTime := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "myapp_job_completion_time_seconds",
			Help: "Time taken to complete a job",
		},
	)

	// 定期推送指标
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		// 模拟作业完成时间
		jobCompletionTime.Set(float64(time.Now().UnixNano()) / float64(time.Second.Nanoseconds()))

		// 推送指标到 Pushgateway
		err := push.New("http://localhost:9091", "myapp_job").
			Collector(jobCompletionTime).
			Grouping("instance", "localhost").
			Push()

		if err != nil {
			log.Printf("Error pushing metrics to Pushgateway: %v", err)
		} else {
			log.Println("Metrics pushed to Pushgateway successfully")
		}
	}
}
