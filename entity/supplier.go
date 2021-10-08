package entities

type supplier struct {
	id_supplier   int `gorm:"primary_key, AUTO_INCREMENT"`
	telp_supplier int
	nama_supplier string
}
