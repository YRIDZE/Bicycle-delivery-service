package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/YRIDZE/Bicycle-delivery-service/conf"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/requests"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/services"
)

type OrderHandler struct {
	cfg      *conf.Config
	services *services.OrderService
}

func NewOrderHandler(cfg conf.Config, repo db_repository.OrderRepositoryI) *OrderHandler {
	s := services.NewOrderService(repo)
	return &OrderHandler{cfg: &cfg, services: s}
}

func (h *OrderHandler) RegisterRoutes(r *http.ServeMux, appH *AppHandlers) {
	r.Handle("/createOrder", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.Create)))
	r.Handle("/getOrderById", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.GetByID)))
	r.Handle("/getOrders", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.GetAll)))
	r.Handle("/updateOrder", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.Update)))
	r.Handle("/deleteOrder", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.Delete)))
}

func (h *OrderHandler) Create(w http.ResponseWriter, req *http.Request) {
	order := new(requests.OrderRequest)
	if err := json.NewDecoder(req.Body).Decode(&order); err != nil {
		h.cfg.Logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}
	order.UserID = req.Context().Value("user").(*models.User).ID
	if err := order.Validate(); err != nil {
		h.cfg.Logger.Error(err.Error())
		requests.ValidationErrorResponse(w, err)
		return
	}

	o, err := h.services.Create(order)
	if err != nil {
		h.cfg.Logger.Error(err.Error())
		http.Error(w, "invalid data", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&models.OrderResponse{ID: o.ID, UserID: o.UserID, Address: o.Address, Status: o.Address, Products: o.Products})
	h.cfg.Logger.Infof("order %d successfully created by User %d", o.ID, o.UserID)
}

func (h *OrderHandler) GetByID(w http.ResponseWriter, req *http.Request) {
	orderID, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil || orderID < 1 {
		h.cfg.Logger.Error(errors.New("invalid id parameter"))
		http.Error(w, "invalid id parameter", http.StatusNotFound)
		return
	}

	o, err := h.services.GetByID(orderID)
	if err != nil {
		h.cfg.Logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&models.OrderResponse{ID: o.ID, UserID: o.UserID, Address: o.Address, Status: o.Address, Products: o.Products})
	h.cfg.Logger.Infof("user %d fetched order %d", o.UserID, o.ID)
}

func (h *OrderHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	userID := req.Context().Value("user").(*models.User).ID
	o, err := h.services.GetAll(userID)
	if err != nil {
		h.cfg.Logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	var resp []models.OrderResponse
	for _, x := range *o {
		resp = append(
			resp, models.OrderResponse{ID: x.ID, UserID: x.UserID, Address: x.Address, Status: x.Address, Products: x.Products},
		)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	h.cfg.Logger.Infof("user %d fetched orders", userID)
}

func (h *OrderHandler) Update(w http.ResponseWriter, req *http.Request) {
	order := new(requests.OrderRequest)
	if err := json.NewDecoder(req.Body).Decode(&order); err != nil {
		h.cfg.Logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}

	order.UserID = req.Context().Value("user").(*models.User).ID
	if err := order.Validate(); err != nil {
		h.cfg.Logger.Error(err.Error())
		requests.ValidationErrorResponse(w, err)
		return
	}

	o, err := h.services.Update(order)
	if err != nil {
		h.cfg.Logger.Error(err.Error())
		http.Error(w, "invalid data", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&models.OrderResponse{ID: o.ID, UserID: o.UserID, Address: o.Address, Status: o.Address, Products: o.Products})
	h.cfg.Logger.Infof("order %d successfully updated", o.ID)
}

func (h *OrderHandler) Delete(w http.ResponseWriter, req *http.Request) {
	orderID, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil || orderID < 1 {
		h.cfg.Logger.Error(errors.New("invalid id parameter"))
		http.Error(w, "invalid id parameter", http.StatusNotFound)
		return
	}

	err = h.services.Delete(orderID)
	if err != nil {
		h.cfg.Logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("order successfully deleted"))
	h.cfg.Logger.Infof("order %d successfully deleted", orderID)
}
