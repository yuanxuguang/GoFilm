package model

import (
	"encoding/json"
	"log"
	"server/config"
	"server/plugin/db"
)

// Category 分类信息
type Category struct {
	Id   int64  `json:"id"`   // 分类ID
	Pid  int64  `json:"pid"`  // 父级分类ID
	Name string `json:"name"` // 分类名称
}

// CategoryTree 分类信息树形结构
type CategoryTree struct {
	*Category
	Children []*CategoryTree `json:"children"` // 子分类信息
}

// SaveCategoryTree 保存影片分类信息
func SaveCategoryTree(tree string) error {
	return db.Rdb.Set(db.Cxt, config.CategoryTreeKey, tree, config.CategoryTreeExpired).Err()
}

// GetCategoryTree 获取影片分类信息
func GetCategoryTree() CategoryTree {
	data := db.Rdb.Get(db.Cxt, config.CategoryTreeKey).Val()
	tree := CategoryTree{}
	_ = json.Unmarshal([]byte(data), &tree)
	return tree
}

// ExistsCategoryTree 查询分类信息是否存在
func ExistsCategoryTree() bool {
	exists, err := db.Rdb.Exists(db.Cxt, config.CategoryTreeKey).Result()
	if err != nil {
		log.Println("ExistsCategoryTree Error", err)
	}
	return exists == 1
}
