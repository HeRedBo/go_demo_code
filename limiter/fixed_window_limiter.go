package limiter

import (
	"fmt"
	"sync/atomic"
	"time"
)

// FixedWindowLimiter 固定窗口限流器
type FixedWindowLimiter struct {
	threshold int64         // 窗口内最大请求数
	window    time.Duration // 窗口大小
	count     int64         // 当前窗口请求数
	lastTime  time.Time     // 窗口起始时间
}

// NewFixedWindowLimiter 创建固定窗口限流器
func NewFixedWindowLimiter(threshold int64, window time.Duration) *FixedWindowLimiter {
	return &FixedWindowLimiter{
		threshold: threshold,
		window:    window,
		lastTime:  time.Now(),
	}
}

// AllowRequest 检查是否允许请求
func (fw *FixedWindowLimiter) AllowRequest() bool {
	now := time.Now()
	// 1. 判断是否进入新窗口
	if now.Sub(fw.lastTime) >= fw.window {
		// 重置窗口
		atomic.StoreInt64(&fw.count, 0)
		fw.lastTime = now
	}
	// 2. 原子递增请求数
	current := atomic.AddInt64(&fw.count, 1)
	return current <= fw.threshold
}

// 测试固定窗口
func testFixedWindow() {
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

func main() {
	testFixedWindow()
}
