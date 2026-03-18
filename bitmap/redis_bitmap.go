package bitmap

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisBitmap struct {
	client    *redis.Client
	ctx       context.Context
	maxOffset uint64 // 业务允许的最大offset（防御编程）
}

// NewRedisBitmap 创建带最大偏移限制的RedisBitmap
// maxOffset: 业务最大ID，超过直接拒绝，防止Redis异常膨胀
func NewRedisBitmap(client *redis.Client, ctx context.Context, maxOffset uint64) *RedisBitmap {
	return &RedisBitmap{
		client:    client,
		ctx:       ctx,
		maxOffset: maxOffset,
	}
}

func (r *RedisBitmap) checkOffset(offset uint64) error {
	if offset > r.maxOffset {
		return fmt.Errorf("offset %d exceeds max allowed %d", offset, r.maxOffset)
	}
	return nil
}

// Set 标记位为1（打卡/上线）
func (r *RedisBitmap) Set(key string, offset uint64) (int64, error) {
	if err := r.checkOffset(offset); err != nil {
		return 0, err
	}
	return r.client.SetBit(r.ctx, key, int64(offset), 1).Result()
}

// Unset 标记位为0（取消/离线）
func (r *RedisBitmap) Unset(key string, offset uint64) (int64, error) {
	if err := r.checkOffset(offset); err != nil {
		return 0, err
	}
	return r.client.SetBit(r.ctx, key, int64(offset), 0).Result()
}

// IsSet 查询位状态
func (r *RedisBitmap) IsSet(key string, offset uint64) (bool, error) {
	if err := r.checkOffset(offset); err != nil {
		return false, err
	}
	val, err := r.client.GetBit(r.ctx, key, int64(offset)).Result()
	if err != nil {
		return false, err
	}
	return val == 1, nil
}

// Count 统计为1的位数
func (r *RedisBitmap) Count(key string, start, end int64) (int64, error) {
	// 构造 BitCount 参数：指定统计的字节范围（start/end）
	bitCount := &redis.BitCount{
		Start: start,
		End:   end,
	}
	return r.client.BitCount(r.ctx, key, bitCount).Result()
}

// BitOp 位运算：AND/OR/XOR/NOT
func (r *RedisBitmap) BitOp(op, destKey string, srcKeys ...string) (int64, error) {
	if op == "" || destKey == "" || len(srcKeys) == 0 {
		return 0, errors.New("invalid parameters")
	}

	// 构造 Redis BITOP 命令的完整参数：["BITOP", op, destKey, srcKey1, srcKey2, ...]
	args := []interface{}{"BITOP", op, destKey}
	for _, key := range srcKeys {
		args = append(args, key)
	}
	// 手动执行 Redis 原始命令（解决 BitOp 方法未定义问题）
	// 核心修复：用 Do 方法执行原生 Redis 命令（所有 go-redis 版本都支持）
	// Do 方法返回 Redis 命令的原始结果，再通过 Int64() 解析
	res, err := r.client.Do(r.ctx, args...).Int64()
	if err != nil {
		return 0, fmt.Errorf("execute BITOP command failed: %w", err)
	}
	return res, nil
}
