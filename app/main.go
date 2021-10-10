package main

import (
	"GITS-Kelompok-CRUD-JWT/config"
	"GITS-Kelompok-CRUD-JWT/routes"

	"github.com/labstack/echo"
)

func main() {
	ech := echo.New()

	//Use route
	routes.InitRoute(ech)

	//Config
	cfg, _ := config.NewConfig(".env")
	
	ech.Logger.Fatal(ech.Start(":" + cfg.Port))
}