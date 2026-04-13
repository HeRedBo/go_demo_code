package main

import (
	"fmt"
	"testing"
	"time"
)

// TestWithoutSeed 测试不设置种子的行为
func TestWithoutSeed(t *testing.T) {
	demo := &RandomSeedDemo{}

	result1 := demo.WithoutSeed(10)
	result2 := demo.WithoutSeed(10)

	t.Logf("第一次调用: %v", result1)
	t.Logf("第二次调用: %v", result2)

	// Go 1.20+ 会自动初始化，所以两次结果应该不同
	if fmt.Sprint(result1) == fmt.Sprint(result2) {
		t.Log("⚠️  两次结果相同（这在 Go 1.20 之前是正常行为）")
	} else {
		t.Log("✅ 两次结果不同（Go 1.20+ 自动初始化种子）")
	}
}

// TestWithFixedSeed 测试固定种子的可重现性
func TestWithFixedSeed(t *testing.T) {
	demo := &RandomSeedDemo{}

	seed := int64(42)
	result1 := demo.WithFixedSeed(seed, 10)
	result2 := demo.WithFixedSeed(seed, 10)
	result3 := demo.WithFixedSeed(seed, 10)

	t.Logf("种子: %d", seed)
	t.Logf("第一次调用: %v", result1)
	t.Logf("第二次调用: %v", result2)
	t.Logf("第三次调用: %v", result3)

	// 固定种子应该产生完全相同的结果
	if fmt.Sprint(result1) != fmt.Sprint(result2) {
		t.Errorf("固定种子应该产生相同结果，但得到: %v vs %v", result1, result2)
	}

	if fmt.Sprint(result2) != fmt.Sprint(result3) {
		t.Errorf("固定种子应该产生相同结果，但得到: %v vs %v", result2, result3)
	}

	t.Log("✅ 固定种子成功产生可重现的随机数序列")
}

// TestWithDifferentFixedSeeds 测试不同固定种子的差异
func TestWithDifferentFixedSeeds(t *testing.T) {
	demo := &RandomSeedDemo{}

	result1 := demo.WithFixedSeed(42, 10)
	result2 := demo.WithFixedSeed(123, 10)
	result3 := demo.WithFixedSeed(999, 10)

	t.Logf("种子 42:  %v", result1)
	t.Logf("种子 123: %v", result2)
	t.Logf("种子 999: %v", result3)

	// 不同种子应该产生不同的结果
	if fmt.Sprint(result1) == fmt.Sprint(result2) {
		t.Error("不同种子应该产生不同结果")
	}

	if fmt.Sprint(result2) == fmt.Sprint(result3) {
		t.Error("不同种子应该产生不同结果")
	}

	t.Log("✅ 不同种子产生不同的随机数序列")
}

// TestWithAutoSeed 测试自动种子（Go 1.20+）
func TestWithAutoSeed(t *testing.T) {
	demo := &RandomSeedDemo{}

	result1 := demo.WithAutoSeed(10)
	result2 := demo.WithAutoSeed(10)

	t.Logf("第一次调用: %v", result1)
	t.Logf("第二次调用: %v", result2)

	// 由于自动种子，两次调用可能不同也可能相同（概率很低）
	// 这里主要验证程序不会崩溃
	t.Log("✅ 自动种子工作正常")
}

// TestWithRandV2 测试新的 rand/v2 包
func TestWithRandV2(t *testing.T) {
	demo := &RandomSeedDemo{}

	result1 := demo.WithRandV2(10)
	result2 := demo.WithRandV2(10)

	t.Logf("第一次调用: %v", result1)
	t.Logf("第二次调用: %v", result2)

	// 验证生成的数字在合理范围内
	for _, num := range result1 {
		if num < 0 || num >= 100 {
			t.Errorf("随机数超出范围: %d", num)
		}
	}

	t.Log("✅ rand/v2 工作正常")
}

// TestSimulateNetworkDelay 测试网络延迟模拟
func TestSimulateNetworkDelay(t *testing.T) {
	demo := &RandomSeedDemo{}

	sources := []string{"Google", "Bing", "Baidu"}

	delays1 := demo.SimulateNetworkDelay(sources)
	delays2 := demo.SimulateNetworkDelay(sources)

	t.Log("第一次模拟:")
	for source, delay := range delays1 {
		t.Logf("  %s: %v", source, delay)
	}

	t.Log("第二次模拟:")
	for source, delay := range delays2 {
		t.Logf("  %s: %v", source, delay)
	}

	// 验证所有延迟都在合理范围内 (100ms - 500ms)
	for source, delay := range delays1 {
		if delay < 100*time.Millisecond || delay > 500*time.Millisecond {
			t.Errorf("%s 的延迟超出范围: %v", source, delay)
		}
	}

	t.Log("✅ 网络延迟模拟工作正常")
}

// BenchmarkWithoutSeed 性能测试：不设置种子
func BenchmarkWithoutSeed(b *testing.B) {
	demo := &RandomSeedDemo{}
	for i := 0; i < b.N; i++ {
		demo.WithoutSeed(100)
	}
}

// BenchmarkWithFixedSeed 性能测试：固定种子
func BenchmarkWithFixedSeed(b *testing.B) {
	demo := &RandomSeedDemo{}
	for i := 0; i < b.N; i++ {
		demo.WithFixedSeed(42, 100)
	}
}

// BenchmarkWithAutoSeed 性能测试：自动种子
func BenchmarkWithAutoSeed(b *testing.B) {
	demo := &RandomSeedDemo{}
	for i := 0; i < b.N; i++ {
		demo.WithAutoSeed(100)
	}
}

// BenchmarkWithRandV2 性能测试：rand/v2
func BenchmarkWithRandV2(b *testing.B) {
	demo := &RandomSeedDemo{}
	for i := 0; i < b.N; i++ {
		demo.WithRandV2(100)
	}
}
