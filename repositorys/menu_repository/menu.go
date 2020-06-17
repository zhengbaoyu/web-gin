package menu_repository

import (
	"github.com/jinzhu/gorm"
	"web-gin/common/pkg/db"
	"web-gin/models"
)

//添加
func Add(menu models.Menus) error {
	err := db.DB.Create(&menu).Error
	return err
}

//查看详情
func View(menu models.Menus) (models.Menus, error) {
	var menuInfo models.Menus
	if err := db.DB.Where("id=?", menu.ID).First(&menuInfo).Error; err != nil {
		return menuInfo, err
	}
	return menuInfo, nil
}

//编辑
func Update(menu models.Menus) error {
	err := db.DB.Model(&models.Menus{}).Where("id=?", menu.ID).Update(menu).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return nil
}

//删除
func Delete(id uint) error {
	menu := models.Menus{DeleteStatus: 1}
	err := db.DB.Model(&models.Menus{}).Where("id=?", id).Update(&menu).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return nil
}
