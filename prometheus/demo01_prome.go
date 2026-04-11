package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// 1. 定义一个指标：程序启动秒数
var (
	upTime = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "go_app_uptime_seconds", // 指标名
			Help: "程序已运行时间",
		},
	)
)

func main011() {
	// 2. 模拟运行时间
	go func() {
		for {
			upTime.Inc() // 每秒 +1
			time.Sleep(1 * time.Second)
		}
	}()

	// 3. 暴露 /metrics 接口
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8082", nil)
}
