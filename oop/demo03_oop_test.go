package main

import (
	"fmt"
	"helloGO/oop/src/mylib"
)

/*
定义动物接口：死、活着
定义动物实现类：鸟、鱼、野兽（跑、捕食）
继承野兽：实现老虎，实现人
业务场景：工作日所有动物都活着、周末人出来捕食，野兽逃跑，其它动物死光光
*/
type Animal interface {
	Live()
	GoDie()
}

type Bird struct {}

func (b *Bird) Live()  {
	fmt.Println("一只鸟在唱歌")
}

func (b *Bird) GoDie()  {
	fmt.Println("一只鸟死翘翘")
}

type Fish struct {}

func (b *Fish) Live()  {
	fmt.Println("一条鱼在水里游")
}

func (b *Fish) GoDie()  {
	fmt.Println("一条鱼死翘翘")
}

type Beast struct {}

func (b *Beast) Live()  {
	fmt.Println("一直野兽在打滚")
}

func (b *Beast) Godie()  {
	fmt.Println("一只野兽死翘翘")
}

func (b *Beast) Run()  {
	fmt.Println("一直野兽在奔跑")
}

func (b *Beast) Hunt()  {
	fmt.Println("一只野兽在捕食")
}

type Tiger struct {
	Beast
}

func (b *Tiger) Hunt() {
	fmt.Println("嗷！本王要用膳了！！！")
}

type Human struct {
	Beast
	Name string
}

func (b *Human) Live() {
	fmt.Println("广东人：好好工作才有大餐吃")
}

func (b *Human) Hunt() {
	fmt.Println("广东人：该吃饭了呢~")
}

func main()  {
	// 创建各种动物实例
	bird := Bird{}
	fish := Fish{}
	tiger :=Tiger{}
	human := Human{}

	// 创建地球家园
	animals := make([]Animal, 0)
	animals = append(animals, &bird)
	animals = append(animals, &fish)
	animals = append(animals, &tiger)
	animals = append(animals, &human)

	weekday :=  mylib.GetRandomInt(6)
	fmt.Printf("今天星期%d\n", weekday)

	if weekday > 0 && weekday < 6 {
		for _, animal := range animals {
			animal.Live()
		}
	}else {
		for _, animal := range animals {

			//使用第二种类型断言
			switch animal.(type) {
			case *Human:
				//将人转换为野兽，并令其捕食
				animal.(*Human).Hunt()
			case *Tiger:
				animal.(*Tiger).Run()
			default:
				animal.GoDie()
			}

			//使用第二种类型断言
			//if human, ok := animal.(*Human); ok {
			//	human.Hunt()
			//} else if tiger, ok := animal.(*Tiger); ok {
			//	tiger.Run()
			//} else {
			//	animal.GoDie()
			//}
		}
	}






}





