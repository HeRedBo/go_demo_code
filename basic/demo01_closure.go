package main

import "fmt"

/*函数也是一种复合类型*/
func main011() {
	f := func(name string, hours int) {
		fmt.Printf("%s头领带队行军%d个小时\n", name, hours)
	}
	fmt.Printf("f的类型是%T，f的值是%v\n", f, f)
}

func main() {
	var p1, p2 int
	//获得武松和鲁达各自的【闭包内层函数】，闭包的作用是保存【各自的内层函数状态】
	f1 := GetDoTaskFunc()
	f2 := GetDoTaskFunc()

	//武松和鲁达犬牙交错地执行任务
	p2 = f2("鲁达", 24)
	p1 = f1("武松", 13)
	p2 = f2("鲁达", 24)
	p1 = f1("武松", 12)
	p1 = f1("武松", 1)

	//查看各自的状态，各自的任务状态被保存在各自的闭包中
	fmt.Println("二哥的进度是", p1)
	fmt.Println("鲁达的进度是", p2)
}

func GetDoTaskFunc() func(name string, hours int) (process int) {
	var progress int = 0
	f := func(name string, hours int) int {
		fmt.Printf("%s头领带队行军%d个小时\n", name, hours)
		progress += hours
		return progress
	}
	return f
}
