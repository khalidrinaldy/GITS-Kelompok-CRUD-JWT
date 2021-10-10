package repository

import (
	"GITS-Kelompok-CRUD-JWT/entity"
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func GetSuppliers(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var suppliers []entity.Supplier
		result := db.Find(&suppliers)
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, &suppliers)
	}
}

func AddSupplier(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var supplier entity.Supplier
		supplier.Name = c.FormValue("name")
		supplier.Phone = c.FormValue("phone")
		result := db.Create(&supplier)
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, &supplier)
	}
}

func UpdateSupplier(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var supplier entity.Supplier
		supplier.Name = c.FormValue("name")
		supplier.Phone = c.FormValue("phone")
		result := db.Model(&supplier).Where("id = ?", c.Param("id")).Updates(entity.Supplier{
			Name: supplier.Name,
			Phone: supplier.Phone,
		})
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, &supplier)
	}
}

func DeleteSupplier(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var supplier entity.Supplier
		result := db.Delete(&supplier, c.Param("id"))
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, &supplier)
	}
}
