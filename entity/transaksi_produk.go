package entity

type Transaction_Products struct {
	ID            int `gorm:"primary_key, AUTO_INCREMENT"`
	TransactionID int `gorm:"primary_key"`
	ProductID     int `gorm:"primary_key"`
	Amount        int
}
