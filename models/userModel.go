package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UID      string `gorm:"primaryKey"`
	Email    string
	Password string
}
