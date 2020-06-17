package models

import (
	"github.com/jinzhu/gorm"
)

type Menu struct {
	gorm.Model
	ParentId  string `json:"parentId" gorm:"comment:'父菜单ID'"`
	Path      string `json:"path" gorm:"comment:'路由path'"`
	Name      string `json:"name" gorm:"comment:'路由name'"`
	Hidden    int    `json:"hidden" gorm:"comment:'是否在列表隐藏'"`
	Component string `json:"component" gorm:"comment:'对应前端文件路径'"`
	Sort      int    `json:"sort" gorm:"comment:'排序标记'"`
	Title     string `json:"title" gorm:"comment:'菜单名'"`
	Icon      string `json:"icon" gorm:"comment:'菜单图标'"`
	Children  []Menu `json:"children"`
}
