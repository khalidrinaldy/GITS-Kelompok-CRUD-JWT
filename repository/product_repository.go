package repository

import (
	"GITS-Kelompok-CRUD-JWT/entity"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func GetAllProducts(db *gorm.DB) echo.HandlerFunc{
	return func(c echo.Context) error {
		var products []entity.Product
		result := db.Find(&products)
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, &products)
	}
}

func AddProduct(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var product entity.Product
		product.Name = c.FormValue("name")
		product.Price,_ = strconv.ParseFloat(c.FormValue("price"), 64)
		product.Quantity,_ = strconv.Atoi(c.FormValue("quantity"))
		product.Status,_ = strconv.ParseBool(c.FormValue("status"))
		product.Created = time.Now()
		product.CategoryID,_ = strconv.Atoi(c.FormValue("category_id"))
		result := db.Create(&product)
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, &product)
	}
}

func UpdateProduct(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var product entity.Product
		product.Name = c.FormValue("name")
		product.Price,_ = strconv.ParseFloat(c.FormValue("price"), 64)
		product.Quantity,_ = strconv.Atoi(c.FormValue("quantity"))
		product.Status,_ = strconv.ParseBool(c.FormValue("status"))
		product.Created = time.Now()
		product.CategoryID,_ = strconv.Atoi(c.FormValue("category_id"))
		result := db.Model(&product).Where("id = ?", c.Param("id")).Updates(entity.Product{
			Name: product.Name,
			Price: product.Price,
			Quantity: product.Quantity,
			Status: product.Status,
			Created: product.Created,
			CategoryID: product.CategoryID,
		})
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, &product)
	}
}

func DeleteProduct(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var product entity.Product
		result := db.Delete(&product, c.Param("id"))
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, &product)
	}
}