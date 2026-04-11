package limiter

import (
	"sync"
	"time"
)

// TokenBucket 令牌桶限流器
type TokenBucket struct {
	capacity     int           // 桶容量
	refillRate   time.Duration // 桶填充速率（每产生一个令牌所需的时间）
	tokens       int           // 桶当前令牌数
	lastRefillAt time.Time     // 最后填充令牌时间
	mu           sync.Mutex    // 互斥锁，保证并发安全
}

// NewTokenBucket 创建一个新的令牌桶
// capacity: 桶的容量，最大令牌数
// rate: 令牌生成速率，例如 time.Millisecond * 100 表示每100毫秒生成一个令牌
func NewTokenBucket(capacity int, refillRate time.Duration) *TokenBucket {
	return &TokenBucket{
		capacity:     capacity,
		refillRate:   refillRate,
		tokens:       capacity,   // 初始化桶满（业务最佳实践）
		lastRefillAt: time.Now(), // 初始化时间戳
	}
}

// 内部工具方法：刷新令牌（消除代码冗余）
func (t *TokenBucket) refill() {
	now := time.Now()
	duration := now.Sub(t.lastRefillAt)

	// 计算应补充的令牌数量
	addTokens := int(duration / t.refillRate)
	if addTokens <= 0 {
		return
	}
	// 补充令牌，不超过容量
	t.tokens += addTokens
	if t.tokens > t.capacity {
		t.tokens = t.capacity
	}
	t.lastRefillAt = now
}

// Allow 检查是否允许请求通过（消耗一个令牌）
func (t *TokenBucket) Allow() bool {
	return t.AllowN(1)
}

// AllowN 检查是否允许N个请求通过（消耗N个令牌）
// AllowN 批量获取 N 个令牌
func (t *TokenBucket) AllowN(n int) bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.refill() // 统一刷新令牌
	if t.tokens >= n {
		t.tokens -= n
		return true
	}
	return false
}

// SetRate 设置新的令牌生成速率
func (t *TokenBucket) SetRate(rate time.Duration) {
	t.mu.Lock()
	defer t.mu.Unlock()

	// 先更新令牌数
	now := time.Now()
	tokensToAdd := int(now.Sub(t.lastRefillAt) / t.refillRate)
	if tokensToAdd > 0 {
		t.tokens += tokensToAdd
		if t.tokens > t.capacity {
			t.tokens = t.capacity
		}
		t.lastRefillAt = now
	}

	// 设置新速率
	t.refillRate = rate
}

// SetCapacity 设置新的桶容量
func (t *TokenBucket) SetCapacity(capacity int) {
	t.mu.Lock()
	defer t.mu.Unlock()
	// 先更新令牌数
	t.refill()
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
	t.refill()
	return t.tokens
}
