package func_program

import "fmt"

/*定义所有星期日*/
//const monday  = 1
//const tuesday  = 2
//const wednesday  = 3
//const thursday  = 4
//const friday  = 5
//const saturday  = 6
//const sunday  = 7

/*使用iota定义常量组*/
const (
	//iota的初始值为零
	Monday = iota + 1

	/*自动延用排头兵的表达式，但iota逐一递增*/
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

/*定义5大常任理事国*/
const (
	USA     = (iota + 1) * 1000
	China
	Russia
	Britain
	France
)

func PrintIoTaCount()  {
	//fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	fmt.Println(USA, China, Russia, Britain, France)
}