package conf

import (
	"os"
	"strconv"

	"github.com/YRIDZE/Bicycle-delivery-service/internal"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var DbPassword string

var AccessSecret string
var RefreshSecret string

var AccessLifetimeMinutes int
var RefreshLifetimeMinutes int

func init() {
	if err := godotenv.Load(); err != nil {
		internal.Log.Fatalf("Could not load .env file. Returned error was: ", err.Error())
		panic(err.Error())
	}

	DbPassword = os.Getenv("DB_PASSWORD")

	AccessSecret = os.Getenv("ACCESS_SECRET")
	RefreshSecret = os.Getenv("REFRESH_SECRET")

	AccessLifetimeMinutes, _ = strconv.Atoi(os.Getenv("ACCESS_LIFETIME_MINUTES"))
	RefreshLifetimeMinutes, _ = strconv.Atoi(os.Getenv("REFRESH_LIFETIME_MINUTES"))

	if err := InitConfig(); err != nil {
		internal.Log.Error("error initializing configs")
	}
}

func InitConfig() error {
	viper.AddConfigPath("conf")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
