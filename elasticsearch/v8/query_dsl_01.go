package main

// 全局ES客户端（复用之前的单例，无需改动）
//var esClient *elasticsearch.Client

// GoodsSearchReq 前端传递的动态查询参数【核心】
// 所有字段都是可选，前端不传则为空，后端动态组装条件
type GoodsSearchReq struct {
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
