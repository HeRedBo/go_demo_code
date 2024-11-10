package main

import (
	"fmt"
	"sync"
	"time"
)

// 业务开发实践

// 该函数的参数为多个业务逻辑函数，且函数个数为变长参数
func allTasks(tasks ...func()) {
	var wg sync.WaitGroup
	for _, t := range tasks {
		wg.Add(1)           // 每次启动添加一个协程等待组
		go func(f func()) { // 匿名函数的参数为业务逻辑函数
			defer func() {
				// 在每个协程内部接收该协程自身抛出来的 panic
				if err := recover(); err != nil {
					fmt.Println("defer panic ", err)
				}
				wg.Done() // 每个协程结束时给 等待组减 1
			}()
			f() // 调用业务函数并执行
		}(t)

	}
	wg.Wait()
}

// 业务逻辑 A
func A() {
	fmt.Println("A func begin")
	panic("error A")
}

// 业务逻辑 B
func B() {
	fmt.Println("B func begin")
}

func main() {
	allTasks(A, B) // 将业务逻辑函数名 A B 传递给封装好的处理函数
	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}
