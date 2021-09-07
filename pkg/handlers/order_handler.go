package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/YRIDZE/Bicycle-delivery-service/internal"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/services"
)

type OrderHandler struct {
	services *services.OrderService
}

func NewOrderHandler(repo db_repository.OrderRepositoryI) *OrderHandler {
	s := services.NewOrderService(repo)
	return &OrderHandler{services: s}
}

func (h *OrderHandler) RegisterRoutes(r *http.ServeMux, appH *AppHandlers) {
	r.Handle("/orderC", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.Create)))
	r.Handle("/order", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.GetByID)))
	r.Handle("/orders", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.GetAll)))
	r.Handle("/orderU", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.Update)))
	r.Handle("/orderD", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.Delete)))
}

func (h *OrderHandler) Create(w http.ResponseWriter, req *http.Request) {
	order := new(models.Order)
	if err := json.NewDecoder(req.Body).Decode(&order); err != nil {
		internal.Log.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}
	order.UserID = req.Context().Value("user").(*models.User).ID
	orderID, err := h.services.Create(order)
	if err != nil {
		internal.Log.Error(err.Error())
		http.Error(w, "invalid data", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("order created"))
	internal.Log.Infof("order %d successfully created by User %d", orderID, order.UserID)

}

func (h *OrderHandler) GetByID(w http.ResponseWriter, req *http.Request) {
	orderID, _ := strconv.Atoi(req.URL.Query().Get("id"))

	order, err := h.services.GetByID(orderID)
	if err != nil {
		internal.Log.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	respJ, _ := json.Marshal(order)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJ)
	internal.Log.Infof("user %d fetched order %d", order.UserID, orderID)

}

func (h *OrderHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	userID := req.Context().Value("user").(*models.User).ID
	order, err := h.services.GetAll(userID)
	if err != nil {
		internal.Log.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	respJ, _ := json.Marshal(order)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJ)
	internal.Log.Infof("user %d fetched orders", userID)

}

func (h *OrderHandler) Update(w http.ResponseWriter, req *http.Request) {
	order := new(models.Order)
	if err := json.NewDecoder(req.Body).Decode(&order); err != nil {
		internal.Log.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}
	order.UserID = req.Context().Value("user").(*models.User).ID
	err := h.services.Update(order)
	if err != nil {
		internal.Log.Error(err.Error())
		http.Error(w, "invalid data", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Order updated"))
	internal.Log.Infof("order %d successfully updated", order.ID)

}

func (h *OrderHandler) Delete(w http.ResponseWriter, req *http.Request) {
	orderID, _ := strconv.Atoi(req.URL.Query().Get("id"))

	err := h.services.Delete(orderID)
	if err != nil {
		internal.Log.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("order successfully deleted"))
	internal.Log.Infof("order %d successfully deleted", orderID)

}
