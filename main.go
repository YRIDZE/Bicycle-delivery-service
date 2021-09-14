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
)

func main() {
	ctx := context.Background()
	cfg := conf.NewConfig()

	db, err := db_repository.NewDB(
		cfg.Logger,
		db_repository.Config{
			Host:     cfg.Host,
			Port:     cfg.DBPort,
			Username: cfg.Username,
			DBName:   cfg.DBName,
			Password: cfg.DbPassword,
		},
	)
	if err != nil {
		cfg.Logger.Fatal("Could not connected to database. Panic!")
		panic(err.Error())
	}

	userRepository := db_repository.NewUserRepository(db)
	tokenRepository := db_repository.NewTokensRepository(db)
	orderRepository := db_repository.NewOrderRepository(db)
	supplierRepository := db_repository.NewSupplierRepository(db)
	productRepository := db_repository.NewProductRepository(db)

	p := parser.NewParser(cfg, supplierRepository, productRepository)
	go p.Parse(ctx)

	userHandler := handlers.NewUserHandler(cfg, userRepository, tokenRepository)
	orderHandler := handlers.NewOrderHandler(cfg, orderRepository)
	supplierHandler := handlers.NewSupplierHandler(cfg, supplierRepository)
	productHandler := handlers.NewProductHandler(cfg, productRepository)
	h := handlers.NewAppHandlers(userHandler, orderHandler, supplierHandler, productHandler)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(cfg.Port, h.InitRoutes()); err != nil {
			cfg.Logger.Fatalf("error occurred while running http server: %s", err.Error())
			return
		}
	}()

	cfg.Logger.Info("App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	cfg.Logger.Info("App Shutting Down")

	shutdownContext, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownContext); err != nil {
		cfg.Logger.Errorf("error occurred on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		cfg.Logger.Errorf("error occurred on db connection close: %s", err.Error())
	}
}
