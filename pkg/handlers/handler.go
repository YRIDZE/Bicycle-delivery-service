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

	router.HandleFunc("/login", h.Login)
	router.HandleFunc("/refresh", h.Refresh)

	//user
	router.Handle("/logout", h.AuthMiddleware(http.HandlerFunc(h.Logout)))
	router.Handle("/user", h.AuthMiddleware(http.HandlerFunc(h.GetUserProfile)))
	router.Handle("/userU", h.AuthMiddleware(http.HandlerFunc(h.UpdateUser)))
	router.Handle("/userD", h.AuthMiddleware(http.HandlerFunc(h.DeleteUser)))
	router.HandleFunc("/userC", h.CreateUser)
	router.HandleFunc("/user/{email}", h.GetUserByEmail)
	router.HandleFunc("/users", h.GetAllUsers)

	//orders
	router.Handle("/orderC", h.AuthMiddleware(http.HandlerFunc(h.CreateOrder)))
	router.Handle("/order/{id}", h.AuthMiddleware(http.HandlerFunc(h.GetOrderByID)))
	router.Handle("/orders", h.AuthMiddleware(http.HandlerFunc(h.GetAllOrders)))
	router.Handle("/orderU", h.AuthMiddleware(http.HandlerFunc(h.UpdateOrders)))
	router.Handle("/orderD/{id}", h.AuthMiddleware(http.HandlerFunc(h.DeleteOrder)))

	return router
}
