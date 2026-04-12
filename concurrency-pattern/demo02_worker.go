package main

import (
	"fmt"
	"time"
)

// Worker 池模式（控制并发数）
//核心
//固定数量的协程处理任务，防止协程爆炸
//适用场景
//爬虫、接口请求、批量任务

// 固定3个worker
func worker2(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker%d 处理任务%d\n", id, j)
		time.Sleep(100)
		results <- j * 2
	}
}
