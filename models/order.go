package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model

	Id             uint      `json:"id"`
	User_id        uint      `json:"user_id"`
	Order_id       string    `json:"order_id"`
	Payment_status bool      `json:"payment_status"`
	Product_title  string    `json:"product_title"`
	Product_image  string    `json:"product_image"`
	Product_price  int64     `json:"product_price"`
	Date           time.Time `json:"date"`
}
