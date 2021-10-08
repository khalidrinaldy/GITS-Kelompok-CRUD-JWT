package entities

type admin struct {
	id_admin       int `gorm:"primary_key, AUTO_INCREMENT"`
	nama_admin     string
	username_admin string
	password_admin string
}
