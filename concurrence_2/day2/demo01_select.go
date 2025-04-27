package main

import "fmt"

/*
select:在能走通的case中随机选一条
何为走不通的路呢?--阻塞之路
所有的case全不通,就走default
*/
func main() {
	ch0 := make(chan int, 1)
	ch1 := make(chan int, 2)
	ch1 <- 3333
	ch2 := make(chan int, 3)
	ch2 <- 123
	ch2 <- 456

	select {
	case ch0 <- 123:
		fmt.Println("fuck [ch0 <- 123]")
	case x := <-ch1:
		fmt.Println("shit  [x := <= ch1]", x)
	case ch2 <- 123:
		fmt.Println("damn", "ch2 <- 123")
	default:
		fmt.Println("力拔山兮气盖世,时不利兮骓不逝...")
	}
}
