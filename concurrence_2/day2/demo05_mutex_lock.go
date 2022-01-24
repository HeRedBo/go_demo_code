package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main0501() {
	var money = 2000
	for i :=0 ; i < 10; i ++ {
		for i :=0 ; i < 10; i ++ {
				for j := 0; j < 1000; j ++   {
					money += 1
				}
		}
	}
	fmt.Println("最终金额1", money)

	var money2 =  2000
	for i :=0 ; i < 10; i ++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j ++   {
				money2 += 1
			}
			wg.Done();
		}()
	}
	wg.Wait()
	fmt.Println("最终金额2", money2)
}


func main() {
	var money = 2000

	// 一只麦
	var mt sync.Mutex

	wg.Add(1)
	go func() {
		fmt.Println("小张开始抢麦")

		mt.Lock()
		fmt.Println("小张抢到了麦，小张开始嗨")
		money -= 300
		<- time.After(10 * time.Second)

		//曲罢弃麦,其它人再去哄抢
		mt.Unlock()
		fmt.Println("小张姐弃麦")
		wg.Done()
	}()



	wg.Add(1)
	go func() {
		fmt.Println("小李开始抢麦")

		mt.Lock()
		fmt.Println("小李抢到了麦，小李开始嗨")
		money -= 300
		<- time.After(10 * time.Second)

		//曲罢弃麦,其它人再去哄抢
		mt.Unlock()
		fmt.Println("小李弃麦")
		wg.Done()
	}()

	/*不抢麦的人不会被阻塞*/

	wg.Add(1)
	go func() {
		fmt.Println("abcdefg")
		<-time.After(1*time.Second)
		fmt.Println("hijklmn")
		<-time.After(1*time.Second)
		fmt.Println("opqrst")
		<-time.After(1*time.Second)
		fmt.Println("uvwxyz")
		wg.Done()
	}()
	wg.Wait()
}