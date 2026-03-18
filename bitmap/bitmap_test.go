package bitmap

import (
	"context"
	"fmt"
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/redis/go-redis/v9"
)

func TestBitmap(t *testing.T) {
	// 实战：记录用户打卡状态
	clockInBitmap := NewBitmap()

	// 标记用户10、5、88已打卡
	clockInBitmap.Set(10)
	clockInBitmap.Set(5)
	clockInBitmap.Set(88)
	// 查询用户5是否打卡
	fmt.Println("用户5是否打卡：", clockInBitmap.IsSet(5)) // 输出：true

	// 查询用户99是否打卡
	dump.Println("用户99是否打卡：", clockInBitmap.IsSet(99)) // 输出：false

	// 标记用户10取消打卡
	clockInBitmap.Unset(10)
	dump.Println("用户10是否打卡：", clockInBitmap.IsSet(10)) // 输出：false

	// 统计打卡总人数
	dump.Println("打卡总人数：", clockInBitmap.Count()) // 输出：2（5和88）
}

func TestFixedBitmap(t *testing.T) {
	// 1. 创建Bitmap实例（用于记录用户打卡状态）
	bm := NewFixedBitmap(100)
	_ = bm.Set(5)
	_ = bm.Set(10)
	_ = bm.Set(2)
	isSet, _ := bm.IsSet(5)
	count := bm.Count()

	dump.P(isSet, count)

	//// 4. 取消用户10的打卡状态
	//clockInBitmap.Unset(10)
	//fmt.Println("用户10是否打卡：", clockInBitmap.IsSet(10)) // 输出：false
	//
	//// 5. 统计打卡总人数
	//fmt.Println("打卡总人数：", clockInBitmap.Count()) // 输出：2（5和100）
}

func TestRedisBitmap(t *testing.T) {
	// 1. 初始化Redis客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // 你的Redis地址
		Password: "123456",         // Redis密码（无则空）
		DB:       0,                // 使用的数据库编号
	})

	// 2. 检查Redis连接
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Redis连接失败：%v", err))
	}
	dump.Println("Redis连接成功")

	// 3. 创建Redis Bitmap实例（最大用户ID 1000000）
	bitmap := NewRedisBitmap(rdb, ctx, 1000000)

	// 4. 测试核心功能
	clockInKey := "clock_in:20260317"
	// 标记用户5打卡
	// 标记用户5打卡
	prevVal, err := bitmap.Set(clockInKey, 5)
	if err != nil {
		fmt.Printf("标记用户5打卡失败：%v\n", err)
	} else {
		fmt.Printf("用户5打卡前状态：%d（0=未打卡，1=已打卡）\n", prevVal)
	}

	// 查询用户5是否打卡
	isSet, err := bitmap.IsSet(clockInKey, 5)
	if err != nil {
		fmt.Printf("查询用户5打卡状态失败：%v\n", err)
	} else {
		fmt.Printf("用户5是否打卡：%t\n", isSet) // 输出 true
	}

	// 统计打卡人数（0,-1 表示统计所有位）
	count, err := bitmap.Count(clockInKey, 0, -1)
	if err != nil {
		fmt.Printf("统计打卡人数失败：%v\n", err)
	} else {
		fmt.Printf("20260317 打卡总人数：%d\n", count) // 输出 1
	}

	// 5. 测试BitOp（合并两天打卡记录）
	clockInKey2 := "clock_in:20260318"
	bitmap.Set(clockInKey2, 5)
	bitmap.Set(clockInKey2, 10)

	// 计算两天都打卡的用户（AND运算）
	destKey := "clock_in:20260317_18_both"
	_, err = bitmap.BitOp("AND", destKey, clockInKey, clockInKey2)
	if err != nil {
		fmt.Printf("位图运算失败：%v\n", err)
	} else {
		// 统计两天都打卡的人数
		bothCount, _ := bitmap.Count(destKey, 0, -1)
		fmt.Printf("两天都打卡的人数：%d\n", bothCount) // 输出 1（仅用户5）
	}

	// 6. 取消用户5打卡
	prevVal, err = bitmap.Unset(clockInKey, 5)
	if err != nil {
		fmt.Printf("取消用户5打卡失败：%v\n", err)
	} else {
		fmt.Printf("用户5取消打卡前状态：%d\n", prevVal) // 输出 1
	}
}
