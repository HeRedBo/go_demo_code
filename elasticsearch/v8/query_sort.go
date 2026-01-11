package main

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"strings"
)

// CustomSortDemo 定制化普通排序示例
func CustomSortDemo() ([]map[string]interface{}, error) {
	queryDSL := strings.NewReader(`{
		"from":0,
		"size":10,
		"_source":["id","title","category","price","sales","score"],
		"query": {
			"bool": {
				"filter": [{"terms": {"category": ["手机", "外设", "耳机"]}}]
			}
		},
		"sort": [
			{"sales": {"order": "desc"}},  // 第一排序：销量降序（核心）
			{"price": {"order": "asc"}},   // 第二排序：价格升序
			{"score": {"order": "desc"}},  // 第三排序：评分降序
			{"_score": {"order": "desc"}}  // 兜底：相关性得分降序
		]
	}`)

	req := esapi.SearchRequest{Index: []string{"demo_goods"}, Body: queryDSL}
	res, err := req.Do(context.Background(), esClient)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// 解析结果同上，略
	var result map[string]interface{}
	_ = json.NewDecoder(res.Body).Decode(&result)
	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	log.Println("✅ 定制化排序结果:", hits)
	return nil, nil
}

// ScriptSortDemo Script脚本排序示例（生产级，重点）
func ScriptSortDemo() ([]map[string]interface{}, error) {
	queryDSL := strings.NewReader(`{
		"from":0,
		"size":10,
		"_source":["id","title","category","price","sales","score"],
		"query": {
			"bool": {
				"filter": [{"term": {"category": "手机"}}]
			}
		},
		"sort": [
			{
				"_script": {
					"type": "number",
					"script": {
						"lang": "painless",
						"source": "doc['sales'].value * 0.8 + doc['score'].value * 100" // 核心：自定义公式
					},
					"order": "desc" // 公式计算值降序
				}
			}
		]
	}`)

	req := esapi.SearchRequest{Index: []string{"demo_goods"}, Body: queryDSL}
	res, err := req.Do(context.Background(), esClient)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result map[string]interface{}
	_ = json.NewDecoder(res.Body).Decode(&result)
	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	log.Println("✅ Script脚本排序结果:", hits)
	return nil, nil
}

// FunctionScoreSortDemo Function Score函数得分排序（终极排序，生产级重点）
func FunctionScoreSortDemo() ([]map[string]interface{}, error) {
	queryDSL := strings.NewReader(`{
		"from":0,
		"size":10,
		"_source":["id","title","category","price","sales","score","is_new"],
		"query": {
			"function_score": {
				"query": { // 基础查询条件：布尔查询，查询所有非零库存商品
					"bool": {
						"filter": [{"range": {"stock": {"gt": 0}}}]
					}
				},
				"functions": [ // 多个得分函数，按顺序加权
					{
						"filter": {"term": {"is_new": true}}, // 条件：新品
						"weight": 3 // 加权：得分×3
					},
					{
						"filter": {"range": {"sales": {"gt": 10000}}}, // 条件：销量>10000
						"weight": 2 // 加权：得分×2
					},
					{
						"filter": {"range": {"score": {"gte": 4.9}}}, // 条件：评分≥4.9
						"weight": 1.5 // 加权：得分×1.5
					}
				],
				"boost_mode": "multiply", // 得分合并方式：原始得分 × 权重
				"score_mode": "sum", // 多函数得分组合方式：求和
				"max_boost": 10 // 最大加权值，防止得分过高
			}
		},
		"sort": [
			{"_score": {"order": "desc"}} // 按加权后的得分降序排序
		]
	}`)

	req := esapi.SearchRequest{Index: []string{"demo_goods"}, Body: queryDSL}
	res, err := req.Do(context.Background(), esClient)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result map[string]interface{}
	_ = json.NewDecoder(res.Body).Decode(&result)
	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	log.Println("✅ Function Score排序结果:", hits)
	return nil, nil
}
