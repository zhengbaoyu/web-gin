package menu_service

import (
	"github.com/jinzhu/gorm"
	"web-gin/common/pkg"
	"web-gin/common/pkg/db"
	"web-gin/models"
	"web-gin/repositorys/menu_repository"
)

type UpdateMenuService struct {
	Id        uint
	ParentId  string
	Path      string
	Name      string
	Hidden    int
	Component string
	Sort      int
	Title     string
	Icon      string
}

func (m *UpdateMenuService) UpdateMenu() int {
	//校验
	var menu models.Menu
	//校验title和name是否已存在
	if err := db.DB.Where("name = ? or title = ?", m.Name, m.Title).Where("id <> ?", m.Id).Find(&menu).Error; err != nil {
		menuM := models.Menu{
			Model:     gorm.Model{ID: m.Id},
			ParentId:  m.ParentId,
			Path:      m.Path,
			Name:      m.Name,
			Hidden:    m.Hidden,
			Component: m.Component,
			Sort:      m.Sort,
			Title:     m.Title,
			Icon:      m.Icon,
		}
		err := menu_repository.Update(menuM)
		if err != nil {
			return pkg.ERR_ADD_MENU_FAIL
		}
	}
	if menu.Name == m.Name {
		return pkg.ERR_ADD_MENU_EXIST_NAME
	}
	if menu.Title == m.Title {
		return pkg.ERR_ADD_MENU_EXIST_TITLE
	}
	return pkg.SUCCESS
}
