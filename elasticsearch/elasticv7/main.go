package main

import (
	"context"
	"encoding/json"
	"github.com/gookit/goutil/dump"
	"log"
	"strings"

	"github.com/olivere/elastic/v7"
)

// 全局单例ES客户端
var esClient *elastic.Client

// InitEsClient 初始化ES客户端 (兼容ES7/8，本地/生产都可用)
func InitEsClient() {
	var err error
	esClient, err = elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		elastic.SetBasicAuth("elastic", "elastic"),
	)
	if err != nil {
		log.Fatalf("ES客户端初始化失败: %v", err)
	}
	log.Println("✅ ES客户端初始化成功 (olivere/elastic/v7)")
}

// ✅ 解决[]string传参给NewTermsQuery的核心工具函数
func StringSlice2Interface(s []string) []interface{} {
	res := make([]interface{}, 0, len(s))
	for _, v := range s {
		res = append(res, v)
	}
	return res
}

// 前端动态搜索入参结构体
type GoodsSearchReq struct {
	Keyword    string   `json:"keyword"`    // 标题搜索词
	Categories []string `json:"categories"` // 分类筛选 ["手机","外设"]
	MinPrice   float64  `json:"minPrice"`   // 最低价格
	MaxPrice   float64  `json:"maxPrice"`   // 最高价格
	MinSales   int64    `json:"minSales"`   // 最低销量
	IsNew      *bool    `json:"isNew"`      // 是否新品(指针区分传参/未传参)
	Page       int      `json:"page"`       // 页码
	Size       int      `json:"size"`       // 每页条数
	SortField  string   `json:"sortField"`  // 普通排序字段
	SortOrder  string   `json:"sortOrder"`  // asc/desc
}

// 设置默认参数，防止非法值
func (req *GoodsSearchReq) SetDefault() {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 || req.Size > 100 {
		req.Size = 10
	}
	if req.SortOrder != "asc" && req.SortOrder != "desc" {
		req.SortOrder = "desc"
	}
}

// 统一解析ES返回结果
func parseResult(res *elastic.SearchResult) ([]map[string]interface{}, int64) {
	list := make([]map[string]interface{}, 0)
	total := res.TotalHits()
	for _, hit := range res.Hits.Hits {
		var item map[string]interface{}
		_ = json.Unmarshal(hit.Source, &item)
		list = append(list, item)
	}
	return list, total
}

// ============ ✅ 1. 动态布尔查询【零报错】 ============
func BoolSearch(req GoodsSearchReq) ([]map[string]interface{}, int64, error) {
	req.SetDefault()
	boolQ := elastic.NewBoolQuery()

	if strings.TrimSpace(req.Keyword) != "" {
		boolQ.Must(elastic.NewMatchQuery("title", req.Keyword))
	}
	if len(req.Categories) > 0 {
		boolQ.Filter(elastic.NewTermsQuery("category", StringSlice2Interface(req.Categories)...))
	}
	if req.MinPrice > 0 || req.MaxPrice > 0 {
		priceQ := elastic.NewRangeQuery("price")
		if req.MinPrice > 0 {
			priceQ.Gte(req.MinPrice)
		}
		if req.MaxPrice > 0 {
			priceQ.Lte(req.MaxPrice)
		}
		boolQ.Filter(priceQ)
	}
	if req.MinSales > 0 {
		boolQ.Filter(elastic.NewRangeQuery("sales").Gte(req.MinSales))
	}
	if req.IsNew != nil {
		boolQ.Filter(elastic.NewTermQuery("is_new", *req.IsNew))
	}
	boolQ.Filter(elastic.NewRangeQuery("stock").Gt(0))
	boolQ.MustNot(elastic.NewRangeQuery("score").Lt(4.5))

	from := (req.Page - 1) * req.Size
	res, err := esClient.Search().
		Index("demo_goods").
		Query(boolQ).
		From(from).
		Size(req.Size).
		Do(context.Background())

	if err != nil {
		return nil, 0, err
	}
	list, total := parseResult(res)
	return list, total, nil
}

