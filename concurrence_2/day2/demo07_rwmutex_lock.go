package main

import (
	"fmt"
	"sync"
	"time"
)

var rwm sync.RWMutex
var wg07 sync.WaitGroup

func main701() {
	var rwMutex sync.RWMutex
	// 锁定为只读模式（可以多路锁定）
	rwMutex.RLock()
	// 解锁只读模式
	rwMutex.RUnlock()
	// 锁定为只写模式（只有一路能锁定）
	rwMutex.Lock()
	// 解锁只写模式
	rwMutex.Unlock()
}


func read(i int) {
	rwm.RLock()
	fmt.Println(i,"reading...")
	<-time.After(5 * time.Second)
	fmt.Println(i,"read over")

	rwm.RUnlock()
	wg07.Done()
}


func write(i int) {
	//锁定为只写模式(全局只有当前一条协程能写,其余协程无论读写都被阻塞)
	rwm.Lock()

	fmt.Println(i,"writing...")
	<-time.After(5 * time.Second)
	fmt.Println(i,"write over")
	//解锁只读模式,引起其它协程(不分读写)对锁的哄抢
	rwm.Unlock()
	wg07.Done()
}