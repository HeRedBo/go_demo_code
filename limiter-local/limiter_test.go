package limiterlocal

import (
	"fmt"
	"sync"
	"sync/atomic"
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

func TestLeakyBucket_FullFeature(t *testing.T) {

	bucket := NewLeakyBucket(10, 2) // 桶容量10，每秒处理2个请求

	// 模拟20次请求（高频入桶）
	for i := 1; i <= 20; i++ {
		time.Sleep(50 * time.Millisecond) // 间隔0.05秒
		if bucket.AddRequest() {
			fmt.Printf("请求%d：入桶（当前水量：%.2f）\n", i, bucket.water)
		} else {
			fmt.Printf("请求%d：被限流（桶已满）\n", i)
		}
	}
}

func TestLeakyBucket(t *testing.T) {
	fmt.Println("=== 漏桶算法测试开始 ===")
	// 创建漏桶：容量 5 个，每秒漏水 2 个（每秒处理2个请求）
	bucket := NewLeakyBucket(5, 2)

	// 模拟 10 个并发请求瞬间进来
	for i := 1; i <= 10; i++ {
		go func(idx int) {
			ok := bucket.AddRequest()
			if ok {
				fmt.Printf("请求 %d：通过 ✅\n", idx)
			} else {
				fmt.Printf("请求 %d：被限流 ❌\n", idx)
			}
		}(i)
	}

	// 等待协程执行完
	time.Sleep(1 * time.Second)

	fmt.Println("\n=== 等待 2 秒后，再次发送 5 个请求 ===")
	time.Sleep(2 * time.Second)

	// 再次发送 5 个请求
	for i := 11; i <= 15; i++ {
		go func(idx int) {
			ok := bucket.AddRequest()
			if ok {
				fmt.Printf("请求 %d：通过 ✅\n", idx)
			} else {
				fmt.Printf("请求 %d：被限流 ❌\n", idx)
			}
		}(i)
	}

	// 等待结束
	time.Sleep(1 * time.Second)
	fmt.Println("=== 测试结束 ===")
}

// region 滑动串口单元测试代码

// go test -v -run=SlidingWindow
// 测试滑动窗口：正常速率不被限流
func TestSlidingWindow_NormalRate_Allow(t *testing.T) {
	// 1秒最多允许 5 个请求
	limiter := NewSlidingWindow(5, time.Second)

	// 连续请求 5 次，都应该允许
	for i := 0; i < 5; i++ {
		if !limiter.Allow() {
			t.Fatalf("第 %d 个请求应该被允许，但被限流了", i+1)
		}
	}
}

// 测试滑动窗口：超过限制会被限流

func TestSlidingWindow_ExceedRate_Reject(t *testing.T) {
	limiter := NewSlidingWindow(5, time.Second)

	// 前 5 个通过
	//for i := 0; i < 5; i++ {
	//	limiter-local.Allow()
	//}
	//// 第 6 个应该被拒绝
	//if limiter-local.Allow() {
	//	t.Fatal("第 6 个请求应该被限流，但通过了")
	//}

	for i := 0; i < 10; i++ {
		if limiter.Allow() {
			fmt.Printf("请求%d：允许通过 ✅\n", i)
		} else {
			fmt.Printf("请求%d：被限流 ❌\n", i)
		}
	}

}

// 测试滑动窗口：窗口滑动后，旧请求会被淘汰
func TestSlidingWindow_WindowSlides_AllowAgain(t *testing.T) {
	limiter := NewSlidingWindow(5, time.Second)

	// 先打满 5 个
	for i := 0; i < 5; i++ {
		limiter.Allow()
	}

	// 等待超过窗口时间（1.1秒），让窗口滑动
	time.Sleep(1100 * time.Millisecond)

	// 再次请求应该通过
	if !limiter.Allow() {
		t.Fatal("窗口滑动后应该允许新请求，但被拒绝了")
	}
}

// 测试滑动窗口：并发安全
func TestSlidingWindow_Concurrent(t *testing.T) {
	limiter := NewSlidingWindow(10, time.Second)
	var success int32
	var wg sync.WaitGroup

	// 并发 20 个请求
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if limiter.Allow() {
				atomic.AddInt32(&success, 1)
			}
		}()
	}
	wg.Wait()

	// 最终成功数必须 = 10
	if success != 10 {
		t.Fatalf("并发测试失败：期望 10 个通过，实际 %d 个", success)
	}
}

// endregion

func TestFixedWindow(t *testing.T) {
	fmt.Println("=== 固定窗口限流器测试开始 ===")
	limiter := NewFixedWindow(5, time.Second) // 1秒最多5次请求
	for i := 1; i <= 8; i++ {
		time.Sleep(100 * time.Millisecond)
		if limiter.Allow() {
			fmt.Printf("请求%d：允许通过 ✅\n", i)
		} else {
			fmt.Printf("请求%d：被限流 ❌\n", i)
		}
	}
	fmt.Println("=== 测试结束 ===")
}

func TestNewFixedWindowLimiter(t *testing.T) {
	fmt.Println("=== 固定窗口限流器测试开始 ===")
	limiter := NewFixedWindowLimiter(5, time.Second) // 1秒最多5次请求
	for i := 1; i <= 8; i++ {
		time.Sleep(100 * time.Millisecond)
		if limiter.AllowRequest() {
			fmt.Printf("请求%d：允许通过 ✅\n", i)
		} else {
			fmt.Printf("请求%d：被限流 ❌\n", i)
		}
	}
	fmt.Println("=== 测试结束 ===")
}
