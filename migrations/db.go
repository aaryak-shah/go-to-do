package main

import (
	"fmt"
	"gotodo/db"
	"gotodo/initializers"
	"gotodo/models"
)

func init() {
	initializers.LoadEnv()
	db.AttachInstance()
	fmt.Println("Succesful initialization")
}

func main() {
	db.Instance.AutoMigrate(&models.User{})
	db.Instance.AutoMigrate(&models.ToDo{})
	db.Instance.AutoMigrate(&models.Item{})
}
