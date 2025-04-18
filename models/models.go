package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName    string
	LastName     string
	EmailAddress string `gorm:"unique"`
	Password     string
}
