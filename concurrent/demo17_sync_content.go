package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func Addnum2(a *int32, b int, deferFunc func()) {
	defer func() {
		deferFunc()
	}()

	for i :=0 ; ; i ++ {
		curNum := atomic.LoadInt32(a)
		newNum := curNum + 1

		time.Sleep(time.Millisecond * 200)
		if atomic.CompareAndSwapInt32(a, curNum, newNum) {
			fmt.Printf("number 当前值： %d [%d-%d]\n",*a,b,i)
			break
		} else {
			//fmt.Printf("The CAS operation failed. [%d-%d]\n", b, i)
		}
	}
}

func main() {
	total := 10
	var num int32
	fmt.Printf("number初始值: %d\n", num)
	fmt.Println("启动子协程...")
	ctx, cancelFunc := context.WithCancel(context.Background())
	for i :=0; i < total; i ++ {
		go Addnum2(&num, i , func() {
			if atomic.LoadInt32(&num) == int32(total) {
				cancelFunc()
			}
		})
	}

	//ctx, cancelFunc := context.WithTimeout(context.Background(), 10 * time.Second)
	//valueCtx := context.WithValue(ctx, "key", "value")
	//defer cancelFunc()
	//for i := 0; i < total; i++ {
	//	go Addnum2(&num, i, func() {
	//		if atomic.LoadInt32(&num) == int32(total) {
	//			fmt.Println("key:", valueCtx.Value("key"))
	//			cancelFunc()
	//		}
	//	})
	//}
	<- ctx.Done()
	fmt.Println("所有子协程执行完毕")
}

