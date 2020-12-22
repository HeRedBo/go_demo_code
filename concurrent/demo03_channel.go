package main

import (
	"fmt"
	"time"
)

func add3(a, b int, ch chan int) {
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	ch <- 1 // 表示把元素 1 发送到通道 ch
}

func main() {
	start := time.Now()
	chs := make([]chan int, 1000)
	for i := 0; i < 1000; i++ {
		chs[i] = make(chan int)
		go add3(1, i, chs[i])
	}
	// 留空表示忽略
	for _, ch := range chs {
		<-ch
	}

	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s):", consume)
}
