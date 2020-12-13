package func_program

import "fmt"

func OneIf() {
	fmt.Println("请输入相亲者姓名：")

	//声明一个字符串变量name
	var name string
	//接收用户输入，存入name变量所在的地址
	fmt.Scan(&name)
	if name == "楠哥"{
		fmt.Println("fuck off!")

		//结束当前函数的执行
		return
	}
	fmt.Println("进入人民公园相亲现场！")
}


func DoubleIf() {
	fmt.Println("请输入相亲者姓名：")
	var name string
	fmt.Scan(&name)

	if name == "杰哥" {
		fmt.Println("安排你和春哥见面")
	} else {

	}

	fmt.Println("GAME OVER!")
}

func MultiIf() {
	fmt.Println("请输入相亲者姓名：")
	var name string
	fmt.Scan(&name)
	if name == "春哥" {
		fmt.Println("安排了你从东门进入")
	} else if name == "明哥" {
		fmt.Println("安排了您和陈一发儿的会面")
	}  else {
		fmt.Println("请等候调度...")
	}
	fmt.Println("GAME OVER")
}