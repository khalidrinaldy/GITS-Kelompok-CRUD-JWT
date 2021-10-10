package entity

type Supplier struct {
	ID          int `gorm:"primary_key, AUTO_INCREMENT"`
	Phone       string
	Name        string
	Transaction Transaction
}
