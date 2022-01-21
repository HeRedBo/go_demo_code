package main

import (
	"fmt"
	"math"
	"time"
)

/*管道数据读写*/
func main() {
	//mSlice := make([]int, 1024)
	//mMap := make(map[string]interface{})
	//channel
	ch := make(chan int)

	go func() {
		//阻塞从管道中读数据
		x := <-ch
		fmt.Println("读出数据", x)
	}()

	//往管道里写数据
	ch <- 123

	time.Sleep(1 * time.Second)
	fmt.Println("Game Over")
}

/*关闭管道的几个注意事项*/
func main072() {

	/*不能关闭一个没有初始化的管道*/
	//var ch chan int
	////panic: close of nil channel
	//close(ch)

	/*管道不能重复关闭*/
	//ch := make(chan int)
	//close(ch)
	////panic: close of closed channel
	//close(ch)

	ch := make(chan int, 3)
	ch <- 123
	close(ch)
	////  send on closed channel
	//ch <- 1234

	go func() {
		//从一个已经关闭的管道中读取数据
		x := <-ch
		fmt.Println("读到", x)

		//已经读完了之后,继续读,读到int类型的默认值
		x = <-ch
		fmt.Println("读到", x)

		//已经读完了之后,继续读,读到int类型的默认值
		x = <-ch
		fmt.Println("读到", x)

		//panic: send on closed channel
		ch <- 456
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Game Over")
}


/*读取管道数据时校验其有效性*/
func main073() {
	//创建一个具有3缓冲能力的管道
	ch := make(chan int, 3)
	ch <- 123
	ch <- 456
	close(ch)

	go func() {
		//x, ok := <-ch
		//fmt.Println(x, ok)

		//x, ok = <-ch
		//fmt.Println(x, ok)
		//
		//x, ok = <-ch
		//fmt.Println(x, ok)

		for i := 0; i < 5; i++ {
			//x, ok := <-ch;
			//fmt.Println(x,ok)

			if x, ok := <-ch; ok == true {
				fmt.Println("读到数据", x)
			} else {
				fmt.Println("数据已读尽")
			}

		}

	}()

	time.Sleep(time.Second)
}


/*遍历管道中的全部数据*/
func main074() {
	ch := make(chan int, 10)
	for i := 0; i < 5; i++ {
		ch <- i * i
	}
	close(ch)

	for x := range ch {
		fmt.Println(x)
	}
}

/*关闭管道时通知读取协程结束遍历*/
func main075() {
	ch := make(chan float64, 10)

	go func() {
		for i := 1; i < 6; i++ {
			ch <- math.Pow(float64(i), 2)
			time.Sleep(1 * time.Second)
		}

		//宣布管道关闭——所有读取协程会结束对该管道的遍历
		close(ch)
		fmt.Println("写入协程完毕！")
	}()

	go func() {
		for x := range ch{
			fmt.Println("读到",x)
		}
		fmt.Println("读取协程完毕！")
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("Game Over!")
}
