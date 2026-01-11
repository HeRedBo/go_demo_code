package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"strings"
)

// ================ 以下是ES DSL的完整映射结构体，可全局复用 ================
// EsSearchDSL ES查询的根DSL结构体，对应完整的查询JSON
type EsSearchDSL struct {
	From   int64        `json:"from"`              // 分页起始位置 (page-1)*size
	Size   int64        `json:"size"`              // 每页条数
	Source []string     `json:"_source,omitempty"` // 返回的字段，可选
	Query  EsQuery      `json:"query"`             // 查询条件核心
	Sort   []EsSortItem `json:"sort,omitempty"`    // 排序规则，可选，动态追加
}

// EsQuery 查询条件容器，目前只实现最常用的bool查询，可扩展match_all等
type EsQuery struct {
	Bool EsBoolQuery `json:"bool"`
}

// EsBoolQuery 布尔查询核心结构体【重中之重】
// must/filter/should/must_not 都是切片，支持动态append追加条件，完美适配动态构造！
type EsBoolQuery struct {
	Must           []map[string]interface{} `json:"must,omitempty"`
	Filter         []map[string]interface{} `json:"filter,omitempty"`
	Should         []map[string]interface{} `json:"should,omitempty"`
	MustNot        []map[string]interface{} `json:"must_not,omitempty"`
	MinShouldMatch int                      `json:"minimum_should_match,omitempty"` // should最小匹配数
}

// EsSortItem 排序项结构体，支持动态排序字段+排序方式
type EsSortItem struct {
	Field string `json:"-"` // 排序字段，比如 sales
	Order string `json:"-"` // 排序方式，asc/desc
}

// MarshalJSON 重写排序项的序列化方法，适配ES的排序语法 {"sales":{"order":"desc"}}
func (s EsSortItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]map[string]string{
		s.Field: {
			"order": s.Order,
		},
	})
}

// SearchGoods 生产级动态构造DSL查询商品【完整实现】
func SearchGoods(req GoodsSearchReq) ([]map[string]interface{}, int64, error) {
	// 1. 参数默认值处理 & 校验
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 || req.Size > 100 { // 限制最大条数，防止分页过大
		req.Size = 10
	}
	if req.SortOrder == "" {
		req.SortOrder = "desc"
	}
	from := int64((req.Page - 1) * req.Size)
	size := int64(req.Size)

	// 2. 初始化ES查询DSL的基础结构体
	dsl := EsSearchDSL{
		From:   from,
		Size:   size,
		Source: []string{"id", "title", "category", "price", "sales", "score", "is_new"}, // 指定返回字段
		Query: EsQuery{
			Bool: EsBoolQuery{
				MinShouldMatch: 0, // should默认不强制匹配
			},
		},
	}

	// 3. 【核心】动态组装查询条件：根据前端参数，有值则追加，无值则忽略
	boolQuery := &dsl.Query.Bool
	// 3.1 标题分词检索：有关键词则追加到must（参与打分）
	if strings.TrimSpace(req.Keyword) != "" {
		boolQuery.Must = append(boolQuery.Must, map[string]interface{}{
			"match": map[string]interface{}{
				"title": req.Keyword,
			},
		})
	}
	// 3.2 分类精准筛选：有分类则追加到filter（不打分，高性能）
	if len(req.Categories) > 0 {
		boolQuery.Filter = append(boolQuery.Filter, map[string]interface{}{
			"terms": map[string]interface{}{
				"category": req.Categories,
			},
		})
	}
	// 3.3 价格区间筛选：有价格则追加到filter
	if req.MinPrice > 0 || req.MaxPrice > 0 {
		rangeCond := map[string]interface{}{}
		if req.MinPrice > 0 {
			rangeCond["gte"] = req.MinPrice
		}
		if req.MaxPrice > 0 {
			rangeCond["lte"] = req.MaxPrice
		}
		boolQuery.Filter = append(boolQuery.Filter, map[string]interface{}{
			"range": map[string]interface{}{
				"price": rangeCond,
			},
		})
	}
	// 3.4 销量筛选：最低销量>0则追加到filter
	if req.MinSales > 0 {
		boolQuery.Filter = append(boolQuery.Filter, map[string]interface{}{
			"range": map[string]interface{}{
				"sales": map[string]interface{}{
					"gte": req.MinSales,
				},
			},
		})
	}
	// 3.5 是否新品筛选：指针非nil表示前端传了该参数
	if req.IsNew != nil {
		boolQuery.Filter = append(boolQuery.Filter, map[string]interface{}{
			"term": map[string]interface{}{
				"is_new": *req.IsNew,
			},
		})
	}
	// 3.6 必加条件：过滤库存>0的商品（固定条件，所有查询都生效）
	boolQuery.Filter = append(boolQuery.Filter, map[string]interface{}{
		"range": map[string]interface{}{
			"stock": map[string]interface{}{
				"gt": 0,
			},
		},
	})

	// 4. 【核心】动态组装排序规则：前端传了排序字段则追加
	if strings.TrimSpace(req.SortField) != "" {
		// 生产建议：做排序字段白名单校验，防止前端传非法字段导致查询失败
		allowSortFields := map[string]bool{"sales": true, "price": true, "score": true, "create_time": true}
		if allowSortFields[req.SortField] {
			dsl.Sort = append(dsl.Sort, EsSortItem{
				Field: req.SortField,
				Order: req.SortOrder,
			})
		}
	}

	// 5. 将Go结构体序列化为标准的ES DSL JSON字符串【无任何语法错误】
	dslJson, err := json.MarshalIndent(dsl, "", "  ") // 带缩进，方便日志查看，生产可改为json.Marshal
	if err != nil {
		return nil, 0, fmt.Errorf("DSL序列化失败: %w", err)
	}
	log.Println("✅ 动态构造的ES DSL:\n", string(dslJson))

	// 6. 执行ES查询
	searchReq := esapi.SearchRequest{
		Index: []string{"goods"},
		Body:  strings.NewReader(string(dslJson)),
	}
	res, err := searchReq.Do(context.Background(), esClient)
	if err != nil {
		return nil, 0, fmt.Errorf("ES查询失败: %w", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		return nil, 0, fmt.Errorf("ES查询响应异常: %s", res.Status())
	}

	// 7. 解析查询结果
	var result map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, 0, fmt.Errorf("解析ES结果失败: %w", err)
	}
	// 总条数
	total := int64(result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))
	// 数据列表
	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	var dataList []map[string]interface{}
	for _, hit := range hits {
		dataList = append(dataList, hit.(map[string]interface{})["_source"].(map[string]interface{}))
	}

	return dataList, total, nil
}
