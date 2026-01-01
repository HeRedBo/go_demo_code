package buildtree

import (
	"github.com/gookit/goutil/dump"
	"testing"
)

func TestBuildTreeByRecursion(t *testing.T) {
	// 模拟数据库查询的原始分类数据
	rawCategories := []*Category{
		{ID: 1, ParentID: 0, Name: "电子产品"},
		{ID: 2, ParentID: 0, Name: "生活用品"},
		{ID: 3, ParentID: 1, Name: "手机"},
		{ID: 4, ParentID: 1, Name: "电脑"},
		{ID: 5, ParentID: 3, Name: "智能手机"},
		{ID: 6, ParentID: 3, Name: "功能手机"},
		{ID: 7, ParentID: 4, Name: "笔记本电脑"},
		{ID: 8, ParentID: 2, Name: "洗漱用品"},
	}
	// 构建无限极分类
	//tree := BuildTreeByRecursion(rawCategories)
	tree := BuildTreeByIteration(rawCategories)
	// 打印结果（直观展示嵌套结构）
	dump.P(tree)
	//printTree(tree, 0)

}
