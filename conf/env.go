package conf

import (
	"context"
	"os"
	"strconv"

	yolo_log "github.com/YRIDZE/yolo-log"
	"github.com/joho/godotenv"
)

var DbPassword string

var AccessSecret string
var RefreshSecret string

var AccessLifetimeMinutes int
var RefreshLifetimeMinutes int

func GetEnv(ctx context.Context) {

	logger := ctx.Value("logger").(*yolo_log.Logger)

	if err := godotenv.Load(); err != nil {
		logger.Fatalf("Could not load .env file. Returned error was: ", err.Error())
		panic(err.Error())
	}

	DbPassword = os.Getenv("DB_PASSWORD")

	AccessSecret = os.Getenv("ACCESS_SECRET")
	RefreshSecret = os.Getenv("REFRESH_SECRET")

	AccessLifetimeMinutes, _ = strconv.Atoi(os.Getenv("ACCESS_LIFETIME_MINUTES"))
	RefreshLifetimeMinutes, _ = strconv.Atoi(os.Getenv("REFRESH_LIFETIME_MINUTES"))

}
