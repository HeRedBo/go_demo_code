package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// --- 模拟数据源 ---
// 模拟三个不同的搜索引擎
var sources = []string{"Google", "Bing", "Baidu"}

// --- 1. 生成器 ---
// 模拟用户输入关键词，产生任务
func queryGenerator(keyword string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		// 这里为了演示简单只发一个，实际可以是读取文件的一行行关键词
		out <- keyword
	}()
	return out
}

// --- 2. 扇出：具体的干活的函数 ---
// 模拟去某个渠道搜索，耗时随机
func searchFromSource(source, keyword string) string {
	// 模拟网络延迟 (100ms - 500ms)
	time.Sleep(time.Duration(rand.Intn(400)+100) * time.Millisecond)
	return fmt.Sprintf("[%s] 搜到结果: %s 相关页面", source, keyword)
}

// --- 3. 核心逻辑：扇出处理 ---
// 接收关键词通道，为每个关键词启动多个协程去不同源搜索
func fanOutSearch(in <-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup

	// 启动一个协程专门负责“分发”
	go func() {
		defer close(out) // 等分发完了，关闭输出通道（注意：这里只是分发完，不是所有结果收完）

		for keyword := range in {
			wg.Add(len(sources)) // 增加计数，有几个源就加几个

			// 【扇出】：针对同一个关键词，同时向三个源发起搜索
			for _, source := range sources {
				go func(s string) {
					defer wg.Done()
					result := searchFromSource(s, keyword)
					out <- result // 把结果扔到输出通道
				}(source)
			}
		}
		// 等待这一批次所有的搜索任务都发完了（注意：这里有个陷阱，见下方解析）
		// 但为了演示扇入，我们通常不在这关闭 out，而是交给 merge 去处理
	}()

	// 注意：上面的逻辑有个小问题，如果 in 关闭了，wg.Wait() 需要有个地方放
	// 为了演示清晰，我们把这个逻辑简化：
	// 我们直接返回 out，让 merge 函数来负责收集所有结果并关闭最终通道
	return out
}

// --- 4. 扇入：合并结果 ---
// 将多个输入通道合并为一个输出通道
func searchMerge(cs ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	// 为每个输入通道启动一个输出协程
	output := func(c <-chan string) {
		defer wg.Done()
		for n := range c {
			out <- n
		}
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// 启动一个协程，等所有输入协程都结束后，关闭输出通道
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// --- 5. 业务编排：把上面串起来 ---
// 这里我们需要稍微修改一下扇出的逻辑，让它返回多个通道，然后再合并
// 或者更常见的模式：输入一个通道，输出一个通道，中间内部扇出扇入
func runSearchPipeline(keyword string) {
	// 1. 产生任务
	queries := queryGenerator(keyword)

	// 2. 准备收集结果
	// 我们需要一个通道来收集所有渠道的结果
	allResults := make(chan string)

	// 用来等待所有搜索协程结束
	var wg sync.WaitGroup

	go func() {
		defer close(allResults) // 任务分发完后，等所有结果都发了，再关闭总结果通道
		// 监听任务通道
		for q := range queries {
			// 【扇出】：针对这个关键词，启动 N 个协程去搜
			for _, source := range sources {
				wg.Add(1)
				go func(s, k string) {
					defer wg.Done()
					res := searchFromSource(s, k)

					// 【危险点】
					// 如果这里出错了（比如网络超时、API 报错），代码直接 return 了
					// if err != nil {
					// 	return
					// }
					// 这看起来好像没事？但如果你的逻辑是“必须收集满 3 个结果才关闭”，或者你在 merge 模式中，某个输入通道因为没有发送数据而没有被关闭，接收方（range allResults）就会永远等待，导致死锁。
					// 可以考虑 结果分钟 和 错误通道

					allResults <- res // 【扇入】：大家都往这一个通道里发结果
				}(source, q)
			}
		}

		// 等所有搜索都做完，再关闭 allResults
		wg.Wait()
	}()

	// 3. 消费结果
	fmt.Printf("🔍 正在搜索: %s ...\n", keyword)
	start := time.Now()

	count := 0
	for result := range allResults {
		count++
		fmt.Printf("✅ 收到结果 %d/%d: %s\n", count, len(sources), result)
	}

	fmt.Printf("🏁 搜索完成，总耗时: %v\n\n", time.Since(start))
}

func main03() {
	// 随机数种子
	// 运行搜索
	runSearchPipeline("Golang 并发编程")
}
