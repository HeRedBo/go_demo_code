package main

import (
	"fmt"
)

func main091() {
	ch := make(chan int)
	go func() {
		x := <-ch
		fmt.Println(x)
	}()
	ch <- 123
	//x := <-ch
	//fmt.Println(x)
}

/*[开塞露协程]开晚了*/
func main092() {
	ch := make(chan int)
	go func() {
		x := <-ch //这里虽然创建了子go程用来读出数据，但是上面会一直阻塞运行不到下面
		fmt.Println(x)
	}()
	ch <- 666 //这里一直阻塞，运行不到下面
}

func main() {
	chHusband := make(chan int)
	chWife := make(chan int)

	// 老公
	go func() {
		for {
			select {
			case <-chHusband:
				chWife <- 123
				fmt.Println("老公：我已经发给你发了123元红包了")
			case chWife <- 1:
				fmt.Println("老公:我已经给你发送1云红包了")
			}
		}
	}()

	//老婆
	select {
	//只要我有钱我就给你发红包
	case <-chWife:
		chHusband <- 123
		fmt.Println("老婆：我已经给你发123元红包了")
	}

	fmt.Println("THE END")
}
