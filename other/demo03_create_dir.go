package main

import (
	"fmt"
	"os"
)

/**
创建文件夹 判断文件夹 是否存在
在获取当前执行文件的执行路径 并在当前文同级 创建 指定文件的文件
*/

func main() {

	// 获取当前文件的执行路径
	//获得当前工作路径：默认当前工程根目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dirName := `other\images`
	dirPath := dir + `\` +dirName


	// 创建目录
	_, err2 := os.Stat(dirPath)
	if os.IsNotExist(err2) {
		os.Mkdir(dirPath, os.ModePerm)
		//fmt.Println("目录已存在",err2)
	}
}
