package conf

import (
	"os"
	"strconv"

	log "github.com/YRIDZE/yolo-log"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type ConfigDB struct {
	Host       string
	Port       string
	Username   string
	DBName     string
	DbPassword string
}

type ConfigServer struct {
	Port string
}

type ConfigToken struct {
	AccessSecret           string
	RefreshSecret          string
	AccessLifetimeMinutes  int
	RefreshLifetimeMinutes int
}

type Config struct {
	ConfigServer
	ConfigToken
	ConfigDB
	Logger *log.Logger
}

func NewConfig() *Config {

	logger, err := log.NewLogger(
		log.LoggerParams{
			ConsoleOutputStream: os.Stdout,
			ConsoleLogLevel:     log.INFO,
			LogFileName:         "logs/all.log",
			FileLogLevel:        log.DEBUG,
		},
	)
	if err != nil {
		panic(err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logger.Fatalf("Could not load .env file. Returned error was: %v", err.Error())
		panic(err.Error())
	}

	if err := InitConfig(); err != nil {
		logger.Fatalf("error initializing configs")
	}

	accessLifetimeMinutes, _ := strconv.Atoi(os.Getenv("ACCESS_LIFETIME_MINUTES"))
	refreshLifetimeMinutes, _ := strconv.Atoi(os.Getenv("REFRESH_LIFETIME_MINUTES"))

	return &Config{
		ConfigServer: ConfigServer{
			Port: viper.GetString("port"),
		},
		ConfigToken: ConfigToken{
			AccessSecret:           os.Getenv("ACCESS_SECRET"),
			RefreshSecret:          os.Getenv("REFRESH_SECRET"),
			AccessLifetimeMinutes:  accessLifetimeMinutes,
			RefreshLifetimeMinutes: refreshLifetimeMinutes,
		},
		ConfigDB: ConfigDB{
			Host:       viper.GetString("db.host"),
			Port:       viper.GetString("db.port"),
			Username:   viper.GetString("db.username"),
			DBName:     viper.GetString("db.dbname"),
			DbPassword: os.Getenv("DB_PASSWORD"),
		},
		Logger: logger,
	}
}

func InitConfig() error {
	viper.AddConfigPath("conf")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
