package repository

import (
	"GITS-Kelompok-CRUD-JWT/config"
	"GITS-Kelompok-CRUD-JWT/entity"
	"GITS-Kelompok-CRUD-JWT/helper"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Registration(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user entity.User
		user.Username = c.FormValue("username")
		user.Password = c.FormValue("password")

		//Check user exist
		result := db.Where("username = ?", user.Username).Find(&user)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected > 0 {
			return c.JSON(http.StatusConflict, helper.ErrorLog(http.StatusConflict, "username already used", "EXT_REF"))
		}

		//Hashing password
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
		if err != nil {
			return c.JSON(http.StatusConflict, helper.ErrorLog(http.StatusConflict, "Error while hashing password", "EXT_REF"))
		}

		//Token
		cfg, _ := config.NewConfig(".env")
		user.Password = string(hash)
		user.Token = helper.JwtGenerator(user.Username, cfg.JWTConfig.SecretKey)

		//Post Registration
		regisResult := db.Create(&user)
		if regisResult.Error != nil {
			return c.JSON(http.StatusConflict, helper.ErrorLog(http.StatusConflict, "Error when registration", "EXT_REF"))
		}

		return c.JSON(http.StatusOK, &user)
	}
}

func Login(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var userInput entity.User
		var userResult entity.User
		userInput.Username = c.FormValue("username")
		userInput.Password = c.FormValue("password")

		//Login
		resLogin := db.Where("username = ?", userInput.Username).Find(&userResult)
		if resLogin.Error != nil {
			return resLogin.Error
		}

		//Check user exist
		if resLogin.RowsAffected == 0 {
			return c.JSON(http.StatusConflict, helper.ErrorLog(http.StatusConflict, "Invalid Username", "EXT_REF"))
		}

		//Check password
		checkPass := bcrypt.CompareHashAndPassword([]byte(userResult.Password), []byte(userInput.Password))
		if checkPass != nil {
			return c.JSON(http.StatusConflict, helper.ErrorLog(http.StatusConflict, "Invalid Password", "EXT_REF"))
		}

		//Token
		cfg, _ := config.NewConfig(".env")
		userResult.Token = helper.JwtGenerator(userResult.Username, cfg.JWTConfig.SecretKey)

		//Login success result
		return c.JSON(http.StatusOK, &userResult)
	}
}
