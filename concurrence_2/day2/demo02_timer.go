package main

import (
	"fmt"
	"time"
)

func main201() {
	//创建3秒钟的定时器
	timer := time.NewTimer(3 * time.Second)
	fmt.Println("定时器创建完毕!")
	fmt.Println(time.Now())
	//阻塞3秒,后才能读出当前时间
	x := <-timer.C
	fmt.Println(x)
}

/*固定时长定时器的第二种写法*/
func main() {
	fmt.Println(time.Now())
	x := <- time.After(3*time.Second)
	//x := <- time.NewTimer(3*time.Second).C
	fmt.Println(x)
}
