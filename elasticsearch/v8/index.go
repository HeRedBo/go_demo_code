package main

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"net/http"
	"strings"
	"time"
)

// CreateGoodsIndex 创建商品索引 goods
func CreateGoodsIndex() error {
	// 判断索引是否已存在
	exists, err := IsIndexExists("demo_goods")
	if err != nil {
		return err
	}
	if exists {
		log.Println("索引 goods 已存在，无需重复创建")
		return nil
	}

	// 构造索引创建请求：包含Mapping映射+分片副本配置
	req := esapi.IndicesCreateRequest{
		Index: "demo_goods",
		Body: strings.NewReader(`{
			"settings": {
				"number_of_shards": 1,
				"number_of_replicas": 0,
				"refresh_interval": "1s"
			},
			"mappings": {
				"properties": {
					"id": {"type": "long"},
					"title": {"type": "text", "analyzer": "ik_max_word"}, // 中文分词，需安装ik插件
					"category": {"type": "keyword"}, // 精准匹配+排序
					"price": {"type": "float"},
					"sales": {"type": "long"}, // 销量
					"stock": {"type": "long"}, // 库存
					"score": {"type": "float"}, // 商品评分
					"is_new": {"type": "boolean"}, // 是否新品
					"create_time": {"type": "date", "format": "yyyy-MM-dd HH:mm:ss"}
				}
			}
		}`),
		Timeout: 10 * time.Second,
	}

	// 执行请求
	res, err := req.Do(context.Background(), esClient)
	if err != nil {
		return fmt.Errorf("创建索引失败: %w", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("创建索引响应异常: %s", res.Status())
	}
	log.Println("✅ 索引 goods 创建成功")
	return nil
}

// IsIndexExists 判断索引是否存在
func IsIndexExists(index string) (bool, error) {
	req := esapi.IndicesExistsRequest{Index: []string{index}}
	res, err := req.Do(context.Background(), esClient)
	if err != nil {
		return false, fmt.Errorf("判断索引失败: %w", err)
	}
	defer res.Body.Close()
	// 200=存在，404=不存在，其他=异常
	return res.StatusCode == http.StatusOK, nil
}
