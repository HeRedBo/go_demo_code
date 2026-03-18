package limiter

import (
	"fmt"
	"testing"
	"time"
)

func TestWindowLimiter(t *testing.T) {

	limiter := NewFixedWindowLimiter(5, time.Second) // 1秒最多5次请求
	for i := 1; i <= 8; i++ {
		time.Sleep(100 * time.Millisecond)
		if limiter.AllowRequest() {
			fmt.Printf("请求%d：允许\n", i)
		} else {
			fmt.Printf("请求%d：限流\n", i)
		}
	}
}

func TestSlidingWindowLimiter(t *testing.T) {
	limiter := NewSlidingWindowLimiter(5, time.Second) // 1秒最多5次请求
	key := "user_456"
	for i := 1; i <= 8; i++ {
		time.Sleep(100 * time.Millisecond)
		if limiter.AllowRequest(key) {
			fmt.Printf("请求%d：允许（%s）\n", i, key)
		} else {
			fmt.Printf("请求%d：限流（%s）\n", i, key)
		}
	}
}
