package limiter

import (
	"sync"
	"time"
)

// TokenBucket 令牌桶限流器
type TokenBucket struct {
	capacity  int           // 桶容量
	rate      time.Duration // 桶填充速率（每产生一个令牌所需的时间）
	tokens    int           // 桶当前令牌数
	timestamp time.Time     // 最后填充令牌时间
	mu        sync.Mutex    // 互斥锁，保证并发安全
}

// NewTokenBucket 创建一个新的令牌桶
// capacity: 桶的容量，最大令牌数
// rate: 令牌生成速率，例如 time.Millisecond * 100 表示每100毫秒生成一个令牌
func NewTokenBucket(capacity int, rate time.Duration) *TokenBucket {
	return &TokenBucket{
		capacity:  capacity,
		rate:      rate,
		tokens:    0,
		timestamp: time.Now(), // 初始化时间戳
	}
}

// Allow 检查是否允许请求通过（消耗一个令牌）
func (t *TokenBucket) Allow() bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	// 计算从上次填充到现在的时间，根据速率增加令牌
	now := time.Now()
	tokensToAdd := int(now.Sub(t.timestamp) / t.rate)
	if tokensToAdd > 0 {
		t.tokens += tokensToAdd
		if t.tokens > t.capacity {
			t.tokens = t.capacity
		}
		t.timestamp = now
	}

	// 如果桶中有令牌，消耗一个并返回true
	if t.tokens > 0 {
		t.tokens--
		return true
	}

	// 桶是空的，请求被拒绝
	return false
}

// AllowN 检查是否允许N个请求通过（消耗N个令牌）
func (t *TokenBucket) AllowN(n int) bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	// 计算从上次填充到现在的时间，根据速率增加令牌
	now := time.Now()
	tokensToAdd := int(now.Sub(t.timestamp) / t.rate)
	if tokensToAdd > 0 {
		t.tokens += tokensToAdd
		if t.tokens > t.capacity {
			t.tokens = t.capacity
		}
		t.timestamp = now
	}

	// 如果桶中的令牌足够，消耗N个并返回true
	if t.tokens >= n {
		t.tokens -= n
		return true
	}

	// 桶中的令牌不足，请求被拒绝
	return false
}

// SetRate 设置新的令牌生成速率
func (t *TokenBucket) SetRate(rate time.Duration) {
	t.mu.Lock()
	defer t.mu.Unlock()

	// 先更新令牌数
	now := time.Now()
	tokensToAdd := int(now.Sub(t.timestamp) / t.rate)
	if tokensToAdd > 0 {
		t.tokens += tokensToAdd
		if t.tokens > t.capacity {
			t.tokens = t.capacity
		}
		t.timestamp = now
	}

	// 设置新速率
	t.rate = rate
}

// SetCapacity 设置新的桶容量
func (t *TokenBucket) SetCapacity(capacity int) {
	t.mu.Lock()
	defer t.mu.Unlock()

	// 先更新令牌数
	now := time.Now()
	tokensToAdd := int(now.Sub(t.timestamp) / t.rate)
	if tokensToAdd > 0 {
		t.tokens += tokensToAdd
		t.timestamp = now
	}

	// 设置新容量
	t.capacity = capacity
	if t.tokens > t.capacity {
		t.tokens = t.capacity
	}
}

// TokensAvailable 获取当前可用的令牌数
func (t *TokenBucket) TokensAvailable() int {
	t.mu.Lock()
	defer t.mu.Unlock()

	// 先更新令牌数
	now := time.Now()
	tokensToAdd := int(now.Sub(t.timestamp) / t.rate)
	if tokensToAdd > 0 {
		t.tokens += tokensToAdd
		if t.tokens > t.capacity {
			t.tokens = t.capacity
		}
		t.timestamp = now
	}

	return t.tokens
}
