package main
import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var counter2 int = 0

func add5(a, b int, lock *sync.RWMutex) {
	c := a + b
	lock.Lock()
	counter2++
	fmt.Printf("%d: %d + %d = %d\n", counter2, a, b, c)
	lock.Unlock()
}

func main() {
	start := time.Now()
	lock := &sync.RWMutex{}
	for i := 0; i < 10; i++ {
		go add5(1, i, lock)
	}

	for {
		lock.Lock()
		c := counter2
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}

	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
