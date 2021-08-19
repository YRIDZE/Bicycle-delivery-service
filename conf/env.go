package conf

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var DbPassword string

var AccessSecret string
var RefreshSecret string

var AccessLifetimeMinutes int
var RefreshLifetimeMinutes int

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	DbPassword = os.Getenv("DB_PASSWORD")

	AccessSecret = os.Getenv("ACCESS_SECRET")
	RefreshSecret = os.Getenv("REFRESH_SECRET")

	AccessLifetimeMinutes, _ = strconv.Atoi(os.Getenv("ACCESS_LIFETIME_MINUTES"))
	RefreshLifetimeMinutes, _ = strconv.Atoi(os.Getenv("REFRESH_LIFETIME_MINUTES"))
}
