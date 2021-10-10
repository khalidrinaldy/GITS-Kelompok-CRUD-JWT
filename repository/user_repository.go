package repository

import (
	"GITS-Kelompok-CRUD-JWT/entity"
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func GetUserList(db *gorm.DB) echo.HandlerFunc{
	return func(c echo.Context) error {
		var users []entity.User
		result := db.Find(&users)
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, &users)
	}
}