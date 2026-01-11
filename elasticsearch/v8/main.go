package main

import (
	"crypto/tls"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"net/http"
)

// ESConfig ES连接配置
type ESConfig struct {
	Addresses  []string // ES地址列表，如: ["http://127.0.0.1:9200"]
	Username   string   // 账号，ES8默认elastic
	Password   string   // 密码，ES8安装时生成的密码
	EnableTLS  bool     // 是否开启TLS加密
	SkipVerify bool     // 生产环境：true=跳过证书验证(临时)，false=配置证书(正式)
}

// 全局ES客户端单例
var esClient *elasticsearch.Client

// InitESClient 初始化ES客户端
func InitESClient(cfg ESConfig) error {
	var err error
	esCfg := elasticsearch.Config{
		Addresses: cfg.Addresses,
		Username:  cfg.Username,
		Password:  cfg.Password,
	}

	// 生产环境：开启TLS配置
	if cfg.EnableTLS {
		esCfg.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: cfg.SkipVerify},
		}
	}

	// 创建客户端
	esClient, err = elasticsearch.NewClient(esCfg)
	if err != nil {
		return fmt.Errorf("es客户端初始化失败: %w", err)
	}

	// 测试连通性
	res, err := esClient.Info()
	if err != nil {
		return fmt.Errorf("es连通性测试失败: %w", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("es服务异常: %s", res.Status())
	}
	log.Println("✅ ES客户端初始化成功，版本:", res.Header.Get("X-Elastic-Product"))
	return nil
}

// 程序入口初始化调用示例
func main() {
	// ============ 方案1：本地开发（无账号密码、关闭TLS，推荐） ============
	localCfg := ESConfig{
		Addresses:  []string{"http://127.0.0.1:9200"},
		Username:   "elastic",
		Password:   "elastic",
		EnableTLS:  false,
		SkipVerify: false,
	}

	// ============ 方案2：生产环境（带账号密码、开启TLS，必须） ============
	// prodCfg := ESConfig{
	// 	Addresses:  []string{"https://192.168.1.100:9200", "https://192.168.1.101:9200"}, // 集群地址
	// 	Username:   "elastic",
	// 	Password:   "你的ES8密码",
	// 	EnableTLS:  true,
	// 	SkipVerify: true, // 生产临时跳过证书，正式环境配置证书路径即可
	// }

	// 初始化客户端
	err := InitESClient(localCfg)
	if err != nil {
		log.Fatal(err)
	}

	// 执行后续操作：创建索引、初始化数据、查询排序等
	//err = CreateGoodsIndex()
	//if err != nil {
	//	dump.P(err)
	//	log.Fatal(err)
	//}

	//_ = InitGoodsData()

	//  调用示例：模拟前端传参，测试动态构造 DSL

	// 模拟前端传递的动态查询参数：搜索手机 + 分类手机 + 价格2000-8000 + 销量≥5000 + 新品 + 按销量降序
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
		SortField:  "sales",
		SortOrder:  "desc",
	}

	// 执行动态查询
	dataList, total, err := SearchGoods(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("✅ 查询结果总条数:", total)
	log.Println("✅ 查询结果列表:", dataList)

}
