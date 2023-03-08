package initializers

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("ERROR: Couldn't load environment")
		log.Fatal(err)
	}
}
