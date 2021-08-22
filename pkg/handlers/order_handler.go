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
	r.Handle("/orderC", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.Create)))
	r.Handle("/order/{id}", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.GetByID)))
	r.Handle("/orders", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.GetAll)))
	r.Handle("/orderU", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.Update)))
	r.Handle("/orderD/{id}", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.Delete)))
}

func (h *OrderHandler) Create(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		order := new(models.Order)
		if err := json.NewDecoder(req.Body).Decode(&order); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		order.UserID = req.Context().Value("user").(*models.User).ID
		_, err := h.services.Create(order)
		if err != nil {
			http.Error(w, "Invalid data", http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Order created"))

	default:
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
	}
}

func (h *OrderHandler) GetByID(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		id, _ := strconv.Atoi(mux.Vars(req)["id"])
		order, err := h.services.GetByID(id)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		respJ, _ := json.Marshal(order)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respJ)
	default:
		http.Error(w, "Only GET is allowed", http.StatusMethodNotAllowed)
	}
}

func (h *OrderHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		order, err := h.services.GetAll(req.Context().Value("user").(*models.User).ID)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		respJ, _ := json.Marshal(order)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respJ)
	default:
		http.Error(w, "Only GET is allowed", http.StatusMethodNotAllowed)
	}
}

func (h *OrderHandler) Update(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "PUT":
		order := new(models.Order)
		if err := json.NewDecoder(req.Body).Decode(&order); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		order.UserID = req.Context().Value("user").(*models.User).ID
		err := h.services.Update(order)
		if err != nil {
			http.Error(w, "Invalid data", http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Order updated"))

	default:
		http.Error(w, "Only PUT is allowed", http.StatusMethodNotAllowed)
	}
}

func (h *OrderHandler) Delete(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "DELETE":
		id, _ := strconv.Atoi(mux.Vars(req)["id"])
		err := h.services.Delete(id)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Order successfully deleted"))
	default:
		http.Error(w, "Only DELETE is allowed", http.StatusMethodNotAllowed)
	}
}
