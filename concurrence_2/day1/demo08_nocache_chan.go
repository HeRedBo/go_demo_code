package main

import (
	"fmt"
	"time"
)

/*
无缓存的管道
没人读就永远写不进(阻塞)
没人写就永远读不出(阻塞)
*/
func main() {

	//创建没有缓存能力的管道
	//ch := make(chan int)
	ch := make(chan int, 0)
	//往管道里写数据
	go func() {
		//没人读管道，就永远写不进去
		ch <- 123
		fmt.Println("数据已写入")
	}()

	go func() {
		time.Sleep(10 * time.Second)
		//没人写就读不出来
		x := <-ch
		fmt.Println("数据已读出", x)
	}()

	time.Sleep(11 * time.Second)
	fmt.Println("Game Over!")

}
