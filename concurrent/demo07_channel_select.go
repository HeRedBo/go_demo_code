package main

import (
	"fmt"
	"math/rand"
)

func main() {
	chs := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}

	index := rand.Intn(3) // 随机生成 0-2 之间的数字
	fmt.Printf("随机索引/数值： %d\n", index)
	chs[index] <- index // 向同道发送随机数字

	// 哪一个通道中有值，哪个对应的分支就会被执行
	select {
	case <-chs[0]:
		fmt.Println("第一个条件分支被选中")
	case <-chs[1]:
		fmt.Println("第二个条件分支被选中")
	case num := <-chs[2]:
		fmt.Println("第三个条件分支被选中:", num)
	default:
		fmt.Println("没有分支被选中")
	}

}
