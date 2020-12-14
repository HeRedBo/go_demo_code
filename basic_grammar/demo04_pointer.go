package basic_grammar

import "fmt"

//指针就是地址
//&value 对值取地址
//*ptr 对地址取值
func Pointer() {

	//声明变量a时，系统开辟了一块内存【地址】，里面存的【值】是123
	var a int = 123
	fmt.Printf("a的类型是%T\n", a)//int
	fmt.Printf("a的值是%v\n", a)//123
	fmt.Printf("a的地址是%p\n", &a)//0x...

	//&a取变量a的地址
	aPointer := &a
	fmt.Printf("aPointer的类型是%T\n",aPointer)//*int

	//将aPointer指向的地址中的值修改为456
	*aPointer = 456
	fmt.Println("*aPointer=",*aPointer)//456
	//a的值就变成了456
	fmt.Println("a=",a)
}