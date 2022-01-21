package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	go wildman()

	time.Sleep(5 * time.Second)

	//主协程自杀=>子协程失去约束
	fmt.Println("啊，我要自裁以谢天下")
	runtime.Goexit()
	fmt.Println("主协程：吾去也!")
}



func wildman() {
	for {
		fmt.Println("我是野人,我爱自由,我讨厌主人")
		time.Sleep(time.Second)
	}
}
