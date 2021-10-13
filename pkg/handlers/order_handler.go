package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/YRIDZE/Bicycle-delivery-service/conf"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/requests"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/services"
	yolo_log "github.com/YRIDZE/yolo-log"
)

type OrderHandler struct {
	cfg      *conf.ConfigToken
	logger   *yolo_log.Logger
	services *services.OrderService
}

func NewOrderHandler(cfg *conf.ConfigToken, logger *yolo_log.Logger, repo db_repository.OrderRepositoryI) *OrderHandler {
	return &OrderHandler{
		cfg:      cfg,
		logger:   logger,
		services: services.NewOrderService(repo),
	}
}

func (h *OrderHandler) RegisterRoutes(r *http.ServeMux, appH *AppHandlers) {
	r.Handle("/createOrder", appH.UserHandler.AuthMiddleware(http.HandlerFunc(h.Create)))
	r.Handle("/getOrders", appH.UserHandler.AuthMiddleware(http.HandlerFunc(h.GetAll)))
}

func (h *OrderHandler) Create(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)

	order := new(requests.OrderRequest)
	if err := json.NewDecoder(req.Body).Decode(&order); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}
	order.UserID = req.Context().Value("user").(*models.User).ID
	if err := order.Validate(); err != nil {
		h.logger.Error(err.Error())
		requests.ValidationErrorResponse(w, err)
		return
	}

	o, err := h.services.Create(order)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "invalid data", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&models.OrderResponse{ID: o.ID, UserID: o.UserID, Address: o.Address, Status: o.Address, Products: o.Products})
	h.logger.Infof("order %d successfully created by User %d", o.ID, o.UserID)
}

func (h *OrderHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)

	userID := req.Context().Value("user").(*models.User).ID
	o, err := h.services.GetAll(userID)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	var resp []models.OrderResponse
	for _, x := range *o {
		resp = append(
			resp, models.OrderResponse{
				ID:               x.ID,
				UserID:           x.UserID,
				Address:          x.Address,
				PhoneNumber:      x.PhoneNumber,
				CustomerName:     x.CustomerName,
				CustomerLastname: x.CustomerLastname,
				Status:           x.Status,
				Products:         x.Products,
				CreatedAt:        x.CreatedAt,
			},
		)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	h.logger.Infof("user %d fetched orders", userID)
}
