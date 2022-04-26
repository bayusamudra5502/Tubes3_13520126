package db

import (
	"gorm.io/gorm"
)

func GetDatabase() *gorm.DB {
	if DB == nil {
		Connect()
	}

	return DB
}
