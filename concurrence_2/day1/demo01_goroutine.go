package main

import (
	"fmt"
	"strconv"
	"time"
)

// 协程并发
func main001() {

	//在一条独立的协程中执行匿名函数
	go func() {
		for {
			fmt.Println("im coroutine")
			time.Sleep(time.Second)
		}
	} ()

	for {
		fmt.Println("我是主协程")
		time.Sleep(time.Second)
	}
}


func main002() {
	go doSomeThing()

	for {
		fmt.Println("我是主协程")
		time.Sleep(time.Second)
	}
}

func doSomeThing() {
	for  {
		fmt.Println("im coroutine")
		time.Sleep(time.Second)
	}
}



func main() {
	//见识百万级并发
	for i:=0;i<10 ;i++  {
		go doSomethingII("小分队"+strconv.Itoa(i))
	}

	for {
		fmt.Println("我是主协程")
		time.Sleep(time.Second)
	}
}

func doSomethingII(grname string) {
	for {
		fmt.Println(grname,"im coroutine")
		time.Sleep(time.Second)
	}
}

