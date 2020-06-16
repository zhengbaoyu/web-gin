package card_repository

import (
	"github.com/jinzhu/gorm"
	"web-gin/common/pkg/db"
	"web-gin/models"
)

func GetDownloadCard(cardTypeId int) ([]models.Card, error) {
	var cards []models.Card
	err := db.DB.Where("card_type_id=?", cardTypeId).Find(&cards).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return cards, err
	}
	return cards, nil
}
