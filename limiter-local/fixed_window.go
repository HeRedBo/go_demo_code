package limiterlocal

import (
	"sync"
	"time"
)

// FixedWindow 固定窗口
type FixedWindow struct {
	mu        sync.Mutex
	rate      int           // 每秒允许数
	interval  time.Duration // 窗口大小
	count     int           // 当前窗口请求数
	windowEnd time.Time     // 当前窗口结束时间
}

func NewFixedWindow(rate int, interval time.Duration) *FixedWindow {
	return &FixedWindow{
		rate:      rate,
		interval:  interval,
		windowEnd: time.Now().Add(interval),
	}
}

func (fw *FixedWindow) Allow() bool {
	fw.mu.Lock()
	defer fw.mu.Unlock()

	now := time.Now()
	// 1. 如果已经超过窗口结束时间 → 重置窗口
	if now.After(fw.windowEnd) {
		fw.count = 0
		fw.windowEnd = now.Add(fw.interval)
	}
	// 2. 判断是否超限
	if fw.count < fw.rate {
		fw.count++
		return true
	}

	return false
}
