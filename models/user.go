package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Id       uint   `json:"id"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
	UserData UserData
}