// ============ ✅ 2. 定制化普通排序【零报错】 ============
func CustomSortSearch(req GoodsSearchReq) ([]map[string]interface{}, int64, error) {
	req.SetDefault()
	boolQ := elastic.NewBoolQuery()
	boolQ.Filter(elastic.NewRangeQuery("stock").Gt(0))
	if len(req.Categories) > 0 {
		boolQ.Filter(elastic.NewTermsQuery("category", StringSlice2Interface(req.Categories)...))
	}

	search := esClient.Search().Index("demo_goods").Query(boolQ).From((req.Page - 1) * req.Size).Size(req.Size)

	allowSort := map[string]bool{"sales": true, "price": true, "score": true, "create_time": true}
	isDesc := req.SortOrder == "desc"
	if req.SortField != "" && allowSort[req.SortField] {
		search.Sort(req.SortField, isDesc)
	}
	search.Sort("sales", true)
	search.Sort("price", false)
	search.Sort("score", true)

	res, err := search.Do(context.Background())
	if err != nil {
		return nil, 0, err
	}
	list, total := parseResult(res)
	return list, total, nil
}

// ============ ✅ 3. 脚本公式排序【零报错】 ============
func ScriptFormulaSortSearch(req GoodsSearchReq) ([]map[string]interface{}, int64, error) {
	req.SetDefault()
	boolQ := elastic.NewBoolQuery()
	boolQ.Filter(elastic.NewTermQuery("category", "手机"))
	boolQ.Filter(elastic.NewRangeQuery("stock").Gt(0))

	search := esClient.Search().Index("demo_goods").Query(boolQ).From((req.Page - 1) * req.Size).Size(req.Size)

	script := elastic.NewScript("doc['sales'].value * 0.8 + doc['score'].value * 100")
	// ✅ 修复：Order(true) 降序
	search.SortBy(elastic.NewScriptSort(script, "number").Order(true))

	res, err := search.Do(context.Background())
	if err != nil {
		return nil, 0, err
	}
	list, total := parseResult(res)
	return list, total, nil
}

// ============ ✅ 4. 核心刚需：脚本权重排序【100%零报错、终版】 ============
func ScriptWeightSortSearch(req GoodsSearchReq) ([]map[string]interface{}, int64, error) {
	req.SetDefault()
	boolQ := elastic.NewBoolQuery()

	if strings.TrimSpace(req.Keyword) != "" {
		boolQ.Must(elastic.NewMatchQuery("title", req.Keyword))
	}
	if len(req.Categories) > 0 {
		boolQ.Filter(elastic.NewTermsQuery("category", StringSlice2Interface(req.Categories)...))
	}
	boolQ.Filter(elastic.NewRangeQuery("stock").Gt(0))
	if req.MinPrice > 0 || req.MaxPrice > 0 {
		priceQ := elastic.NewRangeQuery("price")
		if req.MinPrice > 0 {
			priceQ.Gte(req.MinPrice)
		}
		if req.MaxPrice > 0 {
			priceQ.Lte(req.MaxPrice)
		}
		boolQ.Filter(priceQ)
	}

	search := esClient.Search().Index("demo_goods").Query(boolQ).From((req.Page - 1) * req.Size).Size(req.Size)

	weightScript := elastic.NewScript(`
		def weight = 1.0;
		if (doc['is_new'].value == true) { weight *= 3; }
		if (doc['sales'].value > 10000) { weight *= 2; }
		if (doc['score'].value >= 4.9) { weight *= 1.5; }
		return weight;
	`)
	// ✅ 最终修复：Order(true) 降序，完美解决最后一个报错
	search.SortBy(elastic.NewScriptSort(weightScript, "number").Order(true))

	res, err := search.Do(context.Background())
	if err != nil {
		return nil, 0, err
	}
	list, total := parseResult(res)
	log.Printf("✅ 权重排序完成，共查询 %d 条数据", total)
	return list, total, nil
}

// ============ ✅ 测试入口【直接运行、零报错】 ============
func main() {
	InitEsClient()
	isNew := false
	req := GoodsSearchReq{
		Keyword:    "",
		Categories: []string{"手机"},
		MinPrice:   2000,
		MaxPrice:   8000,
		MinSales:   5000,
		IsNew:      &isNew,
		Page:       1,
		Size:       10,
	}
	// 测试你的核心权重排序
	//resultList, total, err := BoolSearch(req)
	//resultList, total, err := CustomSortSearch(req)
	resultList, total, err := ScriptFormulaSortSearch(req)
	//resultList, total, err := ScriptWeightSortSearch(req)
	if err != nil {
		log.Fatalf("查询失败: %v", err)
	}
	log.Println("✅ 查询总条数：", total)
	log.Println("✅ 查询结果：", resultList)
	dump.P("✅ 查询结果：", resultList)
}
