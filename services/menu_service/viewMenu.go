package menu_service

import (
	"github.com/jinzhu/gorm"
	"web-gin/models"
	"web-gin/repositorys/menu_repository"
)

func ViewMenu(id uint) (models.Menus, error) {
	menuM := models.Menus{
		Model: gorm.Model{ID: id},
	}
	return menu_repository.View(menuM)
}
