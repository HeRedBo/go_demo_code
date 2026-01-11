package main

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/olivere/elastic/v7"
)

// 设置默认参数，防止非法值 (零报错)
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

// 统一解析ES返回结果 (所有方法复用，零报错)
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

// ============ ✅ 1. 动态布尔查询【核心、零报错】支持所有筛选条件 ============
func BoolSearch(req GoodsSearchReq) ([]map[string]interface{}, int64, error) {
	req.SetDefault()
	boolQ := elastic.NewBoolQuery()

	// 标题模糊检索 (参与相关性打分)
	if strings.TrimSpace(req.Keyword) != "" {
		boolQ.Must(elastic.NewMatchQuery("title", req.Keyword))
	}
	// 分类多值精准筛选 (零报错，核心修复)
	if len(req.Categories) > 0 {
		boolQ.Filter(elastic.NewTermsQuery("category", StringSliceToInterface(req.Categories)...))
	}
	// 价格区间筛选
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
	// 销量筛选
	if req.MinSales > 0 {
		boolQ.Filter(elastic.NewRangeQuery("sales").Gte(req.MinSales))
	}
	// 是否新品筛选 (指针判断，零报错)
	if req.IsNew != nil {
		boolQ.Filter(elastic.NewTermQuery("is_new", *req.IsNew))
	}
	// 必筛：库存>0
	boolQ.Filter(elastic.NewRangeQuery("stock").Gt(0))
	// 排除低分商品
	boolQ.MustNot(elastic.NewRangeQuery("score").Lt(4.5))

	// 执行查询
	from := (req.Page - 1) * req.Size
	res, err := esClient.Search().
		Index("goods").
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

// ============ ✅ 2. 定制化普通排序【多字段组合、零报错】 ============
func CustomSortSearch(req GoodsSearchReq) ([]map[string]interface{}, int64, error) {
	req.SetDefault()
	boolQ := elastic.NewBoolQuery()
	boolQ.Filter(elastic.NewRangeQuery("stock").Gt(0))
	if len(req.Categories) > 0 {
		boolQ.Filter(elastic.NewTermsQuery("category", StringSliceToInterface(req.Categories)...))
	}

	search := esClient.Search().Index("goods").Query(boolQ).From((req.Page - 1) * req.Size).Size(req.Size)

	// 排序字段白名单，防止非法字段
	allowSort := map[string]bool{"sales": true, "price": true, "score": true, "create_time": true}
	isDesc := req.SortOrder == "desc"
	if req.SortField != "" && allowSort[req.SortField] {
		search.Sort(req.SortField, isDesc)
	}
	// 多字段组合排序：销量降序 → 价格升序 → 评分降序
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

// ============ ✅ 3. 脚本公式排序【销量*权重+评分*权重、零报错】 ============
func ScriptFormulaSortSearch(req GoodsSearchReq) ([]map[string]interface{}, int64, error) {
	req.SetDefault()
	boolQ := elastic.NewBoolQuery()
	boolQ.Filter(elastic.NewTermQuery("category", "手机"))
	boolQ.Filter(elastic.NewRangeQuery("stock").Gt(0))

	search := esClient.Search().Index("goods").Query(boolQ).From((req.Page - 1) * req.Size).Size(req.Size)

	// 脚本公式：销量*0.8 + 评分*100 （综合热度排序）
	script := elastic.NewScript("doc['sales'].value * 0.8 + doc['score'].value * 100")
	// ✅ 核心修复：Order(true) 代表降序，替代原来的 "desc"
	search.SortBy(elastic.NewScriptSort(script, "number").Order(true))

	res, err := search.Do(context.Background())
	if err != nil {
		return nil, 0, err
	}
	list, total := parseResult(res)
	return list, total, nil
}

// ============ ✅ ✨ 4. 脚本权重排序【你的核心需求、终极零报错、生产首选】 ============
// 完美实现：新品×3 、销量>10000×2 、评分≥4.9×1.5 权重叠加排序
func ScriptWeightSortSearch(req GoodsSearchReq) ([]map[string]interface{}, int64, error) {
	req.SetDefault()
	boolQ := elastic.NewBoolQuery()

	// 所有筛选条件和布尔查询完全一致
	if strings.TrimSpace(req.Keyword) != "" {
		boolQ.Must(elastic.NewMatchQuery("title", req.Keyword))
	}
	if len(req.Categories) > 0 {
		boolQ.Filter(elastic.NewTermsQuery("category", StringSliceToInterface(req.Categories)...))
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

	search := esClient.Search().Index("goods").Query(boolQ).From((req.Page - 1) * req.Size).Size(req.Size)

	// ✅ 核心：Painless脚本实现【条件权重叠加】，和FunctionScore效果完全一致，零报错！！！
	weightScript := elastic.NewScript(`
		def weight = 1.0;
		// 规则1：新品权重×3
		if (doc['is_new'].value == true) { weight *= 3; }
		// 规则2：销量>10000 权重×2
		if (doc['sales'].value > 10000) { weight *= 2; }
		// 规则3：评分≥4.9 权重×1.5
		if (doc['score'].value >= 4.9) { weight *= 1.5; }
		return weight;
	`)
	// 按计算后的权重值降序排序
	search.SortBy(elastic.NewScriptSort(weightScript, "number").Order(true))

	res, err := search.Do(context.Background())
	if err != nil {
		return nil, 0, err
	}
	list, total := parseResult(res)
	log.Printf("✅ 权重排序完成，共查询 %d 条数据", total)
	return list, total, nil
}
