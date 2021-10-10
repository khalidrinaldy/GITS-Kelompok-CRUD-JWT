package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string `json:"username"`
	Password    string `json:"password"`
	Token       string `json:"token,omitempty"`
	Transaction Transaction
}
