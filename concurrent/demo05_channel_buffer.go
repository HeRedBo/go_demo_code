package main

import (
	"fmt"
	"time"
)

func test(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- 1
	}
	close(ch)
}

func main() {
	start := time.Now()
	ch := make(chan int, 20)
	//ch := make(chan int)
	go test(ch)
	for i := range ch {
		fmt.Println("接收到的数据:", i)
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
