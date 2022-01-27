package spiderkit

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

/*生成时间戳_随机数文件名*/
func GetRandomName() string {
	timestamp := strconv.Itoa(int(time.Now().UnixNano()))
	randomNum := strconv.Itoa(GetRandomInt(1000, 10000))
	return timestamp + "_" + randomNum
}


// 获取文件存储目录地址
func GetImagePath(dirName string) string {
	dirPath := GetAbsolutePath(dirName)
	check , _ := IsFileExist(dirPath)
	if !check {
		os.Mkdir(dirPath, os.ModePerm)
	}
	return dirPath
}

func GetAbsolutePath(path string) string{
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return dir +`\` + path


}

/**
@param path string 路径
@return bool  error
 */
func IsFileExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if(os.IsNotExist(err)) {
		return false,nil
	}
	if err == nil {
		return true, nil
	}
	return false, err
}
