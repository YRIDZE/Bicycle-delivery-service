package handlers

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/YRIDZE/Bicycle-delivery-service/conf"
	"github.com/YRIDZE/Bicycle-delivery-service/parser"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
)

type AppHandlers struct {
	handlers    []HandlerI
	userHandler *UserHandler
}

type HandlerI interface {
	RegisterRoutes(router *http.ServeMux, h *AppHandlers)
}

func NewAppHandlers(userHandler *UserHandler, handlers ...HandlerI) *AppHandlers {
	return &AppHandlers{
		handlers:    append(handlers, userHandler),
		userHandler: userHandler,
	}
}

func (h *AppHandlers) InitRoutes() http.Handler {
	router := http.NewServeMux()

	for _, x := range h.handlers {
		x.RegisterRoutes(router, h)
	}

	return router
}

func InitHandlers(ctx context.Context, cfg *conf.Config, db *sql.DB) *AppHandlers {
	userRepository := db_repository.NewUserRepository(db)
	tokenRepository := db_repository.NewTokensRepository(db)
	orderRepository := db_repository.NewOrderRepository(db)
	supplierRepository := db_repository.NewSupplierRepository(db)
	productRepository := db_repository.NewProductRepository(db)
	cartRepository := db_repository.NewCartRepository(db)

	p := parser.NewParser(cfg, supplierRepository, productRepository)
	go p.Parse(ctx)

	userHandler := NewUserHandler(cfg.ConfigToken, cfg.Logger, userRepository, tokenRepository)
	orderHandler := NewOrderHandler(cfg.ConfigToken, cfg.Logger, orderRepository)
	supplierHandler := NewSupplierHandler(cfg.Logger, supplierRepository)
	productHandler := NewProductHandler(cfg.Logger, productRepository)
	cartHandler := NewCartHandler(cfg.ConfigToken, cfg.Logger, cartRepository)
	h := NewAppHandlers(userHandler, orderHandler, supplierHandler, productHandler, cartHandler)

	return h
}
