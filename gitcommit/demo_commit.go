package gitcommit

import (
	"fmt"

	"github.com/gookit/goutil/dump"
)

func main() {
	DemoAction1()
	DemoAction2()

	// 第三次内容提交内容
	DemoAction3()

	// 第三次 最佳新的方法
	DemoAction5(5, 6)
}

func DemoAction1() {
	fmt.Println("Demo Action 1")
	dump.P("this is add  DemoAction1 Content")
}

func DemoAction2() {
	fmt.Println("Demo Action 2")
	dump.P("this is add  DemoAction2 Content")
}

func DemoAction3() {
	fmt.Println("Demo Action 3")
	dump.P(DemoAction4(1, 3))
}

func DemoAction4(num1 int, num2 int) int {
	return num1 + num2
}

func DemoAction5(num1 int, num2 int) int {
	dump.Println(num1 + num2)
	dump.Println("DemoAction5 content")
	return num1 + num2
}
