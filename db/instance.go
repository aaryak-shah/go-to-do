package db

import (
	"fmt"
	"log"
	"os"

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
