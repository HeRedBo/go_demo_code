package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func div(num int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		wg.Done()
	}()
	fmt.Printf("10/%d=%d \n", num, 100/num)
}
func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go div(i)
	}
	wg.Wait()
}

// 以上代码存在异常
