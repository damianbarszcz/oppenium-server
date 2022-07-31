package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Id            uint   `json:"id"`
	Availability  bool   `json:"availability"`
	Product_price uint   `json:"product_price"`
	Description   string `json:"description"`
	Category      string `json:"category"`
	Product_image string `json:"product_image"`
	Product_title string `json:"product_title"`
	Product_url   string `json:"product_url"`
}
