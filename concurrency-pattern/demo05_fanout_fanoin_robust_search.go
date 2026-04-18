package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// --- 1. 定义结果结构体 ---
// 这是一个最佳实践：永远不要只返回数据，要把错误和数据包在一起
type SearchResult struct {
	Source string
	Data   string
	Err    error
}

// --- 模拟搜索引擎 ---
func mockSearch(source string, timeout time.Duration) (string, error) {
	// 模拟随机耗时
	delay := time.Duration(rand.Intn(1000)) * time.Millisecond

	// 模拟超时
	if delay > timeout {
		return "", fmt.Errorf("%s 响应超时 (%v > %v)", source, delay, timeout)
	}

	// 模拟随机错误 (10% 概率挂掉)
	if rand.Float32() < 0.1 {
		return "", fmt.Errorf("%s 服务器内部错误", source)
	}

	time.Sleep(delay)
	return fmt.Sprintf("来自 %s 的搜索结果", source), nil
}

// --- 2. 核心逻辑：带错误处理的扇出/扇入 ---
func robustSearch(keyword string) {
	// 结果通道
	resultChan := make(chan SearchResult)
	var wg sync.WaitGroup

	// 模拟超时时间
	searchTimeout := 800 * time.Millisecond

	fmt.Printf("🔍 开始搜索: %s (超时限制: %v)\n", keyword, searchTimeout)
	start := time.Now()

	// 【扇出】
	for _, source := range []string{"Google", "Bing", "Baidu"} {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()

			// 执行搜索
			data, err := mockSearch(s, searchTimeout)

			// 【关键】无论成功还是失败，都发送结果
			// 这样接收端才知道发生了什么，不会傻等
			resultChan <- SearchResult{
				Source: s,
				Data:   data,
				Err:    err,
			}
		}(source)
	}

	// 【监控协程】：等所有任务做完，关闭通道
	// 这一步非常重要，否则 range 会死锁
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// 【扇入 / 消费结果】
	successCount := 0
	errorCount := 0

	for result := range resultChan {
		if result.Err != nil {
			// 处理错误
			fmt.Printf("❌ [%s] 失败: %v\n", result.Source, result.Err)
			errorCount++
		} else {
			// 处理成功
			fmt.Printf("✅ [%s] 成功: %s\n", result.Source, result.Data)
			successCount++
		}
	}

	fmt.Printf("🏁 搜索结束。耗时: %v, 成功: %d, 失败: %d\n", time.Since(start), successCount, errorCount)
}

func main05() {
	// 运行搜索
	robustSearch("Golang 错误处理")
}
