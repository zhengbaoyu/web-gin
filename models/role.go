package models

import (
	"github.com/jinzhu/gorm"
)

type Roles struct {
	gorm.Model
	AuthorityName string `json:"authorityName" gorm:"comment:'角色名'"`
	DeleteStatus  int    `json:"delete_status"`
	ParentId      string `json:"parentId" gorm:"comment:'父角色ID'"`
	Children      []Roles `json:"children"`
}
