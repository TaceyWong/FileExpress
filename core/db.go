package core

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (err error) {
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	return err
}
