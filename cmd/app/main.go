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
	"github.com/YRIDZE/Bicycle-delivery-service/internal"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/handlers"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	log "github.com/YRIDZE/yolo-log"
	"github.com/spf13/viper"
)

func main() {
	var err error
	internal.Log, err = log.NewLogger(
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

	db, err := db_repository.NewDB(
		db_repository.Config{
			Port:     viper.GetString("db.port"),
			Username: viper.GetString("db.username"),
			DBName:   viper.GetString("db.dbname"),
			Password: conf.DbPassword,
		},
	)
	if err != nil {
		internal.Log.Fatal("Could not connected to database. Panic!")
		panic(err.Error())
	}

	userRepository := db_repository.NewUserRepository(db)
	tokenRepository := db_repository.NewTokensRepository(db)
	orderRepository := db_repository.NewOrderRepository(db)
	supplierRepository := db_repository.NewSupplierRepository(db)
	productRepository := db_repository.NewProductRepository(db)

	p := parser.NewParser(supplierRepository, productRepository)
	go p.Parse()

	userHandler := handlers.NewUserHandler(userRepository, tokenRepository)
	orderHandler := handlers.NewOrderHandler(orderRepository)

	h := handlers.NewAppHandlers(userHandler, orderHandler)

	srv := new(app.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), h.InitRoutes()); err != nil {
			internal.Log.Fatalf("error occurred while running http server: %s", err.Error())
			return
		}
	}()

	internal.Log.Info("App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	internal.Log.Info("App Shutting Down")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		internal.Log.Errorf("error occurred on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		internal.Log.Errorf("error occurred on db connection close: %s", err.Error())
	}
}
