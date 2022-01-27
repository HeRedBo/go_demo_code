package main

import (
	"fmt"
	"helloGO/concurrence_2/day5/spider-2.0/spiderkit"
	"strconv"
	"sync"
	"time"
)

var (
	chSem = make(chan int, 10)
	chaImgMaps = make(chan map[string]string, 1000)
	imgDir = `D:\www\Go\go_demo\concurrence_2\day5\spider-2.0\images\`
	wg4ImageInfo sync.WaitGroup
	wg4Download  sync.WaitGroup
	dirPath = `concurrence_2\day5\spider-2.0\images2\`
)

func main() {
	baseUrl := `https://www.tupianzj.com/meinv/xinggan/list_176`
	imageSavePath := spiderkit.GetImagePath(dirPath)
	// 4 页
	var pageSize = 11
	start := time.Now()
	for i := 1 ; i <= pageSize ; i ++ {
		var url string
		url = baseUrl + "_" + strconv.Itoa( i )  + ".html"
		//imagesData := spiderkit.GetPageImagesInfos(url, imgDir)
		//fmt.Println(imagesData)
		wg4ImageInfo.Add(1)
		go func(theUrl , imgDir string, chaImgMaps chan<- map[string]string ) {
			spiderkit.GetPageImageInfos2Chan(theUrl, imgDir, chaImgMaps)
			wg4ImageInfo.Done()
		}(url, imageSavePath, chaImgMaps)
	}

	go func() {
		wg4ImageInfo.Wait()
		close(chaImgMaps)
	}()

	for imaMap := range chaImgMaps {
		wg4Download.Add(1)
		go func(im map[string]string) {
			chSem<- 123
			// 下载数据
			spiderkit.DownloadFileWithClient(im["url"], im["filename"])
			<-chSem
			wg4Download.Done()
		}(imaMap)
	}
	wg4Download.Wait()
	consume := time.Now().Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
