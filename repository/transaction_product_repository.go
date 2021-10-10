package repository

import (
	"GITS-Kelompok-CRUD-JWT/entity"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func GetTransactionProducts(db *gorm.DB) echo.HandlerFunc{
	return func(c echo.Context) error {
		var transactionProduct []entity.Transaction_Products
		result := db.Find(&transactionProduct)
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, &transactionProduct)
	}
}

func AddTransactionProduct(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var transactionProduct entity.Transaction_Products
		transactionProduct.ProductID,_ = strconv.Atoi(c.FormValue("product_id"))
		transactionProduct.TransactionID,_ = strconv.Atoi(c.FormValue("transaction_id"))
		transactionProduct.Amount,_ = strconv.Atoi(c.FormValue("amount"))
		result := db.Create(&transactionProduct)
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, &transactionProduct)
	}
}

func UpdateTransactionProduct(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var transactionProduct entity.Transaction_Products
		transactionProduct.ProductID,_ = strconv.Atoi(c.FormValue("product_id"))
		transactionProduct.TransactionID,_ = strconv.Atoi(c.FormValue("transaction_id"))
		transactionProduct.Amount,_ = strconv.Atoi(c.FormValue("amount"))
		result := db.Model(&transactionProduct).Where("id = ?", c.Param("id")).Updates(entity.Transaction_Products{
			ProductID: transactionProduct.ProductID,
			TransactionID: transactionProduct.TransactionID,
			Amount: transactionProduct.Amount,
		})
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, &transactionProduct)
	}
}

func DeleteTransactionProduct(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var transactionProduct entity.Transaction_Products
		result := db.Delete(&transactionProduct, c.Param("id"))
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, &transactionProduct)
	}
}