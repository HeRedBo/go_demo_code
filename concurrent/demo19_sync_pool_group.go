package main

import (
	"fmt"
	"sync"
)

func test_put(pool *sync.Pool , deferFunc func()) {
	defer func() {
		deferFunc()
	}()

	value := "Hello , Golang boy"
	pool.Put(value)
}

func main() {
	var wg  sync.WaitGroup
	wg.Add(1)
	var pool = &sync.Pool{
		New: func() interface{} {
			return "Hello ,World!"
		},
	}
	go test_put(pool, wg.Done)
	wg.Wait()
	fmt.Println(pool.Get())
}


