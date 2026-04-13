package main

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
    fmt.Println("\n3️⃣  使用时间种子package main

import (
	"fmt"
	"sync"
)

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func fanOut(in <-chan int, workers int) []<-chan int {
	channels := make([]<-chan int, workers)
	for i := 0; i < workers; i++ {
		channels[i] = square(in)
	}
	return channels
}

func fanIn(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	merged := make(chan int)

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			merged <- n
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

func main2() {
	in := generate(2, 3, 4, 5, 6, 7, 8)
	workers := fanOut(in, 3)
	out := fanIn(workers...)
	for n := range out {
		fmt.Println(n)
	}
}
