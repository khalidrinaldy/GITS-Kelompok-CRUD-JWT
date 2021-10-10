package entity

import "time"

type Transaction struct {
	ID               int `gorm:"primary_key, AUTO_INCREMENT"`
	SupplierID       int
	UserID           int
	Product          []Product `gorm:"many2many:transaction_products;"`
	Total_price      float64
	Transaction_date time.Time
}
