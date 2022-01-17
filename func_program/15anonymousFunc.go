package func_program

import (
	"fmt"
	"time"
)

func DelayDemo1(){
	defer fmt.Println("我是倒数第一")

	defer func() {
		fmt.Println("我是倒数第二")
		fmt.Println("我是倒数第三")
		fmt.Println("我是倒数第四")
	}()

	fmt.Println("文能复习划重点")
	fmt.Println("武能装逼带你飞")
	fmt.Println("上九天揽月")
	fmt.Println("下五洋捉鳖")
}

func DelayDemo2() {

	//go中的代码跑在【子协程】
	go func() {
		fmt.Println("春困")
		time.Sleep(1 * time.Second)
		fmt.Println("秋乏")
		time.Sleep(1 * time.Second)
		fmt.Println("夏打瞌睡")
		time.Sleep(1 * time.Second)
		fmt.Println("冬眠")
		time.Sleep(1 * time.Second)
	}()

	//以下代码跑在【主协程】
	fmt.Println("文能复习划重点")
	time.Sleep(1 * time.Second)
	fmt.Println("武能装逼带你飞")
	time.Sleep(1 * time.Second)
	fmt.Println("上九天揽月")
	time.Sleep(1 * time.Second)
	fmt.Println("下五洋捉鳖")
	time.Sleep(2 * time.Second)
}


