package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/YRIDZE/Bicycle-delivery-service/internal"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/services"
	"github.com/gorilla/mux"
)

type OrderHandler struct {
	services *services.OrderService
}

func NewOrderHandler(repo db_repository.OrderRepositoryI) *OrderHandler {
	s := services.NewOrderService(repo)
	return &OrderHandler{services: s}
}

func (h *OrderHandler) RegisterRoutes(r *mux.Router, appH *AppHandlers) {
	r.Handle("/order", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.Create))).Methods(http.MethodPost)
	r.Handle("/order/{id}", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.GetByID))).Methods(http.MethodGet)
	r.Handle("/orders", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.GetAll))).Methods(http.MethodGet)
	r.Handle("/order", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.Update))).Methods(http.MethodPut)
	r.Handle("/order/{id}", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.Delete))).Methods(http.MethodDelete)
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
	orderID, _ := strconv.Atoi(mux.Vars(req)["id"])
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
	orderID, _ := strconv.Atoi(mux.Vars(req)["id"])
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
