package main

import "fmt"

type Person struct {
	//  封装结构体属性
	name string
	age int
	sex bool
	hobby []string
}

/*
封装结构体的方法
- 无论方法的主语定义为值类型还是指针类型，对象值和对象指针都能够正常访问
- 通常会将主语定义为指针类型——毕竟西门鹤的副本吃了饭，肉不会长到西门鹤本人身上去
*/

func (p *Person) Eat() {
	fmt.Printf("%s爱饕餮\n", p.name)
}

func (p *Person) Drink() {
	fmt.Printf("%s爱喝酒\n", p.name)
}

func (p *Person) Love() {
	fmt.Printf("%s是有感情的\n", p.name)
	p.age  -=1
}

func (p *Person) SelfIntroduce() {
	fmt.Printf("我是%s, 今年%d岁了，", p.name, p.age)
	if p.sex {
		fmt.Println("性别：男")
	} else  {
		fmt.Println("性别：女")
	}
}

//值传递传递的是对象的副本
func MakeHimLove(p Person) {
	fmt.Printf("p的地址是%p\n",&p)
	p.Love()
}

//引用传递传递的是对象的真身
func MakeHisPtrLove(p *Person) {
	fmt.Printf("p的地址是%p\n",p)
	p.Love()
}



/*创建空白对象并访问属性和方法*/
func main0012() {
	//创建空白的Person对象（object）/实例(instance)
	rangge := Person{}

	//设置其属性
	rangge.name = "西门阿让"

	//访问其方法
	rangge.Eat()
	rangge.Drink()
	rangge.Love()
}


func main0113() {
	rangge := Person{name:"西门阿让",sex:true,age:8}
	rangge.name = "西门阿让"

	rangge.SelfIntroduce()
	rangge.Eat()
	rangge.Drink()
	rangge.Love()
}


func main014() {
	rangge := Person{"西门阿让", 8, true, []string{"撸代码", "完美的撸代码"}}
	fmt.Printf("让哥的真身地址是%p\n",&rangge)

	////要求传递值就必须传递值
	//MakeHimLove(rangge)
	//要求传递指针就必须传递指针（指针/地址/引用传递）
	//MakeHisPtrLove(&rangge)

	//值传递传递的是副本，引用传递传递的才是真身
	for i := 0; i < 7; i++ {
		//MakeHimLove(rangge)
		MakeHisPtrLove(&rangge)
	}
	fmt.Printf("暴风雨过后让哥的年龄是%d\n",rangge.age)
}



/*使用值和指针访问方法，效果是一致的*/
func main015() {
	p := &Person{age: 10}
	p.Love()
	fmt.Println(p.age)
}

/*new的参数是结构体的名字,返回的是所有属性都为默认值的对象指针*/
func main016() {
	//new的参数是结构体的名字,返回的是所有属性都为默认值的对象指针
	personPtr := new(Person)
	fmt.Println(personPtr)

	var pp *Person
	fmt.Println(pp)
}

func main() {
	p := Person{"西门阿让", 8, false, []string{"撸代码", "完美的撸代码"}}
	p.SelfIntroduce()
}
