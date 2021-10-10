package routes

import (
	"GITS-Kelompok-CRUD-JWT/db"
	"GITS-Kelompok-CRUD-JWT/repository"

	"github.com/labstack/echo"
)

func InitRoute(e *echo.Echo) {
	//Init db
	database := db.OpenDatabase()

	//Call Migrate Func
	db.Migrate(database)

	//Authentication Routes
	e.POST("/register", repository.Registration(database))
	e.POST("/login", repository.Login(database))

	//Category Routes
	e.GET("/category", repository.GetCategories(database))
	e.POST("/category", repository.AddCategory(database))
	e.PUT("/category/:id", repository.UpdateCategory(database))
	e.DELETE("/category/:id", repository.DeleteCategory(database))

	//User Routes
	e.GET("/user", repository.GetUserList(database))

	//Supplier Routes
	e.GET("/supplier", repository.GetSuppliers(database))
	e.POST("/supplier", repository.AddSupplier(database))
	e.PUT("/supplier/:id", repository.UpdateSupplier(database))
	e.DELETE("/supplier/:id", repository.DeleteSupplier(database))

	//Product Routes
	e.GET("/product", repository.GetAllProducts(database))
	e.POST("/product", repository.AddProduct(database))
	e.PUT("/product/:id", repository.UpdateProduct(database))
	e.DELETE("/product/:id", repository.DeleteProduct(database))

	//Transaction Routes
	e.GET("/transaction", repository.GetTransactions(database))
	e.POST("/transaction", repository.AddTransaction(database))
	e.PUT("/transaction/:id", repository.UpdateTransaction(database))
	e.DELETE("/transaction/:id", repository.DeleteTransaction(database))

	//Transaction Product Routes
	e.GET("/transaction-product", repository.GetTransactionProducts(database))
	e.POST("/transaction-product", repository.AddTransactionProduct(database))
	e.PUT("/transaction-product/:id", repository.UpdateTransactionProduct(database))
	e.DELETE("/transaction-product/:id", repository.DeleteTransactionProduct(database))
}