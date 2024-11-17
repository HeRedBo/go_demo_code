package main

import (
	"strconv"
	"sync"
)

var waitGroup sync.WaitGroup

func main031() {
	baseUrl := "https://www.3gbizhi.com/meinv/index"
	pageLen := 3
	var url string
	for i := 2; i <= pageLen; i++ {
		url = baseUrl + "_" + strconv.Itoa(i) + ".html"
		DownloadPageImg(url)
	}
	waitGroup.Wait()
}

func DownloadPageImg(url string) {
	imageInfos := GetPageImagesInfos(url)
	for _, imageMap := range imageInfos {
		DownloadImgAsync2(imageMap["url"], imageMap["filename"], &waitGroup)
	}
}

func DownloadImgAsync2(url, filename string, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		chSem <- 123
		DownloadImg(url, filename)
		<-chSem
		wg.Done()
	}()
}
