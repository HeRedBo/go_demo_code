package main

import (
	"fmt"
	"time"
)

func main() {
	p1 := make(chan int)
	p2 := make(chan int)
	p3 := make(chan int)
	exist := make(chan bool)

	go func() {
		for i := range p1 {
			fmt.Println("range p1 ", i)
			p2 <- i
			time.Sleep(1 * time.Second)
		}
		close(p2)
	}()

	go func() {
		for i := range p2 {
			fmt.Println("range p2 ", i)
			p3 <- i
			time.Sleep(1 * time.Second)
		}
		close(p3)
	}()

	go func() {
		for i := range p3 {
			fmt.Println("range p3 ", i)
		}
		exist <- true
	}()

	arr := []int{1,2,3,4,5}

	for i :=0; i < 3; i ++ {
		p1 <- arr[i]
	}
	close(p1)
	<-exist
	close(exist)
}
