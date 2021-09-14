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

type SupplierHandler struct {
	cfg      *conf.Config
	services *services.SupplierService
}

func NewSupplierHandler(cfg conf.Config, repo db_repository.SupplierRepositoryI) *SupplierHandler {
	s := services.NewSupplierService(repo)
	return &SupplierHandler{cfg: &cfg, services: s}
}

func (h *SupplierHandler) RegisterRoutes(r *http.ServeMux, appH *AppHandlers) {
	r.HandleFunc("/createSupplier", h.Create)
	r.HandleFunc("/getSupplierById", h.GetByID)
	r.HandleFunc("/getSuppliers", h.GetAll)
	r.HandleFunc("/updateSupplier", h.Update)
	r.HandleFunc("/deleteSupplier", h.Delete)
}

func (h *SupplierHandler) Create(w http.ResponseWriter, req *http.Request) {
	supplier := new(requests.SupplierRequest)

	if err := json.NewDecoder(req.Body).Decode(&supplier); err != nil {
		h.cfg.Logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}

	if err := supplier.Validate(); err != nil {
		h.cfg.Logger.Error(err)
		requests.ValidationErrorResponse(w, err)
		return
	}

	s, err := h.services.Create(supplier)
	if err != nil {
		h.cfg.Logger.Error(err.Error())
		http.Error(w, "invalid data", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&models.SupplierResponse{ID: s.ID, Name: s.Name, Image: s.Image})
	h.cfg.Logger.Infof("supplier %d successfully created", s.ID)
}

func (h *SupplierHandler) GetByID(w http.ResponseWriter, req *http.Request) {
	supplierID, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil || supplierID < 1 {
		h.cfg.Logger.Error(errors.New("invalid id parameter"))
		http.Error(w, "invalid id parameter", http.StatusNotFound)
		return
	}

	s, err := h.services.GetByID(supplierID)
	if err != nil {
		h.cfg.Logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&models.SupplierResponse{ID: s.ID, Name: s.Name, Image: s.Image, Deleted: s.Deleted})
	h.cfg.Logger.Infof("user successfully fetched supplier %d", supplierID)
}

func (h *SupplierHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	s, err := h.services.GetAll()
	if err != nil {
		h.cfg.Logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	var resp []models.SupplierResponse
	for _, x := range *s {
		resp = append(resp, models.SupplierResponse{ID: x.ID, Name: x.Name, Image: x.Image, Deleted: x.Deleted})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	h.cfg.Logger.Infof("user fetched suppliers")
}

func (h *SupplierHandler) Update(w http.ResponseWriter, req *http.Request) {
	supplier := new(requests.SupplierRequest)
	if err := json.NewDecoder(req.Body).Decode(&supplier); err != nil {
		h.cfg.Logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}

	if err := supplier.Validate(); err != nil {
		h.cfg.Logger.Error(err)
		requests.ValidationErrorResponse(w, err)
		return
	}

	s, err := h.services.Update(supplier)
	if err != nil {
		h.cfg.Logger.Error(err.Error())
		http.Error(w, "invalid data", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&models.SupplierResponse{ID: s.ID, Name: s.Name, Image: s.Image, Deleted: s.Deleted})
	h.cfg.Logger.Infof("supplier %d successfully updated", supplier.ID)
}

func (h *SupplierHandler) Delete(w http.ResponseWriter, req *http.Request) {
	supplierID, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil || supplierID < 1 {
		h.cfg.Logger.Error(errors.New("invalid id parameter"))
		http.Error(w, "invalid id parameter", http.StatusNotFound)
		return
	}

	err = h.services.Delete(int32(supplierID))
	if err != nil {
		h.cfg.Logger.Error(err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("supplier successfully deleted"))
	h.cfg.Logger.Infof("supplier %d successfully deleted", supplierID)
}
