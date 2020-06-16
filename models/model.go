package models

type Model struct {
	Id        int `json:"id "gorm:"primary_key"`
	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
}

