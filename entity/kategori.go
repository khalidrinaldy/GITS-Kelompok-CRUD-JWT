package entities

type kategori struct {
	id_kategori   int `gorm:"primary_key, AUTO_INCREMENT"`
	nama_kategori string
}
