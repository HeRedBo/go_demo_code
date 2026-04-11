package limiter

import (
	"fmt"
	"testing"
	"time"
)

func TestTokenBucket_FullFeature(t *testing.T) {
	// 容量 5，每 200ms 生成1个令牌（5 QPS）
	bucket := NewTokenBucket(5, 200*time.Millisecond)

	fmt.Println("=== 初始化完成 ===")
	fmt.Printf("桶容量：%d\n", bucket.capacity)
	fmt.Printf("初始令牌数：%d\n\n", bucket.TokensAvailable())

	// 1. 瞬间请求 10 次
	allow, deny := 0, 0
	for i := 0; i < 10; i++ {
		if bucket.Allow() {
			allow++
		} else {
			deny++
		}
	}
	fmt.Printf("瞬间请求10次：通过=%d，拒绝=%d\n", allow, deny)

	// 2. 批量获取
	fmt.Printf("\n当前剩余令牌：%d\n", bucket.TokensAvailable())
	fmt.Printf("尝试批量获取3个：%t\n", bucket.AllowN(3))
	fmt.Printf("尝试批量获取10个：%t\n", bucket.AllowN(10))

	// 3. 等待令牌恢复
	fmt.Println("\n等待 1 秒...")
	time.Sleep(1 * time.Second)
	fmt.Printf("等待后可用令牌：%d\n", bucket.TokensAvailable())

	// 4. 动态修改速率（变快）
	fmt.Println("\n→ 动态加速：每 100ms 生成1个令牌")
	bucket.SetRate(100 * time.Millisecond)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("加速后可用令牌：%d\n", bucket.TokensAvailable())

	// 5. 动态修改容量
	fmt.Println("\n→ 动态扩容到 10")
	bucket.SetCapacity(10)
	fmt.Printf("扩容后容量：%d\n", bucket.capacity)
	fmt.Printf("扩容后可用令牌：%d\n", bucket.TokensAvailable())
}
