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

type ProductHandler struct {
	services *services.ProductService
}

func NewProductHandler(repo db_repository.ProductRepositoryI) *ProductHandler {
	s := services.NewProductService(repo)
	return &ProductHandler{services: s}
}

func (h *ProductHandler) RegisterRoutes(r *http.ServeMux, appH *AppHandlers) {
	r.HandleFunc("/productC", h.Create)
	r.HandleFunc("/product", h.GetByID)
	r.HandleFunc("/products", h.GetAll)
	r.HandleFunc("/products/", h.GetBySupplier)
	r.HandleFunc("/productU", h.Update)
	r.HandleFunc("/productD", h.Delete)
}

func (h *ProductHandler) Create(w http.ResponseWriter, req *http.Request) {
	product := new(models.Product)
	if err := json.NewDecoder(req.Body).Decode(&product); err != nil {
		internal.Log.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}

	productID, err := h.services.Create(product)
	if err != nil {
		internal.Log.Error(err.Error())
		http.Error(w, "invalid data", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("product created"))
	internal.Log.Infof("product %d successfully created", productID)
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, req *http.Request) {
	productID, _ := strconv.Atoi(req.URL.Query().Get("id"))

	order, err := h.services.GetByID(productID)
	if err != nil {
		internal.Log.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	respJ, _ := json.Marshal(order)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJ)
	internal.Log.Infof("user successfully fetched product %d", productID)
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	product, err := h.services.GetAll()
	if err != nil {
		internal.Log.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	respJ, _ := json.Marshal(product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJ)
	internal.Log.Infof("user fetched all products")

}

func (h *ProductHandler) Update(w http.ResponseWriter, req *http.Request) {
	product := new(models.Product)
	if err := json.NewDecoder(req.Body).Decode(&product); err != nil {
		internal.Log.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}

	err := h.services.Update(product)
	if err != nil {
		internal.Log.Error(err.Error())
		http.Error(w, "invalid data", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("product updated"))
	internal.Log.Infof("product %d successfully updated", product.ID)

}

func (h *ProductHandler) Delete(w http.ResponseWriter, req *http.Request) {
	productID, _ := strconv.Atoi(req.URL.Query().Get("id"))

	err := h.services.Delete(productID)
	if err != nil {
		internal.Log.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("product successfully deleted"))
	internal.Log.Infof("product %d successfully deleted", productID)

}

func (h *ProductHandler) GetBySupplier(w http.ResponseWriter, req *http.Request) {
	supplierID, _ := strconv.Atoi(req.URL.Query().Get("id"))

	order, err := h.services.GetBySupplier(int32(supplierID))
	if err != nil {
		internal.Log.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	respJ, _ := json.Marshal(order)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJ)
	internal.Log.Infof("user successfully fetched supplier %d products ", supplierID)

}
