package gooop

import (
	"github.com/gookit/goutil/dump"
	"testing"
)

func TestClass(t *testing.T) {
	// 创建实例
	p1 := Person{name: "Alice", age: 25}
	p2 := &Person{name: "Bob", age: 30}

	dump.P(p1.GetName()) // Alice
	p2.SetName("Bob Updated")
	dump.P(p2.GetName()) // Bob Updated
}

func TestInterface(t *testing.T) {
	shapes := []Shape{
		Rectangle{width: 3, height: 4},
		Circle{radius: 5},
	}

	for _, shape := range shapes {
		PrintShapeInfo(shape)
	}
}

func TestAccount(t *testing.T) {
	var accont = NewAccount("bo")
	accont.Deposit(100)
	dump.P(accont.GetBalance())
}

func TestEmbedding(t *testing.T) {
	dog := Dog{
		Animal: Animal{name: "Buddy"},
		breed:  "Golden Retriever",
	}

	dog.Speak()                 // Woof! Woof!（调用Dog的方法）
	dog.Animal.Speak()          // Animal sound（可调用嵌入类型的方法）
	dump.Println(dog.GetName()) // Buddy
}
func TestFactory(t *testing.T) {
	var mysqlDb = CreateDatabase("mysql")
	dump.P(mysqlDb.Connect())
	var PgDb = CreateDatabase("postgres")
	dump.P(PgDb.Connect())
}

func TestSing(t *testing.T) {
	var singleton = GetInstance()
	dump.P(singleton)
	var singleton2 = GetInstance()
	dump.P(singleton2)
}
