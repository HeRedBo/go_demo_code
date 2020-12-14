package main

import "fmt"

func main111() {

	defer fmt.Println("疑是银河落九天")

	fmt.Println("日照香炉生紫烟")
	fmt.Println("遥看破布挂前川")
	fmt.Println("飞流直下三千尺")

	//return
}

func lastPassage()  {
	fmt.Println("疑是银河落九天")
	fmt.Println("——李白")
}

func main112() {
	defer lastPassage()

	fmt.Println("日照香炉生紫烟")
	fmt.Println("遥看破布挂前川")
	fmt.Println("飞流直下三千尺")

	//return
}

func main113() {
	defer func() {
		fmt.Println("疑是银河落九天")
		fmt.Println("——李白")
	}()

	fmt.Println("日照香炉生紫烟")
	fmt.Println("遥看破布挂前川")
	fmt.Println("飞流直下三千尺")

	//return
}

func main114() {

	fmt.Println("打开网络")
	defer fmt.Println("关闭网络")

	fmt.Println("打开文件")
	defer fmt.Println("关闭文件")

	fmt.Println("打开数据库")
	defer fmt.Println("关闭数据库")

	fmt.Println("上课")
	fmt.Println("撸代码")
	fmt.Println("做练习")

	//fmt.Println("关灯")
	//fmt.Println("关门")
}
