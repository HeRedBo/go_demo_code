package main

import "fmt"

func main() {

	ch := make(chan int, 4)
	fmt.Println("元素个数为", len(ch), ",缓存能力为", cap(ch))

	ch <- 123
	fmt.Println("元素个数为", len(ch), ",缓存能力为", cap(ch))

	ch <- 123
	fmt.Println("元素个数为", len(ch), ",缓存能力为", cap(ch))
	ch <- 123
	fmt.Println("元素个数为", len(ch), ",缓存能力为", cap(ch))

	//阻塞:
	//主协程被永远阻塞造成死锁:fatal error: all goroutines are asleep - deadlock!
	ch <- 123
	fmt.Println("元素个数为", len(ch), ",缓存能力为", cap(ch))
}
