package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemID    string `gorm:"primaryKey"`
	Details   string
	Completed bool
	ToDoID    string
}
