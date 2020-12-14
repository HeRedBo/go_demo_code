package basic_grammar

import "fmt"

func Slice1()  {
	var slice []string = []string{"a", "b", "c"}
	fmt.Printf("slice的类型是%T，值是%v\n",slice,slice)
}

/**
数组切片
 */
func ArraySlice()  {
	// 先定义一个数组
	months := [...]string{"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December"}

	// 基于数组创建数组切片
	q2 := months[3:6]    // 第二季度
	summer := months[5:8]  // 夏季
	all := months[:]  // all
	firsthalf := months[:6]
	secondhalf := months[6:]

	fmt.Println(q2)
	fmt.Println(summer)
	fmt.Println(all)
	fmt.Println(firsthalf)
	fmt.Println(secondhalf)
}
