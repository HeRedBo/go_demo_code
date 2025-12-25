package gotime

import (
	"fmt"
	"github.com/gookit/goutil/dump"
	"time"
)

// 时间基础类型
func TimeBase() {
	// time.Time 类型表示时间点
	var t time.Time
	fmt.Println("零值时间:", t) // 0001-01-01 00:00:00 +0000 UTC
	dump.P("零值时间:", t)

	// 判断时间是否为零值
	fmt.Println("是否零值:", t.IsZero())
	dump.P("是否零值:", t.IsZero())
}

// CreateTimeExamples 获取当前时间
func CreateTimeExamples() {
	// 当前本地时间
	now := time.Now()
	fmt.Printf("当前本地时间: %v\n", now)
	dump.P("当前本地时间: %v\n", now)

	// 当前UTC时间
	utcNow := time.Now().UTC()
	fmt.Printf("当前UTC时间: %v\n", utcNow)
	dump.P("当前UTC时间: %v\n", utcNow)

	// 自定义时间
	customTime := time.Date(2023, time.December, 25, 10, 30, 45, 0, time.Local)
	fmt.Printf("自定义时间: %v\n", customTime)

	// 解析时间字符串
	parsedTime, err := time.Parse("2006-01-02", "2023-12-25")
	if err != nil {
		fmt.Println("解析错误:", err)
		dump.P("解析错误:", err)
	} else {
		fmt.Printf("解析的时间: %v\n", parsedTime)
		dump.P("解析的时间: %v\n", parsedTime)
	}
}

// FormatExamples 时间格式化
// Go使用特定的参考时间进行格式化：2006-01-02 15:04:05
func FormatExamples() {
	now := time.Now()

	// 常用格式化模式
	fmt.Println("RFC3339格式:", now.Format(time.RFC3339))
	fmt.Println("自定义格式1:", now.Format("2006-01-02"))
	fmt.Println("自定义格式1-2:", now.Format("2006-01-02 15:04:05"))
	fmt.Println("自定义格式2:", now.Format("2006/01/02 15:04:05"))
	fmt.Println("自定义格式3:", now.Format("2006年01月02日 15时04分05秒"))
	fmt.Println("自定义格式4:", now.Format("15:04:05"))
	fmt.Println("自定义格式5:", now.Format("Monday, 02-Jan-06 15:04:05 MST"))
	fmt.Println("自定义格式6:", now.Format("2006-01-02T15:04:05Z07:00"))

	// 12小时制
	fmt.Println("12小时制:", now.Format("2006-01-02 03:04:05 PM"))

	// 只取日期部分
	fmt.Println("仅日期:", now.Format("2006-01-02"))

	// 时间戳字符串
	fmt.Println("Unix时间:", now.Unix())
}
