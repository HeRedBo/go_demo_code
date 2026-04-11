package limiter

import (
	"sync"
	"time"
)

// LeakyBucket 漏桶结构体
type LeakyBucket struct {
	capacity float64    // 桶最大容量
	rate     float64    // 漏水速率（个/秒）
	water    float64    // 当前水量
	lastTime time.Time  // 上次漏水时间
	mu       sync.Mutex // 并发锁
}

// NewLeakyBucket 创建漏桶实例
func NewLeakyBucket(capacity, rate float64) *LeakyBucket {
	return &LeakyBucket{
		capacity: capacity,
		rate:     rate,
		water:    0,
		lastTime: time.Now(),
	}
}

// AddRequest 尝试添加请求到漏桶
func (lb *LeakyBucket) AddRequest() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	// 1. 计算漏水总量
	now := time.Now()
	elapsed := now.Sub(lb.lastTime).Seconds() // 计算当前时间到上次漏水时间的时间间隔
	leakWater := elapsed * lb.rate            // 计算漏水总量

	// 2. 更新当前水量 使用内置函数 如果考虑兼容 使用自定一函数 maxFloat
	lb.water = maxFloat(0.0, lb.water-leakWater) // 当前水量  旧水量 - 漏水量,但是它不能小于0.0
	lb.lastTime = now

	// 3. 判断是否能入桶
	if lb.water < lb.capacity {
		lb.water += 1
		return true
	}
	return false
}

func maxFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

//func max(a, b float64) float64 {
//	if a > b {
//		return a
//	}
//	return b
//}
