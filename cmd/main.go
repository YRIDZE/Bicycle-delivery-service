package main

import (
	"context"
	app "github.com/YRIDZE/Bicycle-delivery-service"
	"github.com/YRIDZE/Bicycle-delivery-service/conf"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/handlers"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	if err := initConfig(); err != nil {
		log.Print("error initializing configs")
	}

	db, _ := db_repository.NewDB(db_repository.Config{
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		Password: conf.DbPassword,
	})
	userRepository := db_repository.NewUserDBRepository(db)
	tokenRepository := db_repository.NewTokensDBRepository(db)
	orderRepository := db_repository.NewOrderDBRepository(db)

	userHandler := handlers.NewUserHandler(userRepository, tokenRepository)
	orderHandler := handlers.NewOrderHandler(orderRepository)

	h := handlers.NewAppHandlers(userHandler, orderHandler)

	srv := new(app.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), h.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Print("App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("App Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Printf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("conf")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
