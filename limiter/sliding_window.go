package limiter

import (
	"sync"
	"time"
)

// SlidingWindow 滑动窗口
type SlidingWindow struct {
	mu         sync.Mutex
	rate       int           // 最大允许数
	window     time.Duration // 窗口大小
	timestamps []time.Time   // 保存请求时间戳
}

func NewSlidingWindow(rate int, window time.Duration) *SlidingWindow {
	return &SlidingWindow{
		rate:       rate,
		window:     window,
		timestamps: make([]time.Time, 0, rate),
	}
}

func (sw *SlidingWindow) Allow() bool {
	sw.mu.Lock()
	defer sw.mu.Unlock()

	now := time.Now()
	// 窗口左边界：当前时间 - 窗口大小
	border := now.Add(-sw.window)

	// 1. 清理掉超出窗口的旧时间戳（只保留最近 window 内的）
	idx := 0
	for i, t := range sw.timestamps {
		if t.After(border) {
			idx = i
			break
		}
	}
	sw.timestamps = sw.timestamps[idx:]

	// 2. 判断是否超限
	if len(sw.timestamps) < sw.rate {
		sw.timestamps = append(sw.timestamps, now)
		return true
	}

	return false
}
