package attrfomat

import (
	"github.com/gookit/goutil/dump"
	"testing"
)

func TestGenerateSKUs(t *testing.T) {
	// 定义商品属性
	attributes := []Attribute{
		{
			Name:   "颜色",
			Values: []string{"黑色", "白色", "红色"},
		},
		{
			Name:   "尺寸",
			Values: []string{"30", "50"},
		},
		{
			Name:   "内存大小",
			Values: []string{"100G", "200G"},
		},
	}

	// 生成所有SKU
	//skus := GenerateSkus(attributes)
	skus := GenerateSKUsOptimized(attributes)
	dump.P(skus)
	// 打印结果
	//fmt.Printf("共生成 %d 个SKU:\n", len(skus))
	//for i, sku := range skus {
	//	fmt.Printf("%d: %s\n", i+1, sku)
	//}
}
