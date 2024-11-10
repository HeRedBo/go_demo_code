package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"regexp"
)

func main01() {
	buf := "abc azc a7c aac 888 a9c  tac"
	//解析正则表达式，如果成功返回解释器
	reg1 := regexp.MustCompile(`a.c`)

	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}
	//根据规则提取关键信息
	result1 := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println("result1 = ", result1)
}

func main02() {
	buf := "abc azc a7c aac 888 a9c  tac"
	//解析正则表达式，如果成功返回解释器
	reg1 := regexp.MustCompile(`a[0-9]c`)

	fmt.Println("reg1", reg1)
	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}
	//根据规则提取关键信息
	result1 := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println("result1 = ", result1)
}

func main03() {
	img_src := `<img src="https://pic2.zhimg.com/80/v2-e0e2deceb65b93e070776fb79323f20d_720w.jpg" data-rawwidth="1080" data-rawheight="1920" data-size="normal" class="origin_image zh-lightbox-thumb lazy" width="1080" data-original="https://pic2.zhimg.com/v2-e0e2deceb65b93e070776fb79323f20d_r.jpg" data-actualsrc="https://pic2.zhimg.com/v2-e0e2deceb65b93e070776fb79323f20d_b.jpg" data-lazy-status="ok">`

	//reImgTagStr := `src="(https.+?)`
	//var reTagSrcStr = `/(\w+\.((jpg)|(jpeg)|(png)|(gif)|(bmp)|(webp)|(swf)|(ico)))`
	var reTagSrcStr = `src="(http.+?)"`

	reg1 := regexp.MustCompile(reTagSrcStr)
	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}

	result1 := reg1.FindAllStringSubmatch(img_src, -1)
	fmt.Println(reflect.TypeOf(result1).String())
	fmt.Println("result1 = ", result1)
}

var (
	rePhone = `(1[3456789]\d)(\d{4})(\d{4})`
	/*超链接*/
	/*
		<a...href="https://www.hao123.com/link/https/?key=http%3A%2F%2Fwww.163.com%2F&amp;&amp;monkey=m-mingzhan-site&amp;c=CCBD97A87DB464FDCD39F475585D5AE6"...>
	*/
	reLink = `<a[\s\S]+?href="(http[\s\S]+?)"`
)

func spiderPhone() {
	html := GetHtml("https://www.haomagujia.com/")
	//fmt.Println(html)

	re := regexp.MustCompile(rePhone)
	allString := re.FindAllStringSubmatch(html, -1)
	for _, x := range allString {
		fmt.Println(x[0])
	}
}

func spiderLink() {
	html := GetHtml(`http://www.hao123.com`)
	//fmt.Println(html)
	re := regexp.MustCompile(reLink)
	rets := re.FindAllStringSubmatch(html, -1)
	for _, x := range rets {
		fmt.Println(x[1])
	}
}

func main010() {
	spiderPhone()
	//spiderLink()

}

func GetHtml(url string) string {
	resp, err := http.Get(url)
	HandleErr(err, "Http.Get")
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	html := string(bytes)

	return html
}

func HandleErr(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}

}
