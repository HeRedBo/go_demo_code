package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		//主协程创建子协程：瞬间完成
		go func(index int) {
			//子协程执行任务，子协程需要一定的初始化时间
			task("子协程" + strconv.Itoa(index))
		}(i)
	}

	time.Sleep(6 * time.Second)
	fmt.Println("GameOver!")

}

func hello(index int) {
	//子协程执行任务，子协程需要一定的初始化时间
	task("子协程" + strconv.Itoa(index))
}

func task(name string) {
	for i := 0; i < 2; i++ {
		if name == "子协程5" {
			//孔融让梨,让别人先执行(优先级排到最后,但并不是一点机会都没有)
			runtime.Gosched()
		}
		fmt.Println(name, i)
	}
}
