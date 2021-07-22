package main

import (
	"fmt"
	"strconv"
	"time"
)

func main011()  {
	//在一条独立的协程中执行匿名函数
	go func() {

		for{
			fmt.Println("im coroutine")
			time.Sleep(time.Second)
		}

	}()

	for {
		fmt.Println("我是主协程")
		time.Sleep(time.Second)
	}
}


func main012() {
	//在一条独立的协程中执行doSomething()
	go doSomething()

	for {
		fmt.Println("我是主协程")
		time.Sleep(time.Second)
	}
}

func doSomething() {
	for {
		fmt.Println("im coroutine")
		time.Sleep(time.Second)
	}
}


func main() {
	//见识百万级并发
	for i:=0;i<1000000 ;i++  {
		go doSomethingII("小分队"+strconv.Itoa(i))
	}

	for {
		fmt.Println("我是主协程")
		time.Sleep(time.Second)
	}
}

func doSomethingII(grname string) {
	for {
		fmt.Println(grname, "im coroutine")
		time.Sleep(time.Second)
	}

}