package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"strings"
)

// BoolQueryDemo 布尔查询完整示例（生产级）
func BoolQueryDemo() ([]map[string]interface{}, error) {
	// 构造布尔查询DSL
	queryDSL := strings.NewReader(`{
		"from": 0,
		"size": 10,
		"_source": ["id","title","category","price","sales","score","is_new"],
		"query": {
			"bool": {
				"must": [
					{"match": {"title": "手机"}} // 标题包含手机，分词检索，参与打分
				],
				"filter": [ // 过滤条件，不打分，性能优先
					{"terms": {"category": ["手机", "外设"]}}, // 分类是手机或外设
					{"range": {"price": {"gte": 2000, "lte": 8000}}}, // 价格区间2000-8000
					{"range": {"sales": {"gte": 5000}}} // 销量≥5000
				],
				"must_not": [
					{"term": {"stock": 0}} // 排除库存为0的商品
				],
				"should": [
					{"term": {"is_new": true}} // 可选：优先新品
				],
				"minimum_should_match": 0 // should至少满足0个，即不强制
			}
		}
	}`)

	// 执行搜索请求
	req := esapi.SearchRequest{
		Index: []string{"demo_goods"},
		Body:  queryDSL,
	}
	res, err := req.Do(context.Background(), esClient)
	if err != nil {
		return nil, fmt.Errorf("布尔查询失败: %w", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		return nil, fmt.Errorf("布尔查询响应异常: %s", res.Status())
	}

	// 解析结果（简化版，生产可封装结构化解析）
	var result map[string]interface{}
	// 这里需要导入 encoding/json
	// import "encoding/json"
	_ = json.NewDecoder(res.Body).Decode(&result)
	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	var dataList []map[string]interface{}
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		dataList = append(dataList, source)
	}
	log.Println("✅ 布尔查询结果:", dataList)
	return dataList, nil
}
