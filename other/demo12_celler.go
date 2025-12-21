package main

import (
	"fmt"
	"runtime"
)

func main() {
	foo()
}

func foo() {
	pcs := callers()
	for _, pc := range pcs {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		fmt.Printf("Function: %s, File: %s, Line: %d\n", fn.Name(), file, line)
	}
}

func callers() []uintptr {
	// 定义一个长度为 32 的 uintptr 类型数组 pcs，用于存储程序计数器的值
	var pcs [32]uintptr

	// 使用 runtime.Callers 函数获取调用栈中的程序计数器值
	// 第一个参数 3 表示跳过调用栈中的前 3 个调用帧，从第 4 个调用帧开始获取
	// 第二个参数 pcs[:] 是一个切片，指向上面定义的数组 pcs，用于存储获取到的程序计数器值
	// runtime.Callers 函数返回实际获取到的程序计数器值的数量
	l := runtime.Callers(3, pcs[:])

	// 返回一个切片，该切片包含了实际获取到的程序计数器值
	// 切片的长度为 l，即实际获取到的程序计数器值的数量
	return pcs[:l]
}
