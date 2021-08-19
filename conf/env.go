package conf

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DB_PASSWORD string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	DB_PASSWORD = os.Getenv("DB_PASSWORD")

}
