package db

import (
	"GITS-Kelompok-CRUD-JWT/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&entity.User{})
}