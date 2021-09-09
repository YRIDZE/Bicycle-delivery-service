package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	app "github.com/YRIDZE/Bicycle-delivery-service"
	"github.com/YRIDZE/Bicycle-delivery-service/cmd/parser"
	"github.com/YRIDZE/Bicycle-delivery-service/conf"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/handlers"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	log "github.com/YRIDZE/yolo-log"
	"github.com/spf13/viper"
)

func main() {
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

	c := context.WithValue(context.Background(), "logger", logger)

	conf.GetEnv(c)
	if err := InitConfig(); err != nil {
		logger.Error("error initializing configs")
	}

	db, err := db_repository.NewDB(
		c,
		db_repository.Config{
			Port:     viper.GetString("db.port"),
			Username: viper.GetString("db.username"),
			DBName:   viper.GetString("db.dbname"),
			Password: conf.DbPassword,
		},
	)
	if err != nil {
		logger.Fatal("Could not connected to database. Panic!")
		panic(err.Error())
	}
	c.Done()

	userRepository := db_repository.NewUserRepository(db)
	tokenRepository := db_repository.NewTokensRepository(db)
	orderRepository := db_repository.NewOrderRepository(db)
	supplierRepository := db_repository.NewSupplierRepository(db)
	productRepository := db_repository.NewProductRepository(db)

	p := parser.NewParser(logger, supplierRepository, productRepository)
	go p.Parse()

	userHandler := handlers.NewUserHandler(logger, userRepository, tokenRepository)
	orderHandler := handlers.NewOrderHandler(logger, orderRepository)
	supplierHandler := handlers.NewSupplierHandler(logger, supplierRepository)
	productHandler := handlers.NewProductHandler(logger, productRepository)

	h := handlers.NewAppHandlers(userHandler, orderHandler, supplierHandler, productHandler)

	srv := new(app.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), h.InitRoutes()); err != nil {
			logger.Fatalf("error occurred while running http server: %s", err.Error())
			return
		}
	}()

	logger.Info("App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Info("App Shutting Down")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Errorf("error occurred on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logger.Errorf("error occurred on db connection close: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("conf")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
