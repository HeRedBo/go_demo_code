package gooop

import "fmt"

/**
嵌入与组合（Embedding）
*/

// 基类（通过嵌入实现类似继承）
type Animal struct {
	name string
}

func (a *Animal) Speak() {
	fmt.Println("Animal sound")
}

func (a *Animal) GetName() string {
	return a.name
}

// 通过嵌入实现组合
type Dog struct {
	Animal // 嵌入Animal，获得其所有方法和字段
	breed  string
}

// 方法重写（不是真正重写，Go没有继承）
//func (d *Dog) Speak() {
//	fmt.Println("Woof! Woof!")
//}

func main4() {
	dog := Dog{
		Animal: Animal{name: "Buddy"},
		breed:  "Golden Retriever",
	}

	dog.Speak()                // Woof! Woof!（调用Dog的方法）
	dog.Animal.Speak()         // Animal sound（可调用嵌入类型的方法）
	fmt.Println(dog.GetName()) // Buddy
}
