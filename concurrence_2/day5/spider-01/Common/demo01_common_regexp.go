package Common

var (
	ReImage = `<img.+?src="(http.+?)".*?>`

	/*
		图片链接中的图片名称
		http://cms-bucket.nosdn.127.net/2018/11/09/7e88b8526ff141129809d8ae7c718e51.jpeg?imageView&thumbnail=185y116&quality=85
		http://img2.money.126.net/chart/hs/time/180x120/0000001.png
		http://cms-bucket.nosdn.127.net/2018/05/31/bc7d30ff42194c35a4743834a77ec97b.png?imageView&thumbnail=90y90&quality=85
	*/
	ReImgName = `/(\w+\.((jpg)|(jpeg)|(png)|(gif)|(bmp)|(webp)|(swf)|(ico)))`

	/*img标签中的alt属性*/
	ReAlt = `alt="([\s\S]+?)"`
)

//
//func main() {
//
//}
