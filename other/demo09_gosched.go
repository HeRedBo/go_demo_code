package main

import (
	"fmt"
	"runtime"
	"time"
)

func say2(s string) {
	for i := 0; i < 3; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main2() {
	go say2("world")
	say2("hello2")
}

func main() {
	go task071()
	go task072()
	time.Sleep(3 * time.Second)
}

func task071() {
	for _, c := range "今生注定我爱你" {
		fmt.Printf("%c\n", c)
		runtime.Gosched()
		time.Sleep(time.Nanosecond)
	}
}

func task072() {
	for _, c := range "FUCKOFF" {
		fmt.Printf("%c\n", c)
		runtime.Gosched()
		time.Sleep(time.Nanosecond)
	}
}
