package spiderkit

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"time"
)

const (
	DIAL_TIMEOUT = 10 * time.Second
	RW_TIMEOUT   = 10 * time.Second
)

var (
	downloadWG sync.WaitGroup
	chSem      = make(chan int, 100)
	httpClient http.Client
)

func init() {
	httpClient = http.Client{
		Transport:&http.Transport{
			Dial : func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw,addr, DIAL_TIMEOUT)
				if err != nil {
					return nil, err
				}
				deadline := time.Now().Add(RW_TIMEOUT)
				conn.SetDeadline(deadline)
				return conn, nil
			},
		},
	}
}



func GetUrlHtml(url string) string{
	resp, err := http.Get(url)
	HandleError(err, "Http.Get")
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	html := string(bytes)
	return html
}


/* 下载文件 */
func DownloadFileWithClient(url, filename  string) {
	//fmt.Println("DownloadFileWithClient...")
	resp, err := httpClient.Get(url)
	if err != nil {
		fmt.Println(filename, "下载失败！")
		return
	}
	defer resp.Body.Close()

	imgBytes, _ := ioutil.ReadAll(resp.Body)
	err = ioutil.WriteFile(filename, imgBytes, 0644)
	if err == nil {
		fmt.Println(filename, "下载成功！")
	} else {
		fmt.Println(filename, "下载失败！")
	}
}


func DownloadImgAsync(url, filename string) {
	downloadWG.Add(1)
	go func() {
		chSem <- 123
		DownloadImg(url, filename)
		<-chSem
		downloadWG.Done()
	} ()
	downloadWG.Wait()
}


func DownloadImg(url string , filename string) {
	fmt.Println("DownloadImg...")

	resp, err := http.Get(url)
	defer resp.Body.Close()

	imgBytes, _ := ioutil.ReadAll(resp.Body)
	err = ioutil.WriteFile( filename, imgBytes, 0644)
	if err == nil {
		fmt.Println(filename, "下载成功！")
	} else {
		fmt.Println("下载文件失败",err)
		fmt.Println(filename, "下载失败！")
	}
}
