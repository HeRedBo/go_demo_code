package basic_grammar

import (
	"fmt"
)

/**
固定长度，固定类型的数据容器
 */

func Array1() {
	//var array = [5]int{3,5,6}
	//fmt.Printf("array 的类型 是%T，值是%v\n",array,array)
	// array 的类型 是[5]int，值是[3 5 6 0 0]

	array := [...]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8}
	fmt.Println(array)

	fmt.Println("数组的长度是", len(array)) //12

	//根据下标对元素进行访问
	fmt.Println("数组的第一个元素是", array[0]) //3
	fmt.Println("数组的第六个元素是", array[5]) //9
	array[0] = 333
	fmt.Println("数组的第一个元素是", array[11]) //333
	////下标越界错误：index out of range


	//遍历数组1
	//for i := 0; i < len(array); i++ {
	//	fmt.Printf("数组的第%d个元素是%d\n",i,array[i])
	//}


	//遍历数组2
	//index是下标，value是值
	for index, value := range array {
		fmt.Println(index, value)
	}


}
