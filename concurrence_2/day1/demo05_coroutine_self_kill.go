package main

import (
	"fmt"
	"runtime"
	"time"
)

func task001 () {
	//4://因协程被杀死而提前触发
	defer fmt.Println("这是亮这个月的党费")

	//2
	fmt.Println("平生我自知")
	//3
	fmt.Println("草堂秋睡足")

	//杀死所在的协程
	runtime.Goexit()

	//因协程被杀死而执行不到
	fmt.Println("窗外日迟迟")

}
func main() {
	go func() {
		//1
		fmt.Println("大梦谁先觉")

		//任务内杀死当前协程
		task001()

		//因协程被杀死而执行不到
		fmt.Println("--诸葛亮")
	}()

	time.Sleep(time.Second)
}
