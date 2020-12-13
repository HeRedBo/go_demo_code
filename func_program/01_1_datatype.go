package func_program

import "fmt"

/*
	%T 类型占位符
	%v 值占位符与

	%d 整型占位符
	%f 浮点型占位符
	%c 字符占位符
	%s 字符串占位符
*/


func DataType () {
	var v1 = 123
	fmt.Printf("v1的类型是%T,值是%d\n",v1,v1)

	var v2 int = 123
	fmt.Printf("v2的类型是%T\n",v2)


	var v3 = 123.0
	fmt.Printf("v3的类型是%T，值是%f\n",v3,v3)

	var v4 float32 = 123.0
	fmt.Printf("v4的类型是%T\n",v4)

	var v5 = "你妹"
	fmt.Printf("v5的类型是%T，值是%s\n",v5,v5)

	var v6 = '俏'
	fmt.Printf("v6的类型是%T\n",v6)
	fmt.Printf("v6的值是%v\n",v6)
	fmt.Printf("v6的字符是%c\n",v6)
	fmt.Printf("20431的字符是%c\n",20431)

	var v7 = (100==(40+60))
	fmt.Printf("v7的类型是%T,值是%v\n",v7,v7)

	var v8 = ('俏'==20431)
	fmt.Printf("v8的类型是%T,值是%v\n",v8,v8)

	fmt.Printf("20431的字符形式是%c\n",20431)
	fmt.Printf("俏的数字形式是%d\n",'俏')
	fmt.Printf("俏在字符集UTF8中的序号是%d\n",'俏')
}