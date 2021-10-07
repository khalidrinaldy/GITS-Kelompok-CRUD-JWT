package routes

import (
	"GITS-Kelompok-CRUD-JWT/db"
	"GITS-Kelompok-CRUD-JWT/repository"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	//Init db
	database := db.OpenDatabase()

	//Call Migrate Func
	db.Migrate(database)

	//Authentication Routes
	e.POST("/register", repository.Registration(database))
	e.POST("/login", repository.Login(database))
}