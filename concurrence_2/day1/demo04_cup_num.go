package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("可用的CPU核数为",runtime.NumCPU())

	//将何用核数设置为1,并返回先前的设置(系统默认设置为全部可用)
	fmt.Println("设置CPU的可用核数为2,先前的设置为",runtime.GOMAXPROCS(2))

	//打印
	fmt.Println("可用的CPU核数为6,先前的设置为",runtime.GOMAXPROCS(6))
}
