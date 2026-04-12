package main

import "fmt"

// 1. 生产者 - 消费者模式（最基础）
// 核心
// 一个 / 多个协程生产数据，一个 / 多个协程消费数据，Channel 做通信。
// 适用场景
// 日志收集、数据处理、任务生产
func producer1(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i // 生产数据
	}
	close(ch) // 生产完毕关闭通道
}

func consumer1(ch chan int) {
	for v := range ch { // 消费数据
		fmt.Println("消费:", v)
	}
}
