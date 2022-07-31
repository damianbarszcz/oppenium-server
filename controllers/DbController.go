package controllers

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db, Err = gorm.Open(sqlite.Open("oppenium.db"), &gorm.Config{})
