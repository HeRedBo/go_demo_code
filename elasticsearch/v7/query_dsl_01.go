package main

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"strings"
)

// 先安装依赖
// go get github.com/olivere/elastic/v7

// 所有字段都是可选，前端不传则为空，后端动态组装条件
type GoodsSearchReq1 struct {
	Keyword    string   `json:"keyword"`    // 商品标题搜索词（可选）
	Categories []string `json:"categories"` // 分类筛选，比如["手机","外设"]（可选）
	MinPrice   float64  `json:"minPrice"`   // 最低价格（可选）
	MaxPrice   float64  `json:"maxPrice"`   // 最高价格（可选）
	MinSales   int64    `json:"minSales"`   // 最低销量（可选）
	IsNew      *bool    `json:"isNew"`      // 是否新品（布尔值用指针，区分false和未传参）
	Page       int      `json:"page"`       // 页码，默认1
	Size       int      `json:"size"`       // 每页条数，默认10
	SortField  string   `json:"sortField"`  // 排序字段，比如 sales/price/score（可选）
	SortOrder  string   `json:"sortOrder"`  // 排序方式，asc/desc，默认desc
}

// ✅ 核心工具函数：解决[]string传参问题
func StringSliceToInterface(s []string) []interface{} {
	interfaceSlice := make([]interface{}, 0, len(s))
	for _, v := range s {
		interfaceSlice = append(interfaceSlice, v)
	}
	return interfaceSlice
}

// 初始化elastic客户端（复用esClient的配置）
func InitElasticClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false), // 本地开发关闭嗅探，生产集群开启
		elastic.SetHealthcheck(false),
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// SearchGoodsByBuilder 基于Builder动态构造DSL查询商品【生产级】
func SearchGoodsByBuilder(req GoodsSearchReq1) ([]map[string]interface{}, int64, error) {
	client, err := InitElasticClient()
	if err != nil {
		return nil, 0, err
	}

	// 初始化bool查询
	boolQuery := elastic.NewBoolQuery()

	// 动态追加条件：链式调用，极简！
	if strings.TrimSpace(req.Keyword) != "" {
		boolQuery.Must(elastic.NewMatchQuery("title", req.Keyword))
	}
	if len(req.Categories) > 0 {
		boolQuery.Filter(elastic.NewTermsQuery("category", StringSliceToInterface(req.Categories)...))
	}
	if req.MinPrice > 0 || req.MaxPrice > 0 {
		rangeQuery := elastic.NewRangeQuery("price")
		if req.MinPrice > 0 {
			rangeQuery.Gte(req.MinPrice)
		}
		if req.MaxPrice > 0 {
			rangeQuery.Lte(req.MaxPrice)
		}
		boolQuery.Filter(rangeQuery)
	}
	if req.IsNew != nil {
		boolQuery.Filter(elastic.NewTermQuery("is_new", *req.IsNew))
	}

	// 分页+排序
	from := (req.Page - 1) * req.Size
	search := client.Search().Index("demo_goods").Query(boolQuery).From(from).Size(req.Size)
	if req.SortField != "" {
		search.Sort(req.SortField, req.SortOrder == "desc")
	}

	// 执行查询
	res, err := search.Do(context.Background())
	if err != nil {
		return nil, 0, err
	}

	// 解析结果
	var dataList []map[string]interface{}
	for _, hit := range res.Hits.Hits {
		var item map[string]interface{}
		_ = json.Unmarshal(hit.Source, &item)
		dataList = append(dataList, item)
	}
	return dataList, res.TotalHits(), nil
}
