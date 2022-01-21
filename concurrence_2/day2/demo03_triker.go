package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个1秒为周期的秒表
	ticker := time.NewTicker(1 * time.Second)

	var i int
	for {
		x := <-ticker.C
		fmt.Print("\r",x)
		i ++
		//10秒后停止计时并退出
		if i > 9 {
			//停掉秒表,会导致ticker.C永远无法读出数据,执着要读会导致死锁(deadlock)
			ticker.Stop()
			break
		}
	}
	fmt.Println("计时结束")

}
