package main

import (
	"fmt"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
				//runtime.Gosched()
				//fmt.Printf("Hello fom goroutine %d \n", i)
			}
		}(i)
	}
	time.Sleep(10 * time.Microsecond)
	fmt.Println(a)

}
