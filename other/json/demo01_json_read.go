package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gookit/goutil/dump"
)

func main01() {
	// 方式一：使用通用的map
	// 读取JSON文件
	fileContent, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Printf("无法读取文件：%v\n", err)
		return
	}

	dump.P(fileContent)

	var data map[string]interface{}
	if err := json.Unmarshal(fileContent, &data); err != nil {
		fmt.Printf("JSON解析失败：%v\n", err)
		return
	}
	// 输出解析后的数据
	fmt.Println("解析后的JSON数据：")
	fmt.Printf("%+v\n", data)
	// 示例：访问特定字段
	if name, exists := data["product_info"].(string); exists {
		fmt.Printf("姓名：%s\n", name)
	}

	jsonData, err2 := json.Marshal(data)
	if err2 != nil {
		fmt.Printf("JSON序列化失败：%v\n", err2)
	}

	dump.P(string(jsonData))
}

// 定义与JSON结构匹配的结构体
type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
}

func main() {
	// 读取JSON文件
	fileContent, err := os.ReadFile("person.json")
	if err != nil {
		fmt.Printf("无法读取文件：%v\n", err)
		return
	}

	// 解析JSON数据到结构体
	var person Person
	if err := json.Unmarshal(fileContent, &person); err != nil {
		fmt.Printf("JSON解析失败：%v\n", err)
		return
	}

	dump.Println(person)
	// 输出解析后的数据
	fmt.Printf("姓名：%s\n", person.Name)
	fmt.Printf("年龄：%d\n", person.Age)
	fmt.Printf("爱好：%v\n", person.Hobbies)
}
