package main

import "testing"

func TestProducerConsumer(t *testing.T) {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}

func TestWorker(t *testing.T) {
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	// 启动3个worker
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 发送5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// 输出结果
	for a := 1; a <= 5; a++ {
		<-results
	}
}
