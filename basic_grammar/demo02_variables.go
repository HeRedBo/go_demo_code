package basic_grammar

import "fmt"

/**
 * 常量、变量、表达式
 */
func Main001() {
	//声明常量（constant）pi，赋值为3.14
	const pi = 3.14

	//声明变量（variable）r，赋值为1.0
	var r = 1.0

	//声明变量area并赋值为（pi*r*r）的结果
	area := pi * r * r

	//打印结果
	fmt.Println("圆的面积是", area)

	//cannot assign to constant 不能赋值给常量
	// pi = 3.1415926

	//对变量r重新赋值
	r = 2.0
	//对area变量重新赋值
	area = pi * r * r
	//打印结果
	fmt.Println("圆的面积是", area)
}

/*
一次性声明多个常量
一次性声明多个变量
*/
func Main002() {
	fmt.Println("我又活过来了~")
	//main021()

	//一次性声明多个常量
	const (
		lightSpeed  = 30 * 10000
		earthCircle = 40000
	)
	fmt.Println("光速常量是", lightSpeed, "地球周长常量是", earthCircle)

	//一次性声明多个变量
	var (
		age    = 20
		height = 175
		weight = 60
	)
	fmt.Println("年龄变量是", age, "身高变量是", height, "体重变量是", weight)

	/*修改变量的值*/
	age    = 30
	height = 175
	weight = 80
	fmt.Println("年龄变量是", age, "身高变量是", height, "体重变量是", weight)
}


/*花式定义变量*/
func Main003() {

	//带var定义，一个或多个，声明类型或不声明类型
	var a int = 2
	var b = 3
	var d, e int = 5, 6
	var f, g = 7, 8

	//冒等，一个或多个，只能动态检测类型
	c := 4
	h, i := 9, 10

	fmt.Println(a, b, c, d, e, f, g, h, i)
}