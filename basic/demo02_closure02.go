package main

import "fmt"

func main() {
	tf1 := GetTaskFunc("武松")
	tf2 := GetTaskFunc("李逵")

	p1 := tf1(13)
	p2 := tf2(31)
	p1 = tf1(12)
	p2 = tf2(2)

	fmt.Println("武松的进度=", p1)
	fmt.Println("李逵的进度=", p2)
}

func GetTaskFunc(name string) func(hours int) (process int) {
	var progress int = 0
	ExecuteTask := func(hours int) int {
		fmt.Printf("%s头领带队行军%d个小时\n", name, hours)
		progress += hours
		return progress
	}
	return ExecuteTask
}
