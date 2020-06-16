package user_repository

import (
	"github.com/jinzhu/gorm"
	"web-gin/common/pkg/db"
	"web-gin/models"
)

func Check(where interface{}) (models.User, error) {
	var user models.User
	err := db.DB.Where("phone=?", where).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return user, err
	}
	return user, nil
}
func Get(id int) (models.User, error) {
	var user models.User
	err := db.DB.Where("id=?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return user, err
	}
	return user, nil
}

func Add(user models.User) error {
	if err := db.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
