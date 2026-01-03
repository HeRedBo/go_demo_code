package attrfomat

import (
	"fmt"
	"strings"
)

// Attribute 表示商品的一个属性，如颜色、尺寸等
type Attribute struct {
	Name   string   // 属性名称
	Values []string // 属性值列表
}

// SKU 表示一个具体的库存单位，包含所有属性的组合
type SKU struct {
	Attributes map[string]string // 属性名到属性值的映射
}

// GenerateSKUs 生成所有可能的SKU组合
func GenerateSKUs(attributes []Attribute) []SKU {
	if len(attributes) == 0 {
		return []SKU{}
	}
	// 初始化结果，使用第一个属性的值作为初始组合
	result := []SKU{}
	for _, val := range attributes[0].Values {
		result = append(result, SKU{Attributes: map[string]string{attributes[0].Name: val}})
	}

	// 依次处理剩余的属性
	for i := 1; i < len(attributes); i++ {
		newResult := []SKU{}
		// 对于每个已有SKU，添加当前属性的每个值
		for _, sku := range result {
			for _, val := range attributes[i].Values {
				// 创建新的SKU，复制现有属性并添加新属性
				newSKU := SKU{
					Attributes: make(map[string]string),
				}
				// 复制现有属性
				for k, v := range sku.Attributes {
					newSKU.Attributes[k] = v
				}
				// 添加新属性
				newSKU.Attributes[attributes[i].Name] = val
				newResult = append(newResult, newSKU)
			}
		}
		result = newResult
	}
	return result
}

// GenerateSkus 生成所有可能的SKU组合
func GenerateSkus(attributes []Attribute) []SKU {
	result := []SKU{}
	if len(attributes) == 0 {
		return result
	}
	// 初始化结构 使用第一个属性值作为初始组合
	for _, val := range attributes[0].Values {
		result = append(result, SKU{
			Attributes: map[string]string{
				attributes[0].Name: val,
			},
		})
	}
	// 依次处理剩余的属性
	for i := 1; i < len(attributes); i++ {
		newResult := []SKU{}
		for _, sku := range result {
			for _, val := range attributes[i].Values {
				// 创建新的SKU
				newSku := SKU{
					Attributes: make(map[string]string),
				}
				// 复制现有SKU
				for k, v := range sku.Attributes {
					newSku.Attributes[k] = v
				}
				// 添加新的SKU属性
				newSku.Attributes[attributes[i].Name] = val
				newResult = append(newResult, newSku)
			}
		}
		// 将心的 组合 复制到 result
		result = newResult
	}
	return result
}

// 优化后：使用预分配和对象池
func GenerateSKUsOptimized(attributes []Attribute) []SKU {
	total := 1
	// 预计算SKU总数
	for _, attr := range attributes {
		total *= len(attr.Values)
	}
	// 预分配结果切片
	result := make([]SKU, total)
	for i := range result {
		result[i].Attributes = make(map[string]string)
	}
	// 生产所有组合
	for i, attr := range attributes {
		step := 1
		for j := 0; j < i; j++ {
			step *= len(attributes[j].Values)
		}
		for k := 0; k < total; k++ {
			idx := (k / step) % len(attr.Values)
			result[k].Attributes[attr.Name] = attr.Values[idx]
		}
	}
	return result
}

// String 返回SKU的字符串表示，用于展示
func (s SKU) String() string {
	var parts []string
	for k, v := range s.Attributes {
		parts = append(parts, fmt.Sprintf("%s:%s", k, v))
	}
	return strings.Join(parts, ", ")
}
