package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
√ 接口：只有方法的定义，没有实现——全部是抽象方法
√ 实现接口：结构体实现接口的全部抽象方法，就称为结构体实现了接口
√ 多态：一个父类/接口有不同的子类实现，本例中【劳动者接口】的具体实现有【程序员】、【产品经理】、【老板】
√ 共性：【程序员】、【产品经理】、【老板】都会劳动和休息
√ 个性：【程序员】、【产品经理】、【老板】的劳动和休息方式各不相同
*/

/**
劳动者父类接口
内含两个抽象方法，工作、休息
*/
type Worker interface {
	//每天工作多少小时，产出何种产品
	Work(hour int) (product string)
	// 休息
	Rest()
}

/**
定义码农结构体
*/
type Coder struct {
	skill string
}

/*码农指针实现Worker接口*/
func (c *Coder) Work(hour int) (product string) {
	fmt.Printf("码农一天工作%d小时\n", hour)
	fmt.Printf("码农正在%s\n", c.skill)
	return "BUG"
}

func (c *Coder) Rest() {
	fmt.Println("休息是什么？？？")
}

/** 程序员特有的方法 */
func (c *Coder) WorkHome() {
	fmt.Println("程序员在家工作")
}

/**
定义产品经理结构体
*/
type ProductManager struct {
	skill string
}

func (pm *ProductManager) Work(hour int) (product string) {
	fmt.Printf("产品一天工作%d小时\n", hour)
	fmt.Printf("产品正在%s\n", pm.skill)
	return "无逻辑产品"
}

func (pm *ProductManager) Rest() {
	fmt.Println("看程序员撸代码")
}

/** 定义老板结构体 */
type Boss struct {
	skill string
}

func (b *Boss) Work(hour int) (product string) {
	fmt.Printf("老板一天工作%d小时\n", hour)
	fmt.Printf("产品正在%s\n", b.skill)
	return "梦想"
}

func (c *Boss) Rest() {
	fmt.Println("一天到晚都在休息")
}

func main() {
	//创建由Worker组成的切片
	workers := make([]Worker, 0)

	workers = append(workers, &Coder{"撸代码"})
	workers = append(workers, &ProductManager{"拍脑门"})
	workers = append(workers, &Boss{"吹牛逼"})

	fmt.Println(workers)

	/*随机生成今天星期几*/
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	weekday := r.Intn(7)
	fmt.Printf("今天是星期%d\n", weekday)

	//工作日全体工作
	if weekday > 0 && weekday < 6 {
		//全体工作
		for _, worker := range workers {
			//调度共性，无视具体类型
			worker.Work(8)
		}
	} else {
		//全体休息
		for _, worker := range workers {
			worker.Rest()
		}

		//码农在家工作其他人休息
		for _, worker := range workers {
			//对接口实例做类型断言（类型判断）
			//判断当前worker是否是Coder指针
			if coder, ok := worker.(*Coder); ok {
				//调度个性
				coder.WorkHome()
			} else {
				worker.Rest()
			}
		}
	}
}
