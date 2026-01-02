package samber

import (
	"fmt"
	"github.com/gookit/goutil/dump"
	"github.com/samber/lo"
)

// 定义结构体（模拟 PHP 多维数组）
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func ArrayDemo() {
	users := []User{
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 2, Name: "Bob", Age: 25},
		{ID: 3, Name: "Charlie", Age: 35},
	}

	// array_column(users, 'Name') → Map 提取字段
	names := lo.Map(users, func(u User, _ int) string {
		return u.Name
	})
	dump.P("Names:", names)
	fmt.Println("Names:", names) // [Alice Bob Charlie]

	// array_filter → Filter
	young := lo.Filter(users, func(u User, _ int) bool {
		return u.Age < 30
	})
	dump.P("Young:", young)      // [{{2 Bob 25}}]
	fmt.Println("Young:", young) // [{{2 Bob 25}}]

	// array_keys / array_values（针对 map）
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	keys := lo.Keys(m)     // []string{"a", "b", "c"}（顺序不定）
	values := lo.Values(m) // []int{1, 2, 3}（顺序不定）
	dump.P("Keys:", keys)
	dump.P("Values:", values)
	fmt.Println("Keys:", keys)
	fmt.Println("Values:", values)
}
