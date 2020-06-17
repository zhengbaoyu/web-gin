package menu_repository

import (
	"github.com/jinzhu/gorm"
	"web-gin/common/pkg/db"
	"web-gin/models"
)

//添加
func Add(menu models.Menu) error {
	err := db.DB.Create(&menu).Error
	return err
}

//查看详情
func View(menu models.Menu) (models.Menu, error) {
	var menuInfo models.Menu
	if err := db.DB.Where("id=?", menu.ID).First(&menuInfo).Error; err != nil {
		return menuInfo, err
	}
	return menuInfo, nil
}

//编辑
func Update(menu models.Menu) error {
	err := db.DB.Model(&models.Menu{}).Where("id=?", menu.ID).Update(menu).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return nil
}
