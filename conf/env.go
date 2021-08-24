package conf

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
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
		logrus.Print("No .env file found")
	}

	DbPassword = os.Getenv("DB_PASSWORD")

	AccessSecret = os.Getenv("ACCESS_SECRET")
	RefreshSecret = os.Getenv("REFRESH_SECRET")

	AccessLifetimeMinutes, _ = strconv.Atoi(os.Getenv("ACCESS_LIFETIME_MINUTES"))
	RefreshLifetimeMinutes, _ = strconv.Atoi(os.Getenv("REFRESH_LIFETIME_MINUTES"))
}
