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
	r.HandleFunc("/getSupplierById", h.GetByID)
	r.HandleFunc("/getSuppliers", h.GetAll)
	r.HandleFunc("/updateSupplier", h.Update)
	r.HandleFunc("/deleteSupplier", h.Delete)
}

func (h *SupplierHandler) Create(w http.ResponseWriter, req *http.Request) {
	supplier := new(models.SupplierResponse)
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

	respJ, _ := json.Marshal(models.SupplierResponse{ID: s.ID, Name: s.Name, Image: s.Image})

	w.WriteHeader(http.StatusCreated)
	w.Write(respJ)
	h.logger.Infof("supplier %d successfully created", s.ID)
}

func (h *SupplierHandler) GetByID(w http.ResponseWriter, req *http.Request) {
	supplierID, _ := strconv.Atoi(req.URL.Query().Get("id"))

	order, err := h.services.GetByID(supplierID)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	respJ, _ := json.Marshal(order)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJ)
	h.logger.Infof("user successfully fetched supplier %d", supplierID)
}

func (h *SupplierHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	supplier, err := h.services.GetAll()
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	respJ, _ := json.Marshal(supplier)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJ)
	h.logger.Infof("user fetched suppliers")
}

func (h *SupplierHandler) Update(w http.ResponseWriter, req *http.Request) {
	supplier := new(models.SupplierResponse)
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
	respJ, _ := json.Marshal(s)

	w.WriteHeader(http.StatusOK)
	w.Write(respJ)
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
