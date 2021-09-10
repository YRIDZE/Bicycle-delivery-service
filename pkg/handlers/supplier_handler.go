package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	s := services.NewSupplierService(repo)
	return &SupplierHandler{logger: logger, services: s}
}

func (h *SupplierHandler) RegisterRoutes(r *http.ServeMux, appH *AppHandlers) {
	r.HandleFunc("/createSupplier", h.Create)
	r.HandleFunc("/getSupplierById/", h.GetByID)
	r.HandleFunc("/getSuppliers", h.GetAll)
	r.HandleFunc("/updateSupplier", h.Update)
	r.HandleFunc("/deleteSupplier/", h.Delete)
}

func (h *SupplierHandler) Create(w http.ResponseWriter, req *http.Request) {
	supplier := new(requests.SupplierRequest)

	if err := json.NewDecoder(req.Body).Decode(&supplier); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}

	s, err := h.services.Create(supplier)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "invalid data", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&models.SupplierResponse{ID: s.ID, Name: s.Name, Image: s.Image})
	h.logger.Infof("supplier %d successfully created", s.ID)
}

func (h *SupplierHandler) GetByID(w http.ResponseWriter, req *http.Request) {
	supplierID, err := requests.Params(req)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	s, err := h.services.GetByID(supplierID)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&models.SupplierResponse{ID: s.ID, Name: s.Name, Image: s.Image, Deleted: s.Deleted})
	h.logger.Infof("user successfully fetched supplier %d", supplierID)
}

func (h *SupplierHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	s, err := h.services.GetAll()
	if err != nil {
		h.logger.Error(err.Error())
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
	h.logger.Infof("user fetched suppliers")
}

func (h *SupplierHandler) Update(w http.ResponseWriter, req *http.Request) {
	supplier := new(requests.SupplierRequest)
	if err := json.NewDecoder(req.Body).Decode(&supplier); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}

	s, err := h.services.Update(supplier)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "invalid data", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&models.SupplierResponse{ID: s.ID, Name: s.Name, Image: s.Image, Deleted: s.Deleted})
	h.logger.Infof("supplier %d successfully updated", supplier.ID)
}

func (h *SupplierHandler) Delete(w http.ResponseWriter, req *http.Request) {
	supplierID, _ := strconv.Atoi(req.URL.Query().Get("id"))

	err := h.services.Delete(int32(supplierID))
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("supplier successfully deleted"))
	h.logger.Infof("supplier %d successfully deleted", supplierID)
}
