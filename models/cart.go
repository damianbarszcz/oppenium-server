package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model

	Id            uint   `json:"id"`
	User_id       uint   `json:"user_id"`
	Product_title string `json:"product_title"`
	Product_price int64  `json:"product_price"`
	Product_image string `json:"product_image"`
}
