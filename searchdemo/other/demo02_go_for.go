package other

import (
	"context"
	"fmt"
	"time"
)

/**
二、被协程（goroutine）包裹的 for {} 循环
这是关键点！在 goroutine 中，break 和 return 的行为不同：

❗ 重要原则：
break：只能跳出当前 goroutine 内的循环
return：只能退出当前 goroutine 的函数，不会影响主 goroutine
而且：你无法从外部“强制”让一个 goroutine 停止（Go 没有 thread.kill()）。
*/

/**
✅ 正确做法：使用 context 或 channel 信号 来协调退出
示例 1：使用 context 控制 goroutine 退出
*/

func GoForWithCentext() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("收到退出信号，goroutine 退出")
				return // 退出 goroutine
			default:
				err := doSomething()
				if err != nil {
					fmt.Println("发生错误，主动退出")
					cancel() // 可选：通知其他 goroutine
					return   // 👈 在 goroutine 中用 return 退出自己
				}
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	cancel()                           // 主动触发退出
	time.Sleep(100 * time.Millisecond) // 等待 goroutine 退出
}

func GoForWithChannel() {
	done := make(chan bool)
	go func() {
		defer close(done)
		for {
			err := doSomething()
			if err != nil {
				fmt.Println("goroutine 因错误退出")
				return // 👈 退出当前 goroutine
			}
			// 模拟工作
			time.Sleep(200 * time.Millisecond)
		}
	}()

	// 等待 goroutine 结束
	<-done
	fmt.Println("主程序继续")
}
