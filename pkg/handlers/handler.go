package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

type AppHandlers struct {
	handlers    []HandlerI
	userHandler *UserHandler
}

type HandlerI interface {
	RegisterRoutes(router *mux.Router, h *AppHandlers)
}

func NewAppHandlers(userHandler *UserHandler, handlers ...HandlerI) *AppHandlers {
	return &AppHandlers{
		handlers:    append(handlers, userHandler),
		userHandler: userHandler,
	}
}

func (h *AppHandlers) InitRoutes() http.Handler {

	router := mux.NewRouter()

	for _, x := range h.handlers {
		x.RegisterRoutes(router, h)
	}

	return router
}
