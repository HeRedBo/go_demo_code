package main

import "fmt"

func main() {
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
