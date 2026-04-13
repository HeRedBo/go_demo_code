package main

import (
	"fmt"
	"math/rand"
	mrand "math/rand/v2"
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
		results[i] = mrand.IntN(100)
	}
	return results
}

// CompareSeeds 对比不同种子设置的效果
func (d *RandomSeedDemo) CompareSeeds() {
	fmt.Println("======================================================================")
	fmt.Println("🎲 随机数种子对比演示")
	fmt.Println("======================================================================")

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

	fmt.Println("\n======================================================================")
	fmt.Println("📝 总结:")
	fmt.Println("   • Go 1.20 之前: 必须手动调用 rand.Seed()")
	fmt.Println("   • Go 1.20-1.21: 自动初始化，但仍可使用 rand.Seed()")
	fmt.Println("   • Go 1.22+: 推荐使用 math/rand/v2 包")
	fmt.Println("   • 固定种子用于测试，时间/自动种子用于生产")
	fmt.Println("======================================================================")
}

// SimulateNetworkDelay 模拟网络延迟（类似 fanout-fanin 场景）
func (d *RandomSeedDemo) SimulateNetworkDelay(sources []string) map[string]time.Duration {
	delays := make(map[string]time.Duration)

	fmt.Printf("\n🌐 模拟 %d 个搜索引擎的网络延迟:\n", len(sources))
	for _, source := range sources {
		delay := time.Duration(mrand.IntN(400)+100) * time.Millisecond
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
}
