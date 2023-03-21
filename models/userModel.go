package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UID      string `gorm:"primaryKey"`
	Email    string `gorm:"unique"`
	Password string
	ToDos    []ToDo `gorm:"foreignKey:UID"`
}
