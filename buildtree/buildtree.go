package buildtree

import (
	"fmt"
)

// Category 分类结构体（包含指针类型的子分类切片，实现引用语义）
type Category struct {
	ID       int64       `json:"id"`
	ParentID int64       `json:"parent_id"`
	Name     string      `json:"name"`
	Children []*Category `json:"children"` // 子分类：指针切片（引用类型，避免值拷贝）
}

// region 递归法
// 递归查找指定父ID的所有子分类
func findChildren(allCategories []*Category, parentID int64) []*Category {
	var children []*Category
	for _, cate := range allCategories {
		if cate.ParentID == parentID {
			// 递归查找当前分类的子分类，并挂载（指针引用，直接修改原对象的Children字段）
			cate.Children = findChildren(allCategories, cate.ID)
			children = append(children, cate)
		}
	}
	return children
}

// 递归法构建无限极分类
func BuildTreeByRecursion(allCategories []*Category) []*Category {
	// 查找顶级分类（ParentID=0 视为根节点）
	return findChildren(allCategories, 0)
}

// 辅助函数：递归打印分类树
func printTree(categories []*Category, level int) {
	indent := ""
	for i := 0; i < level; i++ {
		indent += "  "
	}
	for _, cate := range categories {
		fmt.Printf("%s%s（ID：%d，父ID：%d）\n", indent, cate.Name, cate.ID, cate.ParentID)
		// 递归打印子分类
		printTree(cate.Children, level+1)
	}
}

// endregion

// region 迭代法
// 迭代法构建无限极分类（高性能）

// Category 与递归法一致（指针切片实现引用语义）
// 迭代法构建无限极分类（高性能）

func BuildTreeByIteration(allCategories []*Category) []*Category {
	// 1. 构建 ID -> 分类指针 的Map，用于O(1)快速查找父分类
	cateMap := make(map[int64]*Category, len(allCategories))
	for _, cate := range allCategories {
		cateMap[cate.ID] = cate
	}
	var tree []*Category
	// 2. 迭代遍历所有分类，挂载子分类到父分类
	for _, cate := range allCategories {
		if cate.ParentID == 0 {
			// 顶级分类，直接加入结果集
			tree = append(tree, cate)
			continue
		}
		// 通过Map快速查找父分类（O(1)效率）
		parentCate, exists := cateMap[cate.ParentID]
		if exists {
			// 挂载当前分类到父分类的Children切片（指针引用，无值拷贝）
			parentCate.Children = append(parentCate.Children, cate)
		}
	}
	return tree
}

func main() {
	// 模拟原始分类数据（与递归法一致）
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
	tree := BuildTreeByIteration(rawCategories)

	// 打印结果
	printTree(tree, 0)
}

// region
