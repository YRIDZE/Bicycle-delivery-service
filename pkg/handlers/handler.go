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
	router.Handle("/profile", h.AuthMiddleware(http.HandlerFunc(h.GetProfile)))
	router.Handle("/update", h.AuthMiddleware(http.HandlerFunc(h.Update)))
	router.Handle("/delete", h.AuthMiddleware(http.HandlerFunc(h.Delete)))
	router.HandleFunc("/create", h.Create)
	router.HandleFunc("/getByEmail/{email}", h.GetByEmail)
	router.HandleFunc("/", h.GetAll)

	//orders
	router.Handle("/createOrder", h.AuthMiddleware(http.HandlerFunc(h.CreateOrder)))
	router.Handle("/getOrder/{id}", h.AuthMiddleware(http.HandlerFunc(h.GetOrderByID)))
	router.Handle("/getOrders", h.AuthMiddleware(http.HandlerFunc(h.GetAllOrders)))
	router.Handle("/updateOrder", h.AuthMiddleware(http.HandlerFunc(h.UpdateOrders)))
	router.Handle("/deleteOrder/{id}", h.AuthMiddleware(http.HandlerFunc(h.DeleteOrder)))

	return router
}
