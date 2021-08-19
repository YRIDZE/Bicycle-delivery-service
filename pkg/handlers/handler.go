package handlers

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/services"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h Handler) InitRoutes() http.Handler {

	router := mux.NewRouter()

	router.HandleFunc("/create", h.Create)
	router.HandleFunc("/getByEmail/{email}", h.GetByEmail)
	router.HandleFunc("/", h.GetAll)

	router.HandleFunc("/login", h.Login)
	router.HandleFunc("/refresh", h.Refresh)

	router.Handle("/logout", h.AuthMiddleware(http.HandlerFunc(h.Logout)))
	router.Handle("/delete", h.AuthMiddleware(http.HandlerFunc(h.Delete)))
	router.Handle("/update", h.AuthMiddleware(http.HandlerFunc(h.Update)))
	router.Handle("/profile", h.AuthMiddleware(http.HandlerFunc(h.GetProfile)))

	return router
}
