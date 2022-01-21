package main

import (
	"fmt"
	"time"
)

/** 写入阻塞 */
func main901() {
	ch := make(chan int, 3)
	go func() {
		ch <- 1
		fmt.Println("写入1")
		fmt.Println("写入1")
		ch <- 2
		fmt.Println("写入2")
		ch <- 3
		fmt.Println("写入3")

		//缓存能力已耗尽,没人读就再也写不进去了(阻塞)
		ch <- 4
		fmt.Println("写入4")
	}()

	time.Sleep(1 * time.Second)
}

/*读取阻塞*/
func main902() {
	ch := make(chan int, 3)
	//close(ch)
	go func () {
		x := <-ch
		fmt.Println("读到", x)

		x = <-ch
		fmt.Println("读到", x)

		x = <-ch
		fmt.Println("读到", x)

		x = <-ch
		fmt.Println("读到", x)
	}()
	time.Sleep(5 * time.Second)
}


/*管道的元素个数和缓存能力*/
func main() {
	ch := make(chan int, 3)
	fmt.Println("元素个数为",len(ch),",缓存能力为",cap(ch))

	ch <- 123
	fmt.Println("元素个数为",len(ch),",缓存能力为",cap(ch))

	ch <- 123
	fmt.Println("元素个数为",len(ch),",缓存能力为",cap(ch))
	ch <- 123
	fmt.Println("元素个数为",len(ch),",缓存能力为",cap(ch))

	//阻塞:
	//主协程被永远阻塞造成死锁:fatal error: all goroutines are asleep - deadlock!
	ch <- 123
	fmt.Println("元素个数为",len(ch),",缓存能力为",cap(ch))

}
