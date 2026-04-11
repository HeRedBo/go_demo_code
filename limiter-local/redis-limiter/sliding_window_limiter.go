package redislimiter

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// SlidingWindowLimiter 滑动窗口限流器（Redis版）
type SlidingWindowLimiter struct {
	client    *redis.Client
	prefix    string        // Redis key前缀
	threshold int64         // 窗口内最大请求数
	window    time.Duration // 窗口大小
}

// NewSlidingWindowLimiter 创建滑动窗口限流器
func NewSlidingWindowLimiter(threshold int64, window time.Duration) *SlidingWindowLimiter {
	// 初始化Redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // 如有密码请填写
		DB:       0,
	})
	// 测试连接
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic("Redis连接失败：" + err.Error())
	}

	return &SlidingWindowLimiter{
		client:    client,
		prefix:    "limit:sliding:",
		threshold: threshold,
		window:    window,
	}
}

// AllowRequest 检查是否允许请求
func (sw *SlidingWindowLimiter) AllowRequest(key string) bool {
	ctx := context.Background()
	redisKey := sw.prefix + key
	now := time.Now().UnixMilli() // 当前时间戳（毫秒）
	windowStart := now - sw.window.Milliseconds()

	// 1. 移除窗口外的请求记录
	_, err := sw.client.ZRemRangeByScore(ctx, redisKey, "0", fmt.Sprintf("%d", windowStart)).Result()
	if err != nil {
		fmt.Printf("清理过期数据失败：%v\n", err)
		return false
	}

	// 2. 统计当前窗口内请求数
	currentCount, err := sw.client.ZCard(ctx, redisKey).Result()
	if err != nil {
		fmt.Printf("统计请求数失败：%v\n", err)
		return false
	}

	// 3. 判断是否超过阈值
	if currentCount < sw.threshold {
		// 记录当前请求（用时间戳+随机数做唯一值）
		member := fmt.Sprintf("%d_%d", now, time.Now().UnixNano())
		_, err := sw.client.ZAdd(ctx, redisKey, redis.Z{Score: float64(now), Member: member}).Result()
		if err != nil {
			fmt.Printf("记录请求失败：%v\n", err)
			return false
		}
		// 设置过期时间
		sw.client.Expire(ctx, redisKey, sw.window+time.Second)
		return true
	}
	return false
}
