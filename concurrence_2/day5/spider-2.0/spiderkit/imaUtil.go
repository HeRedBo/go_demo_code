package spiderkit

import (
	"fmt"
	"helloGO/concurrence_2/day5/spider-01/Common"
	"path/filepath"
	"regexp"
	"strings"
)


func GetPageImageInfos2Chan(url, imgDir string , chImaMaps chan<- map[string]string) {
	imageInfos := GetPageImagesInfos(url,imgDir)
	fmt.Println("imageInfo", imageInfos)
	for _, infoMap := range imageInfos {
		chImaMaps <- infoMap
	}
}


func GetPageImagesInfos(url string,  imgDir string) []map[string]string {
	html  := GetUrlHtml(url)
	html = string(Common.ConvertToByte(html, "gbk", "utf8"))
	re := regexp.MustCompile(Common.ReImage)
	rets := re.FindAllStringSubmatch(html, -1)
	imagesInfos := make([]map[string]string, 0)
	for _, ret := range rets {
		imgInfo := make(map[string]string)
		imgUrl := ret[1]
		imgInfo["url"] = imgUrl
		imgInfo["filename"] = GetImgNameFromTag(ret[0], imgUrl, imgDir)
		imagesInfos = append(imagesInfos, imgInfo)
	}
	return imagesInfos
}


/**
从<img>标签中提取文件名（含地址）
有 alt 使用alt 文件名， 没有使用时间戳_随机数做文件名
参数：
imgTag 图片<img>标签
imgDir 目录位置
suffix 文件名后缀
*/
func GetImgNameFromTag(imgTag, imgUrl, imgDir string) string {
	var filename string
	//fmt.Println(imgTag)
	//fmt.Println(imgUrl)
	//获取图片格式
	imgName := GetImgNameFromImageUrl(imgUrl)
	suffix := ".jpg"
	if imgName != "" {
		suffix = imgName[strings.LastIndex(imgName, "."):]
	}
	//fmt.Println(imgName)

	//尝试从imgTag中提取alt
	re := regexp.MustCompile(Common.ReAlt)
	rets := re.FindAllStringSubmatch(imgTag, 1)
	if len(rets) > 0 && imgName != "" {
		//首选alt
		alt := rets[0][1]
		alt = strings.Replace(alt, ":", "_", -1)
		filename = alt + imgName
	} else if imgName != "" {
		//次选链接中的文件名
		filename = imgName
	} else {
		//最末时间戳+随机数
		filename = GetRandomName() + suffix
	}
	filename = imgDir + filename
	return filename

}


func GetImgNameFromImageUrl(imgUrl string) string {
	_, filename := filepath.Split(imgUrl)
	if filename !=  "" {
		return filename
	} else {
		return ""
	}
}
