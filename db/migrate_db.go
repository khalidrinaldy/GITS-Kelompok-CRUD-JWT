package db

import (
	"GITS-Kelompok-CRUD-JWT/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Category{})
	db.AutoMigrate(&entity.Product{})
	db.AutoMigrate(&entity.Supplier{})
	db.AutoMigrate(&entity.Transaction{})
	db.AutoMigrate(&entity.Transaction_Products{})
}