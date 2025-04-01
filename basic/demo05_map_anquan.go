package main

import (
	"fmt"
	"sync"
	"time"
)

var m map[string]interface{}
var lock *sync.RWMutex

func main() {
	// map 初始化
	m = make(map[string]interface{})
	lock = new(sync.RWMutex)
	for i := 0; i < 2000; i++ {
		go writeM("name", "zhangsan")
		go readM("name")
		go writeM("age", 18)
		go readM("age")
		go writeM("name", "zhangsan1")
		go readM("name")
		go writeM("age", 19)
		go readM("age")
	}
	time.Sleep(10 * time.Second)
}


func readM(key string) {
	lock.RLock()
	fmt.Println(m[key])
	lock.RUnlock()
}
func writeM(key string, value interface{}) {
	lock.Lock()
	m[key] = value
	lock.Unlock()
}
