package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"
)

func main01() {
	newRandStream := func() <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			// 死循环：不断向channel中放数据，直到阻塞
			for {
				randStream <- rand.Int()
			}
		}()

		return randStream
	}

	randStream := newRandStream()
	fmt.Println("3 random ints:")

	// 只消耗3个数据，然后去做其他的事情，此时生产者阻塞，
	// 若主goroutine不处理生产者goroutine，则就产生了泄露
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}

	fmt.Fprintf(os.Stderr, "%d\n", runtime.NumGoroutine())
	time.Sleep(10e9)
	fmt.Fprintf(os.Stderr, "%d\n", runtime.NumGoroutine())

}

func main() {

	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)

			for {
				select {
				case randStream <- rand.Int():
				case <-done:  // 得到通知，结束自己
					return
				}
			}
		}()

		return randStream
	}


	done := make(chan interface{})
	randStream := newRandStream(done)
	fmt.Println("3 random ints:")

	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}

	// 通知子协程结束自己
	// done <- struct{}{}
	close(done)
	// Simulate ongoing work
	time.Sleep(1 * time.Second)
}