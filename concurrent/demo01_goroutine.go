package main

import (
	"fmt"
	"time"
)

func add(a, b int) {
	var c = a + b
	fmt.Printf("%d + %d = %d", a, b, c)
	fmt.Println()
}

func main() {
	// go add(1 ,2)
	for i := 0; i < 10; i++ {
		go add(1, i)
	}
	time.Sleep(1e9)
}
