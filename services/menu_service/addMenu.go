package menu_service

import (
	"web-gin/common/pkg"
	"web-gin/common/pkg/db"
	"web-gin/models"
	"web-gin/repositorys/menu_repository"
)

type AddMenuService struct {
	ParentId  string
	Path      string
	Name      string
	Hidden    int
	Component string
	Sort      int
	Title     string
	Icon      string
}

func (m *AddMenuService) AddMenu() int {
	var menu models.Menus
	//校验title和name是否已存在
	if err := db.DB.Where("name = ? or title = ?", m.Name, m.Title).Find(&menu).Error; err != nil {
		addMenuData := models.Menus{
			ParentId:  m.ParentId,
			Path:      m.Path,
			Name:      m.Name,
			Hidden:    m.Hidden,
			Component: m.Component,
			Sort:      m.Sort,
			Title:     m.Title,
			Icon:      m.Icon,
		}
		err := menu_repository.Add(addMenuData)
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
