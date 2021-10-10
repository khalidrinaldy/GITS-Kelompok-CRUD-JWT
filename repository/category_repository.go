package repository

import (
	"GITS-Kelompok-CRUD-JWT/entity"
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func GetCategories(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var categories []entity.Category
		result := db.Find(&categories)
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, &categories)
	}
}

func AddCategory(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var category entity.Category
		category.Name = c.FormValue("name")
		result := db.Create(&category)
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, &category)
	}
}

func UpdateCategory(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var category entity.Category
		category.Name = c.FormValue("name")
		result := db.Model(&category).Where("id = ?", c.Param("id")).Update("name", category.Name)
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, &category)
	}
}

func DeleteCategory(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var category entity.Category
		result := db.Delete(&category, c.Param("id"))
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, &category)
	}
}