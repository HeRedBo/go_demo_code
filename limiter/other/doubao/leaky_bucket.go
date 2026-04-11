package doubao

import (
	"time"
)

// LeakyBucket 漏桶结构体
type LeakyBucket struct {
	capacity float64       // 桶最大容量
	rate     float64       // 漏水速率（个/秒）
	water    float64       // 当前水量
	lastTime time.Time     // 上次漏水时间
	mu       chan struct{} // 并发锁
}

// NewLeakyBucket 创建漏桶实例
func NewLeakyBucket(capacity, rate float64) *LeakyBucket {

	return &LeakyBucket{
		capacity: capacity,
		rate:     rate,
		water:    0,
		lastTime: time.Now(),
		mu:       make(chan struct{}, 1),
	}
}

// AddRequest 尝试添加请求到漏桶
// 返回 true 表示请求被接受，false 表示被拒绝
func (lb *LeakyBucket) AddRequest() bool {
	lb.mu <- struct{}{}
	defer func() { <-lb.mu }()

	// 1. 计算漏水总量
	now := time.Now()
	elapsed := now.Sub(lb.lastTime).Seconds()
	leakWater := elapsed * lb.rate

	// 2. 更新当前水量
	lb.water = max(0.0, lb.water-leakWater)
	lb.lastTime = now

	// 3. 判断是否能入桶
	if lb.water < lb.capacity {
		lb.water += 1
		return true
	}
	return false
}

// GetWaterLevel 获取当前桶内水位
func (lb *LeakyBucket) GetWaterLevel() float64 {
	lb.mu <- struct{}{}
	defer func() { <-lb.mu }()

	// 先更新水位
	now := time.Now()
	elapsed := now.Sub(lb.lastTime).Seconds()
	leakWater := elapsed * lb.rate
	lb.water = max(0.0, lb.water-leakWater)
	lb.lastTime = now

	return lb.water
}

// Reset 重置漏桶
func (lb *LeakyBucket) Reset() {
	lb.mu <- struct{}{}
	defer func() { <-lb.mu }()

	lb.water = 0
	lb.lastTime = time.Now()
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
