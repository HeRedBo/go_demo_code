package qianwen

import (
	"fmt"
	"testing"
	"time"
)

func TestNewLeaky(t *testing.T) {
	bucket := NewLeakyBucket(5, 200*time.Millisecond)
	fmt.Println("开始模拟请求...")
	// 模拟 10 个并发请求
	for i := 0; i < 10; i++ {
		go func(id int) {
			if bucket.Allow() {
				fmt.Printf("请求 %d: ✅ 允许通过\n", id)
			} else {
				fmt.Printf("请求 %d: ❌ 拒绝（限流）\n", id)
			}
		}(i)
	}
	// 等待一段时间让 goroutine 执行完
	time.Sleep(2 * time.Second)
}
