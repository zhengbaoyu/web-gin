package models

type Card struct {
	Id              int    `json:"id"`
	CardIdentifyNum string `json:"card_identify_num"`
	CardPwd         string `json:"card_pwd"`
	QrCode          string `json:"qr_code"`
}
