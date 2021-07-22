package main

import (
	"fmt"
	"sync"
)

func main() {
	var pool = &sync.Pool{
		New: func() interface{} {
			return "hello Word!"
		},
	}
	value := "Hello Red Bo!"
	pool.Put(value)
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
}