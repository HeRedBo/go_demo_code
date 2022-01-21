package main

import (
	. "fmt"
	"strconv"
	"time"
)

func main() {
	//见识百万级并发
	for i := 0; i < 100; i++ {
		go doSomethingIII("小分队" + strconv.Itoa(i))
	}

	time.Sleep(1 * time.Second)
	Println("Game Over")
}
func doSomethingIII(grname string) {
	for {
		Println(grname, "im coroutine")
		time.Sleep(time.Second)
	}
}