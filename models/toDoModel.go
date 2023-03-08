package models

import "gorm.io/gorm"

type ToDo struct {
	gorm.Model
	ToDoID string `gorm:"primaryKey"`
	Title  string
	UID    string
}
