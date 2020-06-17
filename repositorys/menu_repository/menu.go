package menu_repository

import (
	"web-gin/common/pkg/db"
	"web-gin/models"
)

func Add(menu models.Menu) error {
	err := db.DB.Create(&menu).Error
	return err
}
