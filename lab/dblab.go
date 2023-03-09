package main

import (
	"fmt"
	"gotodo/db"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func AttachInstance() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("POSTGRES_SERVER"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASS"),
		os.Getenv("POSTGRES_DBNAME"),
	)

	var instanceConnectionError error
	Instance, instanceConnectionError = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if instanceConnectionError != nil {
		fmt.Println("ERROR: Couldn't connect to postgres instance")
		log.Fatal(instanceConnectionError)
	}

	return Instance
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("ERROR: Couldn't load environment")
		log.Fatal(err)
	}
}

func init() {
	LoadEnv()
	AttachInstance()
}

func main() {
	result := db.Instance.Where("email = ?", user.Email).First(&user)
	fmt.Printf("Result: %#v", result)
}
