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
func (a *Account) Save(amount int) int {
	// 抢锁
	a.mt.Lock()
	fmt.Println("正在存钱...")
	<-time.After(3 * time.Second)
	a.money += amount
	fmt.Println("存钱结束")
	a.mt.Unlock()
	return a.money
}

/* 取钱
 */
func (a *Account) GetMoney(amount int) int {
	// 抢钱
	a.mt.Lock()
	fmt.Println("正在取钱")
	<-time.After(3 * time.Second)
	a.money -= amount
	fmt.Println("取钱结束")
	a.mt.Unlock()
	return amount
}

func (a *Account) Query() int {
	fmt.Println("当前余额",a.money)
	return a.money
}


/* 数据初始化 */
var waitGroup sync.WaitGroup
var account *Account

func init() {
	account = new(Account)
	account.money = 10000
}

func main() {

	waitGroup.Add(1)
	go func() {
		account.Save(500)
		waitGroup.Done()
	}()

	waitGroup.Add(1)
	go func() {
		account.GetMoney(500)
		waitGroup.Done()
	}()

	waitGroup.Add(1)
	go func() {
		account.Query()
		waitGroup.Done()
	}()
	waitGroup.Wait()
	fmt.Println("END")
}
