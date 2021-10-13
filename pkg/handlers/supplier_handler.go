package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/requests"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/services"
	yolo_log "github.com/YRIDZE/yolo-log"
)

type SupplierHandler struct {
	logger   *yolo_log.Logger
	services *services.SupplierService
}

func NewSupplierHandler(logger *yolo_log.Logger, repo db_repository.SupplierRepositoryI) *SupplierHandler {
	return &SupplierHandler{
		logger:   logger,
		services: services.NewSupplierService(repo),
	}
}

func (h *SupplierHandler) RegisterRoutes(r *http.ServeMux, appH *AppHandlers) {
	r.HandleFunc("/createSupplier", h.Create)
	r.HandleFunc("/getSuppliers", h.GetAll)
	r.HandleFunc("/getSupplierTypes", h.GetTypes)
}

func (h *SupplierHandler) Create(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)

	supplier := new(requests.SupplierRequest)
	if err := json.NewDecoder(req.Body).Decode(&supplier); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}

	if err := supplier.Validate(); err != nil {
		h.logger.Error(err)
		requests.ValidationErrorResponse(w, err)
		return
	}

	s, err := h.services.Create(supplier)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "invalid data", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(
		&models.SupplierResponse{
			ID:    s.ID,
			Name:  s.Name,
			Type:  s.Type,
			Image: s.Image,
			WorkHours: models.WorkingHours{
				Opening: s.WorkHours.Opening,
				Closing: s.WorkHours.Closing,
			},
		},
	)
	h.logger.Infof("supplier %d successfully created", s.ID)
}

func (h *SupplierHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)

	s, err := h.services.GetAll()
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	var resp []models.SupplierResponse
	for _, x := range *s {
		resp = append(
			resp,
			models.SupplierResponse{
				ID:    x.ID,
				Name:  x.Name,
				Type:  x.Type,
				Image: x.Image,
				WorkHours: models.WorkingHours{
					Opening: x.WorkHours.Opening,
					Closing: x.WorkHours.Closing,
				},
				Deleted: x.Deleted,
			},
		)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	h.logger.Infof("fetched suppliers")
}

func (h *SupplierHandler) GetTypes(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)

	supplierTypes, err := h.services.GetTypes()
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(supplierTypes)
	h.logger.Infof("fetched supplier types")
}
