package main

import (
	"fmt"
	"helloGO/func_program"
)

func main ()  {
	fmt.Println("hello World")

	//func_program.DataType();

	//func_program.Hello()
	//func_program.VarConst()
	//
	//func_program.VarConst2()
	//
	//func_program.VarConst3()
	//
	//func_program.UpperFunc()

	//func_program.PrintIoTaCount()

	// 基本的数学运算
	//func_program.BaseMath()

	// 高级数学运算
	//func_program.HighMath()

	// 整型和浮点型的互化
	//func_program.Int22Float()

	// 自加运算
	// func_program.SelfAutoMath()
	// 比较运算
	//func_program.Compare()

	//逻辑运算
	//func_program.Logic()

	// if 选择结构
	// func_program.OneIf();

	// if 双分支结构
	//func_program.DoubleIf()

	// //多分支结构
	//func_program.MultiIf()

	// /*条件表达式的结果落在多个不同的孤立值上*/
	//func_program.SwitchOneByOne()


	//func_program.SwitchMerge()

	//func_program.Switch3()

	//func_program.GoTo()

	//func_program.Defer1()
	//func_program.Defer2()
	//func_program.Defer4()
	// hello world
	//basic_grammar.Hello()

	// 匿名函数：
	// func_program.DelayDemo1()
	//func_program.DelayDemo2()

	// 递归
	func_program.TimeIt(func_program.GetSum,10000000)
	func_program.TimeIt(func_program.ReF,10000000)
	//fmt.Printf(func_program.GetSum(1000000))
	fmt.Printf("递归结果值是%f\n",func_program.GetSum(1000000))
	fmt.Printf("递归结果值是%f\n",func_program.ReF(1000000))
	//func_program.TimeIt(func_program.ReF,100000)

	// 变量表达式
	//basic_grammar.Main001()
	//basic_grammar.Main002()
	//basic_grammar.Main003()

	// 基本数据类型
	//basic_grammar.DateType()
	// 指针
	//basic_grammar.Pointer()

	// 数组
	//basic_grammar.Array1()
	// 9x9 乘法表
	// basic_grammar.Math9x9()
	// 切片
	//basic_grammar.SliceBase()
	//basic_grammar.SliceAppend()
	//basic_grammar.SliceAutoAppend()
	//basic_grammar.Slice1()
	//basic_grammar.ArraySlice()

	// 递归处理相关数据


	// 异常处理
	//exception_handling.ErrorPanic()

	// 恐慌与处理


}
