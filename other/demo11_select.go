package main

import (
	"fmt"
	"github.com/gookit/goutil/dump"
	"time"
)

// go select 使用场景理解1 超时处理
func main1101() {
	ch := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
	}()
	select {
	case val := <-ch:
		fmt.Println("Received value:", val)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout!")
		dump.P()
	}
}

// go select 使用场景理解2 多路复用
func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- 2
	}()
	select {
	case val := <-ch1:
		fmt.Println("Received from ch1:", val)
	case val := <-ch2:
		fmt.Println("Received from ch2:", val)
	}
}

// go select 使用场景理解3. 非阻塞通道操作
func main1103() {
	ch := make(chan int)
	select {
	case ch <- 1:
		fmt.Println("Sent value to channel")
	default:
		fmt.Println("Channel is not ready to receive")
	}

}

// 空的 select 语句：空的 select 语句会永远阻塞，例如

// 通道关闭的影响
func main1104() {
	ch := make(chan int)
	close(ch)
	select {
	case val, ok := <-ch:
		if ok {
			fmt.Println("Received value:", val)
		} else {
			fmt.Println("Channel is closed")
		}
	}
}
