package main

import (
	"strconv"
	"sync"
)

var (
	wg4ImgInfo  sync.WaitGroup
	wg4Download sync.WaitGroup
	chImgMaps   = make(chan map[string]string, 1000)
)

func main() {
	baseUrl := "https://www.3gbizhi.com/meinv/index"
	pageLen := 4
	var url string
	for i := 1; i <= pageLen; i++ {
		url = baseUrl + "_" + strconv.Itoa(i) + ".html"
		wg4ImgInfo.Add(1)
		go func(theUrl string) {
			GetImgInfosFormPage(theUrl)
			wg4ImgInfo.Wait()
		}(url)
	}
	go func() {
		wg4ImgInfo.Wait()
		close(chImgMaps)
	}()

	// 下载图片信息 使用协程
	for imgMap := range chImgMaps {
		wg4Download.Add(1)
		go func(im map[string]string) {
			chSem <- 123
			DownloadImgWithClient(im["url"], im["filename"])
			<-chSem
			wg4Download.Done()
		}(imgMap)
	}
	wg4Download.Wait()
}

func GetImgInfosFormPage(url string) {
	imageInfos := GetPageImagesInfos(url)
	for _, imageMap := range imageInfos {
		chImgMaps <- imageMap
	}
}
