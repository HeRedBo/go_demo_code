package main

import (
	"fmt"
	"strconv"
	"time"
)

/*生产者-消费者模型1:生产-消费*/
func main101() {
	//商店管道:生产者和消费者传递产品的纽带
	chanShop := make(chan string, 100)

	//生产者协程
	go func() {
		for {
			product := strconv.Itoa(time.Now().Nanosecond())
			chanShop <- "产品" + product
			fmt.Println("生产了产品： ",product)
			time.Sleep(time.Second)
		}

	}()

	// 消费这协程
	go func() {
		for {
			product := <-chanShop
			fmt.Println("消费了产品： ",product)
			time.Sleep(time.Second)
		}
	}()
	// 保持主协程正常运行
	for {
		time.Sleep(time.Second)
	}

}

/*生产者-消费者模型2:生产-物流-消费*/
func main() {
	chanStorage := make(chan string , 10 )
	chanShop := make(chan string , 10)

	go producer(chanStorage)
	go logistics(chanStorage, chanShop)
	go consumer(chanShop)

	for {
		time.Sleep(time.Second)
	}

}

/* 生产者 */
func producer(chanStorage chan string) {
	for i := 0 ; i < 10; i ++  {
		product := "[" + strconv.Itoa(i) + "]"+ strconv.Itoa(time.Now().Nanosecond())
		chanStorage <- "产品" + product

		fmt.Println("生产了产品", product)
		time.Sleep(time.Second)
	}
	close(chanStorage)
	fmt.Println("仓库已关张，生产结束！")
}

/* 物流公司 */
func logistics(chanStorage, chanShop chan string) {
	for p := range chanStorage {
		chanShop <- p
		fmt.Println("转运了", p)
	}
	close(chanShop)
	fmt.Println("物流转运全部完成，商店已关张")
}

func consumer(chanShop chan string) {
	for product := range chanShop {
		fmt.Println("消费了产品", product)
		fmt.Println()
	}
	fmt.Println("商品已全部被消费")
}