package main

import (
	"gotodo/db"
	"gotodo/initializers"
	"gotodo/models"
)

func init() {
	initializers.LoadEnv()
	db.AttachInstance()
}

func main() {
	db.Instance.AutoMigrate(&models.ToDo{})
}
