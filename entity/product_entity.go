package entity

import (
	"time"
)

type Product struct {
	ID         int `gorm:"primary_key, AUTO_INCREMENT"`
	Name       string
	Price      float64
	Quantity   int
	Status     bool
	CategoryID int
	Created    time.Time
}
