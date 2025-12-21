package gooop

/**
 * 1.1 结构体（Struct）作为"类"
 * Go没有"类"的概念，使用结构体作为数据载体
 */

// 定义结构体
type Person struct {
	name string
	age  int
}

// 方法定义（值接收者）
func (p Person) GetName() string {
	return p.name
}

// 方法定义（指针接收者）
func (p *Person) SetName(name string) {
	p.name = name
}
