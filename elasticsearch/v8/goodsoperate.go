package main

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gookit/goutil/dump"
	"log"
	"strings"
)

// InitGoodsData 初始化商品测试数据（批量写入10条，覆盖所有查询/排序场景）
func InitGoodsData() error {
	// 构造批量写入的body，_bulk格式：每行指令+每行数据，换行分隔
	var bulkBody strings.Builder
	goodsList := []map[string]interface{}{
		{"id": 1, "title": "苹果15 Pro 256G", "category": "手机", "price": 6999.0, "sales": 10000, "stock": 500, "score": 4.9, "is_new": true, "create_time": "2025-01-01 10:00:00"},
		{"id": 2, "title": "华为Mate70 512G", "category": "手机", "price": 7999.0, "sales": 8000, "stock": 300, "score": 4.8, "is_new": true, "create_time": "2025-01-02 10:00:00"},
		{"id": 3, "title": "小米14 128G", "category": "手机", "price": 3999.0, "sales": 15000, "stock": 1000, "score": 4.7, "is_new": false, "create_time": "2024-12-01 10:00:00"},
		{"id": 4, "title": "罗技G502游戏鼠标", "category": "外设", "price": 299.0, "sales": 20000, "stock": 2000, "score": 4.9, "is_new": false, "create_time": "2024-11-01 10:00:00"},
		{"id": 5, "title": "雷蛇黑寡妇机械键盘", "category": "外设", "price": 599.0, "sales": 12000, "stock": 800, "score": 4.8, "is_new": true, "create_time": "2025-01-03 10:00:00"},
		{"id": 6, "title": "苹果AirPods Pro2", "category": "耳机", "price": 1799.0, "sales": 18000, "stock": 1500, "score": 4.9, "is_new": true, "create_time": "2025-01-04 10:00:00"},
		{"id": 7, "title": "华为FreeBuds Pro3", "category": "耳机", "price": 1299.0, "sales": 13000, "stock": 900, "score": 4.7, "is_new": true, "create_time": "2025-01-05 10:00:00"},
		{"id": 8, "title": "小米平板7 Pro", "category": "平板", "price": 2499.0, "sales": 7000, "stock": 400, "score": 4.6, "is_new": true, "create_time": "2025-01-06 10:00:00"},
		{"id": 9, "title": "iPad Air6", "category": "平板", "price": 4799.0, "sales": 6000, "stock": 200, "score": 4.8, "is_new": false, "create_time": "2024-12-05 10:00:00"},
		{"id": 10, "title": "红米K70 至尊版", "category": "手机", "price": 2999.0, "sales": 25000, "stock": 1200, "score": 4.7, "is_new": false, "create_time": "2024-12-10 10:00:00"},
	}

	// 拼接_bulk格式
	for _, goods := range goodsList {
		// 指令行：index表示新增，_index指定索引，_id指定文档ID（建议用业务ID）
		bulkBody.WriteString(fmt.Sprintf(`{"index": {"_index": "demo_goods", "_id": %d}}`+"\n", goods["id"]))
		// 数据行
		bulkBody.WriteString(fmt.Sprintf(`{"id":%d,"title":"%s","category":"%s","price":%f,"sales":%d,"stock":%d,"score":%f,"is_new":%t,"create_time":"%s"}`+"\n",
			goods["id"], goods["title"], goods["category"], goods["price"], goods["sales"], goods["stock"], goods["score"], goods["is_new"], goods["create_time"]))
	}

	dump.P(bulkBody)
	dump.P(bulkBody.String())
	// 执行批量写入
	req := esapi.BulkRequest{
		Body:    strings.NewReader(bulkBody.String()),
		Refresh: "true", // 写入后立即刷新，测试用；生产环境建议去掉，提升性能
	}
	res, err := req.Do(context.Background(), esClient)
	if err != nil {
		return fmt.Errorf("批量初始化数据失败: %w", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("批量初始化数据响应异常: %s", res.Status())
	}
	log.Println("✅ 10条商品测试数据初始化完成")
	return nil
}

// 单条新增
func AddGoodsData(id int64, data map[string]interface{}) error {
	req := esapi.IndexRequest{
		Index:      "demo_goods",
		DocumentID: fmt.Sprintf("%d", id),
		Body:       strings.NewReader(fmt.Sprintf(`%v`, data)),
		Refresh:    "true",
	}
	res, err := req.Do(context.Background(), esClient)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}

// 单条更新（指定字段）
func UpdateGoodsData(id int64, updateData map[string]interface{}) error {
	req := esapi.UpdateRequest{
		Index:      "demo_goods",
		DocumentID: fmt.Sprintf("%d", id),
		Body:       strings.NewReader(fmt.Sprintf(`{"doc":%v}`, updateData)),
		Refresh:    "true",
	}
	res, err := req.Do(context.Background(), esClient)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}

// 单条删除
func DeleteGoodsData(id int64) error {
	req := esapi.DeleteRequest{
		Index:      "demo_goods",
		DocumentID: fmt.Sprintf("%d", id),
	}
	res, err := req.Do(context.Background(), esClient)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
