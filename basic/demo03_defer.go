package main

import "fmt"

func main() {
	//if (true) {
	//	defer fmt.Println("1")
	//} else {
	//	defer fmt.Println("2")
	//}
	//fmt.Println("3")

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