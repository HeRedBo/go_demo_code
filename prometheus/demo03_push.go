package main

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func main() {
	// 1. 创建指标
	jobSuccess := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "my_batch_job_success",
		Help: "批处理任务是否成功 1=成功 0=失败",
	})

	// 2. 模拟业务执行
	println("执行批处理任务中...")
	time.Sleep(2 * time.Second)

	// 3. 执行完成，推送给 PushGateway
	err := push.New("http://localhost:9091", "my_batch_job").
		Collector(jobSuccess).
		Push()

	if err != nil {
		jobSuccess.Set(0) // 失败
		println("推送失败:", err)
	} else {
		jobSuccess.Set(1) // 成功
		println("推送成功！")
	}
}
