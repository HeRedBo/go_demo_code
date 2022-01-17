package basic_grammar

import "fmt"

/*
切片相当于长度可以动态扩张的数组
*/

//array[start:end]从数组身上截取下标为[start,end)片段，形成切片
//start代表开始下标，不写默认代表从头开始切
//end代表结束下标（本身不被包含），不写默认截取到末尾

func SliceBase() {
	var array = [10]int{0, 11, 22, 33, 44, 55, 66, 77, 88, 99}

	//含头不含尾，从array的第0项截取到第9项
	slice := array[0:10]
	fmt.Printf("array的类型是%T，值是%v\n", array, array) //[10]int,[0 11 22 33 44 55 66 77 88 99]
	fmt.Printf("slice的类型是%T，值是%v\n", slice, slice) //[]int,[0 11 22 33 44 55 66 77 88 99]

	slice = array[0:5]
	fmt.Printf("slice的类型是%T，值是%v\n", slice, slice) //

	slice = array[2:5]
	fmt.Printf("slice的类型是%T，值是%v\n", slice, slice)

	slice = array[:5]
	fmt.Printf("slice的类型是%T，值是%v\n", slice, slice)

	slice = array[2:]
	fmt.Printf("slice的类型是%T，值是%v\n", slice, slice)

	slice = array[:]
	fmt.Printf("slice的类型是%T，值是%v\n", slice, slice)

	//也可以对切片进行截取，得到切片
	sonSlice := slice[:]
	fmt.Printf("sonSlice的类型是%T，值是%v\n", sonSlice, sonSlice)

}

//向切片中追加元素
//遍历切片
func SliceAppend() {
	//初始化一个没有元素的整型切片
	var slice []int = []int{}
	fmt.Printf("类型是%T\n", slice)                     //[]int
	fmt.Printf("切片的长度是%d,值是%v\n", len(slice), slice) //0,[]

	slice = append(slice, 0)
	fmt.Printf("切片的长度是%d,值是%v\n", len(slice), slice) //1,[0]

	slice = append(slice, 11, 22, 33)
	fmt.Printf("切片的长度是%d,值是%v\n", len(slice), slice) //4,[0 11 22 33]

	// 切片遍历
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice的第%v个元素的值是%v\n", i, slice[i])
	}

	for index, value := range slice {
		fmt.Printf("slice的第%d个元素是%d\n", index, value)
	}

}

//cap(slice)获得切片的容量
//创建之初，容量等于长度
//扩张时，一旦容量无法满足需要，就以翻倍的策略扩容
func SliceAutoAppend() {
	var slice = []int{1, 2, 3} //len=3,cap=3
	slice = append(slice, 4)   //len=4,cap=6
	fmt.Printf("slice的长度是%d，容量是%d\n", len(slice), cap(slice))
}



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
