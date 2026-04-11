package qianwen

import (
	"sync"
	"time"
)

// LeakyBucket 结构体定义了漏桶的属性
type LeakyBucket struct {
	capacity int           // 桶的最大容量（能处理的最大突发请求数）
	water    int           // 当前桶里的水量（当前积压的请求数）
	rate     time.Duration // 漏水速率（每隔多久处理一个请求）
	lastLeak time.Time     // 上次漏水的时间
	mu       sync.Mutex    // 互斥锁，保证并发安全
}

// NewLeakyBucket 创建一个新的漏桶
// capacity: 桶容量
// rate: 漏水速率（例如 100ms 漏一滴，代表每秒处理10个请求）
func NewLeakyBucket(capacity int, rate time.Duration) *LeakyBucket {
	if rate <= 0 {
		panic("漏桶速率必须大于 0")
	}
	return &LeakyBucket{
		capacity: capacity,
		water:    0, // 初始时桶是空的
		rate:     rate,
		lastLeak: time.Now(),
	}
}

// Allow 尝试向桶中加水（处理请求）
// 如果桶未满，则允许请求进入（返回 true）
// 如果桶已满，则拒绝请求（返回 false）
func (lb *LeakyBucket) Allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	// 1. 先执行“漏水”操作，计算这段时间内漏掉了多少水
	now := time.Now()
	// 计算经过的时间能漏掉多少请求
	elapsed := now.Sub(lb.lastLeak)
	// 漏掉的数量 = 经过时间 / 漏水速率
	leakCount := int(elapsed / lb.rate)

	if leakCount > 0 {
		lb.water -= leakCount
		if lb.water < 0 {
			lb.water = 0
		}
		lb.lastLeak = now
	}

	// 2. 尝试加水（接受新请求）
	// 如果当前水量小于容量，说明桶没满，可以加水
	if lb.water < lb.capacity {
		lb.water++
		return true
	}

	// 桶满了，拒绝请求（限流触发）
	return false
}
