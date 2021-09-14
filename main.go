package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/YRIDZE/Bicycle-delivery-service/conf"
	"github.com/YRIDZE/Bicycle-delivery-service/parser"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/handlers"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/YRIDZE/Bicycle-delivery-service/server"
	log "github.com/YRIDZE/yolo-log"
	"github.com/spf13/viper"
)

func main() {
	ctx := context.Background()

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
	loggerContext := context.WithValue(ctx, "logger", logger)

	conf.GetEnv(loggerContext)
	db, err := db_repository.NewDB(
		loggerContext,
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

	userRepository := db_repository.NewUserRepository(db)
	tokenRepository := db_repository.NewTokensRepository(db)
	orderRepository := db_repository.NewOrderRepository(db)
	supplierRepository := db_repository.NewSupplierRepository(db)
	productRepository := db_repository.NewProductRepository(db)

	p := parser.NewParser(logger, supplierRepository, productRepository)
	go p.Parse(ctx)

	userHandler := handlers.NewUserHandler(logger, userRepository, tokenRepository)
	orderHandler := handlers.NewOrderHandler(logger, orderRepository)
	supplierHandler := handlers.NewSupplierHandler(logger, supplierRepository)
	productHandler := handlers.NewProductHandler(logger, productRepository)
	h := handlers.NewAppHandlers(userHandler, orderHandler, supplierHandler, productHandler)

	srv := new(server.Server)
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

	shutdownContext, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownContext); err != nil {
		logger.Errorf("error occurred on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logger.Errorf("error occurred on db connection close: %s", err.Error())
	}
}
