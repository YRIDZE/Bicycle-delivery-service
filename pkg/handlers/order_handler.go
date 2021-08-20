package handlers

import (
	"encoding/json"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *Handler) CreateOrder(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		order := new(models.Order)
		if err := json.NewDecoder(req.Body).Decode(&order); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		order.UserID = req.Context().Value("user").(*models.User).ID
		_, err := h.services.CreateOrder(order)
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

func (h *Handler) GetOrderByID(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		id, _ := strconv.Atoi(mux.Vars(req)["id"])
		order, err := h.services.GetOrderByID(id)
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

func (h *Handler) GetAllOrders(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		order, err := h.services.GetAllOrders(req.Context().Value("user").(*models.User).ID)
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

func (h *Handler) UpdateOrders(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "PUT":
		order := new(models.Order)
		if err := json.NewDecoder(req.Body).Decode(&order); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		order.UserID = req.Context().Value("user").(*models.User).ID
		err := h.services.UpdateOrder(order)
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

func (h *Handler) DeleteOrder(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "DELETE":
		id, _ := strconv.Atoi(mux.Vars(req)["id"])
		err := h.services.DeleteOrder(id)
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
