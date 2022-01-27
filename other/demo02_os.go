package main

import (
	"fmt"
	"os"
)

func main() {
	//获得当前工作路径：默认当前工程根目录
	dir, err := os.Getwd()
	fmt.Println(dir)
	fmt.Println(err)


	//获得指定的环境变量
	//paths := os.Getenv("Path")
	//goroot := os.Getenv("GOROOT")
	//fmt.Println(paths)//.....
	//fmt.Println(goroot)//c:/go

	//获得所有环境变量
	envs := os.Environ()
	for _,env := range envs{
		fmt.Println(env)
	}



}
