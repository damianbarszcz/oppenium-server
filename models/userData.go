package models

import "gorm.io/gorm"

type UserData struct {
	gorm.Model

	User_id      uint   `json:"user_id"`
	First_name   string `json:"first_name"`
	Last_name    string `json:"last_name"`
	Phone_number string `json:"phone_number"`
}
