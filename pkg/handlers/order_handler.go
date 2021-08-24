package handlers

import (
	"encoding/json"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/services"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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
		models.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	order.UserID = req.Context().Value("user").(*models.User).ID
	_, err := h.services.Create(order)
	if err != nil {
		models.ErrorResponse(w, "Invalid data", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Order created"))
}

func (h *OrderHandler) GetByID(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	order, err := h.services.GetByID(id)
	if err != nil {
		models.ErrorResponse(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	respJ, _ := json.Marshal(order)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJ)
}

func (h *OrderHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	order, err := h.services.GetAll(req.Context().Value("user").(*models.User).ID)
	if err != nil {
		models.ErrorResponse(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	respJ, _ := json.Marshal(order)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJ)
}

func (h *OrderHandler) Update(w http.ResponseWriter, req *http.Request) {
	order := new(models.Order)
	if err := json.NewDecoder(req.Body).Decode(&order); err != nil {
		models.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	order.UserID = req.Context().Value("user").(*models.User).ID
	err := h.services.Update(order)
	if err != nil {
		models.ErrorResponse(w, "Invalid data", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Order updated"))
}

func (h *OrderHandler) Delete(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	err := h.services.Delete(id)
	if err != nil {
		models.ErrorResponse(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Order successfully deleted"))
}
