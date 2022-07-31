package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model

	Id          uint `json:"id"`
	User_id     uint `json:"user_id"`
	Total_price uint `json:"total_price"`
	Is_paid     bool `json:"is_paid"`
}
