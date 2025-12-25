package gocaller

import (
	"fmt"
	"github.com/gookit/goutil/dump"
	"runtime"
	"strings"
)

func GoCaller() {
	// 获取调用者的信息
	pc, file, line, ok := runtime.Caller(0)
	if ok {
		dump.P(fmt.Sprintf("文件: %s, 行号: %d\n", file, line))
		dump.P(pc) // 调用函数的程序计数器（通常不使用，但可用于获取函数名）
		// 通过pc获取函数名
		fn := runtime.FuncForPC(pc)
		dump.P(fmt.Sprintf("函数名: %s\n", fn.Name()))
	}
}

func Log(msg string) {
	// 跳过1层，获取调用Log的函数的信息
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		// 提取文件名（去掉路径）
		fileName := file[strings.LastIndex(file, "/")+1:]
		// 获取函数名
		funcName := runtime.FuncForPC(pc).Name()
		dump.P(fmt.Sprintf("[%s:%d] [%s] %s\n", fileName, line, funcName, msg))
		//fmt.Printf("[%s:%d] [%s] %s\n", fileName, line, funcName, msg)
	} else {
		dump.P(msg)
		//fmt.Println(msg)
	}
}

func Foo() {
	Log("这是一条日志")
}

// DemoCaller 1. 获取调用者信息
func DemoCaller() {
	// skip=0: 获取当前函数(demoCaller)的信息
	// skip=1: 获取调用者(main)的信息
	// skip=2: 获取更上层的调用者信息
	for i := 0; i <= 2; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if ok {
			// 通过pc获取函数名
			fn := runtime.FuncForPC(pc)
			dump.P(fmt.Sprintf("skip=%d: 文件=%s, 行号=%d, 函数=%s\n",
				i, file, line, fn.Name()))
		} else {
			dump.Println(fmt.Sprintf("skip=%d: 无法获取信息\n", i))
		}
	}
}
