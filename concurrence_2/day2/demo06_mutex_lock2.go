package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	money int
	mt sync.Mutex
}

// 存钱
func (a Account) Save(amount int) int {
	// 抢锁
	a.mt.Lock()
	fmt.Println("正在存钱...")
	<-time.After(10 * time.Second)
	a.money += amount
	fmt.Println("存钱结束")
	a.mt.Unlock()
	return a.money
}




func main() {

}
