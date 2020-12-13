package func_program

import (
	"fmt"
)

func SwitchOneByOne()  {

	fmt.Println("请输入今天星期几？1-7")

	var userinput string
	fmt.Scan(&userinput)

	switch userinput {
	case  "1" :
		fmt.Println("垂头丧气，意犹未尽")
	case "2":
		fmt.Println("工作状态")
	case "3":
		fmt.Println("工作状态")
	case "4":
		fmt.Println("黎明前的黑暗")
	case "5":
		fmt.Println("胜利大逃亡")
	case "6":
		fmt.Println("一起去浪")
	case "7":
		fmt.Println("一起去浪")
	default:
		fmt.Println("对不起您的智商余额不足请尽快充值！")
	}

}

// /*合并相同的case*/


func SwitchMerge() {
	fmt.Println("请输入今天星期几？1-7")

	var userinput string
	fmt.Scan(&userinput)

	switch userinput {
	case "1":
		fmt.Println("垂头丧气，意犹未尽")
	case "2", "3":
		fmt.Println("工作状态")
	case "4":
		fmt.Println("黎明前的黑暗")
	case "5":
		fmt.Println("胜利大逃亡")
	case "6", "7":
		fmt.Println("一起去浪")
	default:
		fmt.Println("对不起您的智商余额不足请尽快充值！")
	}
}

func Switch3() {
	fmt.Println("请输入今天星期几？1-7")
	var userinput string
	fmt.Scan(&userinput)
	//fmt.Printf("userinput的类型是%T,值是%d\n",userinput,userinput)
	fmt.Printf("userinput的类型是%T,值是%d\n",userinput,userinput)
	//var age int
	//age, _  := strconv.Atoi(userinput)
	//fmt.Printf("userinput的类型是%T,值是%d\n",age,age)
	return
	//fmt.Println(age)
	var age int
	age = 35

	switch {
	case age < 18:
		fmt.Println("你是一个少年")
	case age >= 18 && age < 36:
		fmt.Println("你是一个青年人")
	case age >= 36 && age < 60:
		fmt.Println("你是一个中年人")
	case age >= 60 && age < 100:
		fmt.Println("你是一个老年人")
	default:
		fmt.Println("你是一个传奇")
	}
}