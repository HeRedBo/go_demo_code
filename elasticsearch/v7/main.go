package main

import "log"

func main01() {

}

func main() {
	InitEsClient()
	// 模拟前端传参：搜索手机 + 分类手机 + 价格2000-8000 + 销量≥5000 + 新品
	isNew := true
	req := GoodsSearchReq{
		Keyword:    "手机",
		Categories: []string{"手机"},
		MinPrice:   2000,
		MaxPrice:   8000,
		MinSales:   5000,
		IsNew:      &isNew,
		Page:       1,
		Size:       10,
	}
	list, total, err := BoolQueryDemo(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("查询结果：", list, "总条数：", total)
}
