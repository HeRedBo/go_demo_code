package main

import "sync"

func main() {
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
