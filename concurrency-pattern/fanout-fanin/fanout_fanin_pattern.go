package main

import (
	"fmt"
	"math/rand/v2"
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
	delay := time.Duration(rand.IntN(400)+100) * time.Millisecond
	time.Sleep(delay)
	return fmt.Sprintf("[%s] 搜到结果: %s 相关页面", source, keyword)
}

// --- 3. 核心逻辑：扇出处理 ---
// 接收关键词通道，为每个关键词启动多个协程去不同源搜索
func fanOutSearch(in <-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup

	// 启动一个协程专门负责"分发"
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
func merge(cs ...<-chan string) <-chan string {
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

func main() {
	// Go 1.20+ 不需要手动设置 rand.Seed，math/rand/v2 已经自动初始化

	// 运行搜索
	runSearchPipelipackage main

import (
    "fmt"
    "math/rand"
    "math/rand/v2"
    "time"
)

// RandomSeedDemo 演示随机数种子的作用
type RandomSeedDemo struct{}

// WithoutSeed 不设置种子 - 每次运行产生相同的随机数序列
func (d *RandomSeedDemo) WithoutSeed(count int) []int {
    results := make([]int, count)
    for i := 0; i < count; i++ {
        results[i] = rand.Intn(100)
    }
    return results
}

// WithFixedSeed 使用固定种子 - 每次运行产生相同但可预测的随机数序列
func (d *RandomSeedDemo) WithFixedSeed(seed int64, count int) []int {
    r := rand.New(rand.NewSource(seed))
    results := make([]int, count)
    for i := 0; i < count; i++ {
        results[i] = r.Intn(100)
    }
    return results
}

// WithTimeSeed 使用时间种子（旧方式）- Go 1.20 之前的做法
func (d *RandomSeedDemo) WithTimeSeed(count int) []int {
    rand.Seed(time.Now().UnixNano())
    results := make([]int, count)
    for i := 0; i < count; i++ {
        results[i] = rand.Intn(100)
    }
    return results
}

// WithAutoSeed 自动种子（新方式）- Go 1.20+ 推荐
func (d *RandomSeedDemo) WithAutoSeed(count int) []int {
    results := make([]int, count)
    for i := 0; i < count; i++ {
        results[i] = rand.Intn(100)
    }
    return results
}

// WithRandV2 使用新的 rand/v2 包 - Go 1.22+ 推荐
func (d *RandomSeedDemo) WithRandV2(count int) []int {
    results := make([]int, count)
    for i := 0; i < count; i++ {
        results[i] = rand.IntN(100)
    }
    return results
}

// CompareSeeds 对比不同种子设置的效果
func (d *RandomSeedDemo) CompareSeeds() {
    fmt.Println("=" + string(make([]byte, 70)))
    fmt.Println("🎲 随机数种子对比演示")
    fmt.Println("=" + string(make([]byte, 70)))

    count := 5

    // 1. 不设置种子（Go 1.20 之前会有问题）
    fmt.Println("\n1️⃣  不设置种子（默认行为）:")
    fmt.Println("   第一次调用:", d.WithoutSeed(count))
    fmt.Println("   第二次调用:", d.WithoutSeed(count))
    fmt.Println("   💡 Go 1.20+ 已自动初始化，两次结果会不同")

    // 2. 使用固定种子
    fmt.Println("\n2️⃣  使用固定种子 (seed=42):")
    result1 := d.WithFixedSeed(42, count)
    result2 := d.WithFixedSeed(42, count)
    fmt.Println("   第一次调用:", result1)
    fmt.Println("   第二次调用:", result2)
    if fmt.Sprint(result1) == fmt.Sprint(result2) {
        fmt.Println("   ✅ 两次结果完全相同（可重现）")
    }

    // 3. 使用时间种子（旧方式）
    fmt.Println("\n3️⃣  使用时间种子（Go 1.20 之前的方式）:")
    fmt.Println("   第一次调用:", d.WithTimeSeed(count))
    time.Sleep(1 * time.Millisecond)
    fmt.Println("   第二次调用:", d.WithTimeSeed(count))
    fmt.Println("   ⚠️  注意：rand.Seed 在 Go 1.20+ 已废弃")

    // 4. 自动种子（Go 1.20+ 推荐）
    fmt.Println("\n4️⃣  自动种子（Go 1.20+ 推荐方式）:")
    fmt.Println("   第一次调用:", d.WithAutoSeed(count))
    fmt.Println("   第二次调用:", d.WithAutoSeed(count))
    fmt.Println("   ✅ 无需手动设置，每次运行都不同")

    // 5. 使用 rand/v2（Go 1.22+ 最新推荐）
    fmt.Println("\n5️⃣  使用 math/rand/v2（Go 1.22+ 最新推荐）:")
    fmt.Println("   第一次调用:", d.WithRandV2(count))
    fmt.Println("   第二次调用:", d.WithRandV2(count))
    fmt.Println("   ✅ 更好的 API，自动初始化")

    fmt.Println("\n" + "=" + string(make([]byte, 70)))
    fmt.Println("📝 总结:")
    fmt.Println("   • Go 1.20 之前: 必须手动调用 rand.Seed()")
    fmt.Println("   • Go 1.20-1.21: 自动初始化，但仍可使用 rand.Seed()")
    fmt.Println("   • Go 1.22+: 推荐使用 math/rand/v2 包")
    fmt.Println("   • 固定种子用于测试，时间/自动种子用于生产")
    fmt.Println("=" + string(make([]byte, 70)))
}

// SimulateNetworkDelay 模拟网络延迟（类似 fanout-fanin 场景）
func (d *RandomSeedDemo) SimulateNetworkDelay(sources []string) map[string]time.Duration {
    delays := make(map[string]time.Duration)
    
    fmt.Printf("\n🌐 模拟 %d 个搜索引擎的网络延迟:\n", len(sources))
    for _, source := range sources {
        delay := time.Duration(rand.IntN(400)+100) * time.Millisecond
        delays[source] = delay
        fmt.Printf("   %-10s: %v\n", source, delay)
    }
    
    return delays
}

func main() {
    demo := &RandomSeedDemo{}
    
    // 运行对比演示
    demo.CompareSeeds()
    
    // 模拟实际场景
    sources := []string{"Google", "Bing", "Baidu"}
    fmt.Println("\n🔍 实际应用场景演示:")
    fmt.Println("第一次运行:")
    demo.SimulateNetworkDelay(sources)
    
    fmt.Println("\n第二次运行:")
    demo.SimulateNetworkDelay(sources)
    
    fmt.Println("\n💡 每次运行的延迟都不同，真实模拟网络环境！")
}ne("Golang 并发编程")
}
