package repository

import (
	"GITS-Kelompok-CRUD-JWT/entity"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func GetTransactions(db *gorm.DB) echo.HandlerFunc{
	return func(c echo.Context) error {
		var transactions []entity.Transaction
		result := db.Find(&transactions)
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, &transactions)
	}
}

func AddTransaction(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var transaction entity.Transaction
		transaction.SupplierID,_ = strconv.Atoi(c.FormValue("supplier_id"))
		transaction.UserID,_ = strconv.Atoi(c.FormValue("user_id"))
		transaction.Total_price,_ = strconv.ParseFloat(c.FormValue("total_price"), 64)
		transaction.Transaction_date = time.Now()
		result := db.Create(&transaction)
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, &transaction)
	}
}

func UpdateTransaction(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var transaction entity.Transaction
		transaction.SupplierID,_ = strconv.Atoi(c.FormValue("supplier_id"))
		transaction.UserID,_ = strconv.Atoi(c.FormValue("user_id"))
		transaction.Total_price,_ = strconv.ParseFloat(c.FormValue("total_price"), 64)
		transaction.Transaction_date = time.Now()
		result := db.Model(&transaction).Where("id = ?", c.Param("id")).Updates(entity.Transaction{
			SupplierID: transaction.SupplierID,
			UserID: transaction.UserID,
			Total_price: transaction.Total_price,
			Transaction_date: transaction.Transaction_date,
		})
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, &transaction)
	}
}

func DeleteTransaction(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var transaction entity.Transaction
		result := db.Delete(&transaction, c.Param("id"))
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, &transaction)
	}
}