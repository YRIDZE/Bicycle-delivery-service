package server

import (
	"net/http"
	"time"

	"github.com/YRIDZE/Bicycle-delivery-service/conf"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/handlers"
)

func InitServer(cfg *conf.Config, h *handlers.AppHandlers) *http.Server {
	srv := &http.Server{
		Addr:           ":" + cfg.ConfigServer.Port,
		Handler:        h.InitRoutes(),
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return srv
}
