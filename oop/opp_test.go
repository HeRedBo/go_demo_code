package opp

import "testing"

func TestO1Class(t *testing.T) {
	//创建空白的Person对象（object）/实例(instance)
	rangge := Person{}
	//设置其属性
	rangge.name = "西门阿让"
	//访问其方法
	rangge.Eat()
	rangge.Drink()
	rangge.Love()
}
