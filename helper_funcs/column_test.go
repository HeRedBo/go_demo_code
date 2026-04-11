package helper_funcs

import (
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/samber/lo"
)

// 示例：结构体切片提取字段
type User struct {
	ID   int
	Name string
	Age  int
}

type Product struct {
	ID    int
	Price float64
}

func TestColumn(t *testing.T) {
	// 1. 结构体切片示例
	users := []User{
		{ID: 1, Name: "张三", Age: 20},
		{ID: 2, Name: "李四", Age: 25},
		{ID: 3, Name: "王五", Age: 30},
	}

	// 提取 Name 字段
	names, err := Column(users, "Name")
	if err != nil {
		dump.P("提取失败:", err)
		return
	}
	dump.P("提取的姓名列表:", names) // [张三 李四 王五]
	dump.P(names)

	// 提取 ID 字段
	ids, _ := Column(users, "ID")
	dump.P("提取的ID列表:", ids) // [1 2 3]
	dump.P(ids)
	// 2. 映射切片示例
	maps := []map[string]interface{}{
		{"id": 101, "title": "Go入门"},
		{"id": 102, "title": "PHP进阶"},
	}
	titles, _ := Column(maps, "title")
	dump.P("提取的标题列表:", titles) // [Go入门 PHP进阶]
	dump.P(maps)
}

func TestSamberIo(t *testing.T) {
	products := []Product{
		{ID: 1, Price: 99.9},
		{ID: 2, Price: 199.9},
	}
	//
	ids := lo.Map(products, func(p Product, _ int) int { return p.ID })
	// ✅ lo.Collect + lo.Map 正确用法（兼容所有泛型版本）
	prices := lo.Map(products, func(p Product, _ int) float64 { return p.Price })

	dump.P("产品价格:", prices) // [99.9 199.9]
	dump.P("产品ID:", ids)    // [1 2]
}

func TestSamberDemo(t *testing.T) {
	// 模拟 PHP 数组
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Charlie"},
	}

	// array_column
	names := lo.Map(users, func(u User, _ int) string { return u.Name })
	dump.P(names) // [Alice Bob Charlie]

	// array_values / array_keys (对 map)
	userMap := map[int]string{1: "Alice", 2: "Bob", 3: "Charlie"}
	keys := lo.Keys(userMap)
	vals := lo.Values(userMap)
	dump.P(keys, vals) // [1 2 3] [Alice Bob Charlie]

	// array_filter
	evenIDs := lo.Filter(users, func(u User, _ int) bool {
		return u.ID%2 == 0
	})
	dump.P(evenIDs) // [{2 Bob}]

	// array_reduce 拼接名字
	concatenated := lo.Reduce(users, func(agg string, u User, _ int) string {
		return agg + u.Name
	}, "")
	dump.P(concatenated) // AliceBobCharlie

	// in_array
	containsBob := lo.Contains(names, "Bob")
	dump.P(containsBob) // true

	// array_unique
	dup := []int{1, 2, 2, 3}
	uniq := lo.Uniq(dup)
	dump.P(uniq) // [1 2 3]

	// array_chunk
	chunked := lo.Chunk([]int{1, 2, 3, 4, 5}, 2)
	dump.P(chunked) // [[1 2] [3 4] [5]]
}
