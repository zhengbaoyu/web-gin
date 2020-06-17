package check_form

import "github.com/jinzhu/gorm"

type AddMenu struct {
	gorm.Model
	ParentId  string `form:"parent_id" valid:"Required"`
	Path      string `form:"path"`
	Name      string `form:"name" valid:"Required"`
	Hidden    int    `form:"hidden"`
	Component string `form:"component"`
	Title     string `form:"title" valid:"Required"`
	Icon      string `form:"icon"`
	Sort      int    `form:"sort"`
}
