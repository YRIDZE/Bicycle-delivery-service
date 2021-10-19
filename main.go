package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/YRIDZE/Bicycle-delivery-service/conf"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/handlers"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/YRIDZE/Bicycle-delivery-service/server"
)

func main() {
	ctx := context.Background()
	cfg := conf.NewConfig()

	db := db_repository.InitDB(cfg)
	h := handlers.InitHandlers(ctx, cfg, db)
	srv := server.InitServer(cfg, h)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
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
