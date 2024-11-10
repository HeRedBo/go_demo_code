package main

import (
	"fmt"
	"sync"
)

// https://zhuanlan.zhihu.com/p/374464199
// 相关实例代码

func A (i int) {
	fmt.Printf("我是A ：%d",i)
	fmt.Println()
	panic("崩溃了")

}
func main() {
	var wg sync.WaitGroup
	fmt.Println("我是 main")
	wg.Add(1)

	go func(i int) {

		defer func() {
			if err := recover() ; err != nil {
				fmt.Println("外部函数获取一层测试恢复",err)
			}
		}()
		defer func() {
			wg.Done()
		}()
		A(i)
	}(1)
	wg.Wait()
	fmt.Println("执行完了")
}
