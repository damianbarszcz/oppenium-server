package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Id            uint   `json:"id"`
	Product_title string `json:"product_title"`
	Product_price int64  `json:"product_price"`
	Product_image string `json:"product_image"`
	Description   string `json:"description"`
	Category      string `json:"category"`
	Availability  bool   `json:"availability"`
	Product_url   string `json:"product_url"`
}
