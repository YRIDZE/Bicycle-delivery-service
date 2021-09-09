package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/services"
	yolo_log "github.com/YRIDZE/yolo-log"
)

type OrderHandler struct {
	logger   *yolo_log.Logger
	services *services.OrderService
}

func NewOrderHandler(logger *yolo_log.Logger, repo db_repository.OrderRepositoryI) *OrderHandler {
	s := services.NewOrderService(repo)
	return &OrderHandler{logger: logger, services: s}
}

func (h *OrderHandler) RegisterRoutes(r *http.ServeMux, appH *AppHandlers) {
	r.Handle("/createOrder", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.Create)))
	r.Handle("/getOrderById", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.GetByID)))
	r.Handle("/getOrders", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.GetAll)))
	r.Handle("/updateOrder", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.Update)))
	r.Handle("/deleteOrder", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.Delete)))
}

func (h *OrderHandler) Create(w http.ResponseWriter, req *http.Request) {
	order := new(models.Order)
	if err := json.NewDecoder(req.Body).Decode(&order); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}
	order.UserID = req.Context().Value("user").(*models.User).ID
	o, err := h.services.Create(order)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "invalid data", http.StatusUnauthorized)
		return
	}
	respJ, _ := json.Marshal(o)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(respJ)

	h.logger.Infof("order %d successfully created by User %d", o.ID, order.UserID)
}

func (h *OrderHandler) GetByID(w http.ResponseWriter, req *http.Request) {
	orderID, _ := strconv.Atoi(req.URL.Query().Get("id"))

	order, err := h.services.GetByID(orderID)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	respJ, _ := json.Marshal(order)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJ)

	h.logger.Infof("user %d fetched order %d", order.UserID, orderID)
}

func (h *OrderHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	userID := req.Context().Value("user").(*models.User).ID
	order, err := h.services.GetAll(userID)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	respJ, _ := json.Marshal(order)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJ)

	h.logger.Infof("user %d fetched orders", userID)
}

func (h *OrderHandler) Update(w http.ResponseWriter, req *http.Request) {
	order := new(models.Order)
	if err := json.NewDecoder(req.Body).Decode(&order); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}

	order.UserID = req.Context().Value("user").(*models.User).ID
	o, err := h.services.Update(order)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "invalid data", http.StatusUnauthorized)
		return
	}
	respJ, _ := json.Marshal(o)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJ)
	h.logger.Infof("order %d successfully updated", order.ID)
}

func (h *OrderHandler) Delete(w http.ResponseWriter, req *http.Request) {
	orderID, _ := strconv.Atoi(req.URL.Query().Get("id"))

	err := h.services.Delete(orderID)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("order successfully deleted"))
	h.logger.Infof("order %d successfully deleted", orderID)
}
