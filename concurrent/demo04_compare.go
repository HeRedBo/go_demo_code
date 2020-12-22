package main

import (
	"fmt"
	"time"
)

func add4(a, b int) {
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
}

func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		add4(1, i)
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s):", consume)
}
