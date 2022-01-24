package main

import (
	"fmt"
	"sync"
	"time"
)

func main0401() {
	// 创建任务等等组
	var wg sync.WaitGroup

	// 向等的组中添加任务
	wg.Add(1)
	wg.Add(1)
	wg.Add(1)

	// 从等待组中抹掉任务
	wg.Done()
	wg.Done()
	wg.Done()

	//阻塞等待至等待组中的任务数为0
	wg.Wait()

}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for i:=0; i < 5; i ++  {
			fmt.Println("子协程1", i)
			<- time.After( 1 * time.Second)
		}
		fmt.Println("子协程1结束任务")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		var i int
		ticker := time.NewTicker(1 * time.Second)
		for {
			<-ticker.C
			i++
			fmt.Println("子协程2","秒表:",i)
			if i > 9 {
				break
			}
		}
		fmt.Println("子协程2结束任务")
		wg.Done()
	}()

	//等待组阻塞等待至记录清零为止
	wg.Wait()
	fmt.Println("END")


}