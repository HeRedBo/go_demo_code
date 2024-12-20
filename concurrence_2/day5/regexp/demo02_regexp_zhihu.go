package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"regexp"
)

var savePath = "/Users/hehongbo/www/GO/go_demo_code/concurrence_2/day5/regexp/images"

func main0201() {
	var url = `https://www.3gbizhi.com/meinv`
	html := GetHtml(url)
	fmt.Println(html)
	var reLink = `<img[\s\S]+?src="(http[\s\S]+?)"[\s\S]*?>`
	re := regexp.MustCompile(reLink)
	rets := re.FindAllStringSubmatch(html, -1)
	fmt.Println(rets)
	for _, x := range rets {
		//fmt.Println(x[1])
		DownloadFileWithClient(x[1])
	}
}

func DownloadFileWithClient(url string) {
	//fmt.Println("DownloadFileWithClient...")
	resp, err := http.Get(url)
	_, filename := filepath.Split(url)
	if err != nil {
		fmt.Println(filename, "下载失败！")
		return
	}
	defer resp.Body.Close()

	imgBytes, _ := ioutil.ReadAll(resp.Body)
	err = ioutil.WriteFile(savePath+filename, imgBytes, 0644)
	if err == nil {
		fmt.Println(filename, "下载成功！")
	} else {
		fmt.Println(filename, "下载失败！")
	}

}
