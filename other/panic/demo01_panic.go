package main

import (
	"fmt"
	"time"
)

//在多协程并发环境下，我们常常会碰到以下两个问题。假设我们现在有 2 个协程，我们叫它们协程 A 和 B 。
//
//【问题1】如果协程 A 发生了 panic ，协程 B 是否会因为协程 A 的 panic 而挂掉？  会
//【问题2】如果协程 A 发生了 panic ，协程 B 是否能用 recover 捕获到协程 A 的 panic ？  不能

// 【问题1】如果协程 A 发生了 panic ，协程 B 是否会因为协程 A 的 panic 而挂掉？
func main() {

	go func() {
		for {
			fmt.Println("协程 AAA")
		}
	}()

	// 协程B
	go func() {

		// 协助 defer 增加异常捕获 不会再影响 协程A 的运行
		//  哪个协程发生 panic，就需要在哪个协程自身中 recover 。
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("协程 B err is %s", err)
			}
		}()
		time.Sleep(1 * time.Microsecond) // 确保 协程 A 先运行起来
		panic("协程 B panic")
	}()

	time.Sleep(10 * time.Second) // 充分等待协程 B 触发 panic 完成和协程 A 执行完毕
	fmt.Println("main end")

}
