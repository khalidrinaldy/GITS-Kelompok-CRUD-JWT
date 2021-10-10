package entity

type Category struct {
	ID      int `gorm:"primary_key, AUTO_INCREMENT"`
	Name    string
	Product Product
}
