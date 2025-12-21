package other

import (
	"fmt"
	"github.com/pkg/errors"
)

// 关于 GO 语言中 for 循环 错误跳出 break 和 return 的相关的知识点

//一、普通函数中的 for {} 循环
// 场景 1：只想跳出当前循环（继续执行后续代码）

func doSomething() error {
	return errors.New("asdasd")
	//return nil
}
func DemoFor() {
	for {
		err := doSomething()
		if err != nil {
			break // 跳出 for 循环，继续执行下面的代码
		}
	}
	fmt.Println("循环结束了，程序继续运行")
}

func Process() error {
	for {
		err := doSomething()
		if err != nil {
			return err // 直接退出函数，不再执行任何后续逻辑
		}
	}
	return nil
}

// 使用 break：仅退出循环，函数继续执行。
