package main

import (
	"fmt"
	"strconv"
	"time"
)

//模拟商店，能缓存10件商品
var chanShop chan string = make(chan string, 10)

//模拟消费指令，无缓存能力，生产者生产完毕，写入消费指令，在此之前，消费者阻塞等待消费指令
var chanTel chan bool = make(chan bool)

//主协程去死管道
var chanGoDie chan string  = make(chan string, 0)


func main() {
	//生产者协程
	go Producer()
	//消费者协程
	go Consumer()

	//主协程不能死
	<- chanGoDie
	fmt.Println("GAME OVER!")
}

/* 生产者 */
func Producer() {
	for i := 0 ; i < 10; i ++  {
		//生产商品
		product := "产品" +  "[" + strconv.Itoa(i) + "]"+ strconv.Itoa(time.Now().Nanosecond())
		//丢入商店管道
		chanShop <- product
		fmt.Println("生产了产品", product)
		time.Sleep(time.Second)
	}
	close(chanShop)
	fmt.Println("生产完毕")
	fmt.Println()

	//发送吃的指令给消费者（人家一直阻塞等待这个指令）
	chanTel <- true
}

func Consumer() {
	//阻塞等待吃的指令，此处chanTel中数据内容无关紧要，要的是其阻塞效果
	<-chanTel
	//阻塞等待吃的指令，此处chanTel中数据内容无关紧要，要的是其阻塞效果
	for product := range chanShop {
		fmt.Println("消费了产品", product)
		fmt.Println()
	}
	fmt.Println("商品已全部被消费")

	//所有子协程任务都已结束，通知主协程去死
	chanGoDie <- "去死"
}

