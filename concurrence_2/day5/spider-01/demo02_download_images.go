package main

import (
	"fmt"
	"github.com/gookit/goutil/dump"
	"helloGO/concurrence_2/day5/spider-01/Common"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	//url     = `https://www.keke345.cc/` // 已失效
	//url     = "https://www.tupianzj.com/meinv/xinggan/list_176" // 已失效
	url = `https://www.3gbizhi.com/meinv`
	//reImage = `<img src="(https.+?)".*?>`
	reImage = `<img.+?src="(http.+?)".*?>`
	//reImage = `<img\s+src\s*=\s*["']?([^"']+)["']?`
	//reImage = `<img[\s\S]+?src="(http[\s\S]+?)"[\s\S]*?>`
	// imgDir = "D:\\www\\Go\\go_demo\\concurrence_2\\day5\\spider-01\\images\\"
	imgDir     = "/Users/hehongbo/www/GO/go_demo_code/concurrence_2/day5/spider-01/images2/"
	randomMT   sync.Mutex
	downloadWG sync.WaitGroup
	chSem      = make(chan int, 5)
)

func main011() {
	start := time.Now()
	imginfos := GetPageImagesInfos(url)
	//GetPageImagesInfos3(url)
	for _, imginfoMap := range imginfos {
		//DownloadImg(imginfoMap["url"], imginfoMap["filename"])
		DownloadImgAsync(imginfoMap["url"], imginfoMap["filename"])
	}
	//fmt.Println(imginfos)
	//end := time.Now()
	//consume := end.Sub(start).Seconds()
	consume := time.Now().Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)

}

func GetPageImagesInfos(url string) []map[string]string {
	html := GetUrlHtml(url)
	//html = string(Common.ConvertToByte(html, "gbk", "utf8"))
	re := regexp.MustCompile(reImage)
	rets := re.FindAllStringSubmatch(html, -1)
	fmt.Println("捕获图片张数：", len(rets))
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

func DownloadImgAsync(url, filename string) {
	downloadWG.Add(1)
	go func() {
		chSem <- 123
		DownloadImg(url, filename)
		<-chSem
		downloadWG.Done()
	}()
	downloadWG.Wait()
}

func DownloadImg(url string, filename string) {
	fmt.Println("DownloadImg...")

	resp, err := http.Get(url)
	defer resp.Body.Close()

	imgBytes, _ := ioutil.ReadAll(resp.Body)
	err = ioutil.WriteFile(filename, imgBytes, 0644)
	if err == nil {
		fmt.Println(filename, "下载成功！")
	} else {
		fmt.Println("下载文件失败", err)
		fmt.Println(filename, "下载失败！")
	}
}

/*
*
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

/*生成[start,end)之间的随机数*/
func GetRandomInt(start, end int) int {
	randomMT.Lock()
	<-time.After(1 * time.Nanosecond)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ret := start + r.Intn(end-start)
	randomMT.Unlock()
	return ret
}

/*生成时间戳_随机数文件名*/
func GetRandomName() string {
	timestamp := strconv.Itoa(int(time.Now().UnixNano()))
	randomNum := strconv.Itoa(GetRandomInt(1000, 10000))
	return timestamp + "_" + randomNum
}

/*从imgUrl中摘取图片名称*/
func GetImgNameFromImgUrl(imgUrl string) string {
	re := regexp.MustCompile(Common.ReImgName)
	rets := re.FindAllStringSubmatch(imgUrl, -1)
	fmt.Println(rets)
	if len(rets) > 0 {
		return rets[0][1]
	} else {
		return ""
	}
}

func GetImgNameFromImageUrl(imgUrl string) string {
	_, filename := filepath.Split(imgUrl)
	if filename != "" {
		return filename
	} else {
		return ""
	}
}

func HandleError(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}

}

func GetUrlHtml(url string) string {
	resp, err := http.Get(url)
	HandleError(err, "Http.Get")
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	html := string(bytes)
	return html
}

var httpClient http.Client

func init() {
	httpClient = http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: 5 * time.Second, // 连接超时时间
			}).DialContext,
			TLSHandshakeTimeout: 5 * time.Second,  // TLS 握手超时时间
			IdleConnTimeout:     90 * time.Second, // 空闲连接超时时间
		},
		Timeout: 10 * time.Second, // 总超时时间（包括连接、TLS握手、请求和响应）
	}
}

func DownloadImgWithClient(url, filename string) {
	dump.Println("DownloadImgWithClient...")
	resp, err := httpClient.Get(url)
	if err != nil {
		dump.Println(filename, url, "文件链接请求失败")
		return
	}
	defer resp.Body.Close()
	imageBytes, _ := io.ReadAll(resp.Body)
	err = os.WriteFile(filename, imageBytes, 0644)
	if err == nil {
		dump.Println(filename, "下载成功")
	} else {
		dump.Println(filename, "下载失败")
	}
}
