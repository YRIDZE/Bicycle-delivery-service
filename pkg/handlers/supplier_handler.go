package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/YRIDZE/Bicycle-delivery-service/internal"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/services"
)

type SupplierHandler struct {
	services *services.SupplierService
}

func NewSupplierHandler(repo db_repository.SupplierRepositoryI) *SupplierHandler {
	s := services.NewSupplierService(repo)
	return &SupplierHandler{services: s}
}

func (h *SupplierHandler) RegisterRoutes(r *http.ServeMux, appH *AppHandlers) {
	r.HandleFunc("/supplierC", h.Create)
	r.HandleFunc("/supplier", h.GetByID)
	r.HandleFunc("/suppliers", h.GetAll)
	r.HandleFunc("/supplierU", h.Update)
	r.HandleFunc("/supplierD", h.Delete)
}

func (h *SupplierHandler) Create(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		supplier := new(models.Supplier)
		if err := json.NewDecoder(req.Body).Decode(&supplier); err != nil {
			internal.Log.Error(err.Error())
			http.Error(w, "something went wrong", http.StatusBadRequest)
			return
		}
		supplierID, err := h.services.Create(supplier)
		if err != nil {
			internal.Log.Error(err.Error())
			http.Error(w, "invalid data", http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("supplier created"))
		internal.Log.Infof("supplier %d successfully created", supplierID)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func (h *SupplierHandler) GetByID(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		supplierID, _ := strconv.Atoi(req.URL.Query().Get("id"))

		order, err := h.services.GetByID(supplierID)
		if err != nil {
			internal.Log.Error(err.Error())
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		respJ, _ := json.Marshal(order)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respJ)
		internal.Log.Infof("user successfully fetched supplier %d", supplierID)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func (h *SupplierHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		supplier, err := h.services.GetAll()
		if err != nil {
			internal.Log.Error(err.Error())
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		respJ, _ := json.Marshal(supplier)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respJ)
		internal.Log.Infof("user fetched suppliers")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func (h *SupplierHandler) Update(w http.ResponseWriter, req *http.Request) {
	if req.Method == "PUT" {
		supplier := new(models.SupplierResponse)
		if err := json.NewDecoder(req.Body).Decode(&supplier); err != nil {
			internal.Log.Error(err.Error())
			http.Error(w, "something went wrong", http.StatusBadRequest)
			return
		}

		err := h.services.Update(supplier)
		if err != nil {
			internal.Log.Error(err.Error())
			http.Error(w, "invalid data", http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("supplier updated"))
		internal.Log.Infof("supplier %d successfully updated", supplier.ID)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func (h *SupplierHandler) Delete(w http.ResponseWriter, req *http.Request) {
	if req.Method == "DELETE" {
		supplierID, _ := strconv.Atoi(req.URL.Query().Get("id"))

		err := h.services.Delete(int32(supplierID))
		if err != nil {
			internal.Log.Error(err.Error())
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("supplier successfully deleted"))
		internal.Log.Infof("supplier %d successfully deleted", supplierID)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
