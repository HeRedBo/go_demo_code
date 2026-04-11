package doubao

import (
	"sync"
	"time"
)

// TokenBucket 令牌桶限流器结构体（优化版）
type TokenBucket struct {
	mutex        sync.Mutex    // 互斥锁：保证高并发下线程安全
	capacity     int           // 桶的最大容量（最大并发/最大突发流量）
	refillRate   time.Duration // 填充间隔：每隔多久生成 1 个令牌
	tokens       int           // 当前剩余令牌数
	lastRefillAt time.Time     // 上一次刷新令牌的时间
}

// NewTokenBucket 创建令牌桶限流器
// capacity：令牌桶最大容量
// refillRate：生成 1 个令牌的时间间隔（例如 100*time.Millisecond = 10ms 生成1个，即 100 QPS）
func NewTokenBucket(capacity int, refillRate time.Duration) *TokenBucket {
	now := time.Now()
	return &TokenBucket{
		capacity:     capacity,
		refillRate:   refillRate,
		tokens:       capacity, // 初始化时桶是满的，更符合实际使用
		lastRefillAt: now,
	}
}

// Allow 判断是否允许通过（获取 1 个令牌）
// 返回 true：允许通过；false：限流拒绝
func (t *TokenBucket) Allow() bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// 1. 计算当前需要补充的令牌数
	now := time.Now()
	duration := now.Sub(t.lastRefillAt)
	addTokens := int(duration / t.refillRate)

	// 2. 补充令牌，不超过桶容量
	if addTokens > 0 {
		t.tokens += addTokens
		if t.tokens > t.capacity {
			t.tokens = t.capacity
		}
		t.lastRefillAt = now
	}

	// 3. 没有令牌，直接拒绝
	if t.tokens < 1 {
		return false
	}

	// 4. 消耗 1 个令牌
	t.tokens--
	return true
}
