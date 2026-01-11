package main

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/olivere/elastic/v7"
)

// 全局单例客户端，初始化一次即可
var esClient *elastic.Client

// InitEsClient 初始化olivere/elastic/v7客户端
func InitEsClient() {
	var err error
	// ============ 方案1：本地开发（无账号密码，ES8关闭TLS，推荐） ============
	esClient, err = elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false), // 本地环境关闭节点嗅探，生产集群开启
		elastic.SetHealthcheck(false),
	)

	// ============ 方案2：生产环境（ES8带账号密码+TLS加密，直接用） ============
	// esClient, err = elastic.NewClient(
	// 	elastic.SetURL("https://192.168.1.100:9200"),
	// 	elastic.SetBasicAuth("elastic", "你的ES8密码"),
	// 	elastic.SetSniff(true),
	// 	elastic.SetSSL(true),
	// 	elastic.SetInsecure(true), // 跳过证书验证，生产可配置证书
	// )

	if err != nil {
		log.Fatalf("初始化ES客户端失败: %v", err)
	}
	log.Println("✅ olivere/elastic/v7 客户端初始化成功")
}

// GoodsSearchReq 前端动态搜索入参【生产标准】
// 所有字段为可选，无传参则不参与查询/排序，完美适配动态构造
type GoodsSearchReq struct {
	Keyword    string   `json:"keyword"`    // 标题搜索词
	Categories []string `json:"categories"` // 分类筛选 ["手机","外设"]
	MinPrice   float64  `json:"minPrice"`   // 最低价格
	MaxPrice   float64  `json:"maxPrice"`   // 最高价格
	MinSales   int64    `json:"minSales"`   // 最低销量
	IsNew      *bool    `json:"isNew"`      // 是否新品 ✨布尔值用指针，区分false和未传参
	Page       int      `json:"page"`       // 页码，默认1
	Size       int      `json:"size"`       // 每页条数，默认10
	SortField  string   `json:"sortField"`  // 排序字段
	SortOrder  string   `json:"sortOrder"`  // asc/desc，默认desc
}

// 统一的结果解析函数，所有示例复用
func parseEsResult(res *elastic.SearchResult) ([]map[string]interface{}, int64) {
	var list []map[string]interface{}
	total := res.TotalHits()
	for _, hit := range res.Hits.Hits {
		var item map[string]interface{}
		_ = json.Unmarshal(hit.Source, &item)
		list = append(list, item)
	}
	return list, total
}

// setDefault 补全默认参数，防止非法值
func (req *GoodsSearchReq) setDefault() {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 || req.Size > 100 { // 限制最大条数，防止分页溢出
		req.Size = 10
	}
	if req.SortOrder == "" || (req.SortOrder != "asc" && req.SortOrder != "desc") {
		req.SortOrder = "desc"
	}
}

// BoolQueryDemo 动态构造布尔查询-完整示例
// 需求：标题含关键词+分类筛选+价格区间+销量筛选+是否新品+过滤无库存+排除低分商品
func BoolQueryDemo(req GoodsSearchReq) ([]map[string]interface{}, int64, error) {
	req.setDefault()
	// 1. 创建布尔查询核心对象
	boolQuery := elastic.NewBoolQuery()

	// 2. 动态追加查询条件 - 链式调用，按需追加，无值则不追加
	// 标题分词检索：有关键词则加入Must(参与打分)
	if strings.TrimSpace(req.Keyword) != "" {
		boolQuery.Must(elastic.NewMatchQuery("title", req.Keyword))
	}
	// 分类精准筛选：有分类则加入Filter(不打分，高性能)
	if len(req.Categories) > 0 {
		boolQuery.Filter(elastic.NewTermsQuery("category", StringSliceToInterface(req.Categories)...))
	}
	// 价格区间筛选：加入Filter
	if req.MinPrice > 0 || req.MaxPrice > 0 {
		priceQuery := elastic.NewRangeQuery("price")
		if req.MinPrice > 0 {
			priceQuery.Gte(req.MinPrice)
		}
		if req.MaxPrice > 0 {
			priceQuery.Lte(req.MaxPrice)
		}
		boolQuery.Filter(priceQuery)
	}
	// 销量筛选：加入Filter
	if req.MinSales > 0 {
		boolQuery.Filter(elastic.NewRangeQuery("sales").Gte(req.MinSales))
	}
	// 是否新品：指针非nil表示前端传参，加入Filter
	if req.IsNew != nil {
		boolQuery.Filter(elastic.NewTermQuery("is_new", *req.IsNew))
	}
	// 固定过滤：库存>0的商品（所有查询都生效）
	boolQuery.Filter(elastic.NewRangeQuery("stock").Gt(0))
	// 排除低分商品：评分<4.5的商品不展示，加入MustNot
	boolQuery.MustNot(elastic.NewRangeQuery("score").Lt(4.5))

	// 3. 执行查询：链式配置索引、分页、查询条件
	from := (req.Page - 1) * req.Size
	res, err := esClient.Search().
		Index("goods").   // 指定索引
		Query(boolQuery). // 绑定布尔查询
		From(from).       // 分页起始
		Size(req.Size).   // 每页条数
		Do(context.Background())

	if err != nil {
		return nil, 0, err
	}
	// 解析结果
	list, total := parseEsResult(res)
	log.Printf("布尔查询完成，共查询到 %d 条数据", total)
	return list, total, nil
}
