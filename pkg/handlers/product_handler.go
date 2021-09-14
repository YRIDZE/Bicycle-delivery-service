package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/YRIDZE/Bicycle-delivery-service/conf"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/requests"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/services"
)

type ProductHandler struct {
	cfg      *conf.Config
	services *services.ProductService
}

func NewProductHandler(cfg conf.Config, repo db_repository.ProductRepositoryI) *ProductHandler {
	s := services.NewProductService(repo)
	return &ProductHandler{cfg: &cfg, services: s}
}

func (h *ProductHandler) RegisterRoutes(r *http.ServeMux, appH *AppHandlers) {
	r.HandleFunc("/createProduct", h.Create)
	r.HandleFunc("/getProductById", h.GetByID)
	r.HandleFunc("/getProducts", h.GetAll)
	r.HandleFunc("/getProductBySupplier", h.GetBySupplier)
	r.HandleFunc("/updateProduct", h.Update)
	r.HandleFunc("/deleteProduct", h.Delete)
}

func (h *ProductHandler) Create(w http.ResponseWriter, req *http.Request) {
	product := new(requests.ProductRequest)
	if err := json.NewDecoder(req.Body).Decode(&product); err != nil {
		h.cfg.Logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}

	p, err := h.services.Create(product)
	if err != nil {
		h.cfg.Logger.Error(err.Error())
		http.Error(w, "invalid data", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&models.ProductResponse{ID: p.ID, SupplierID: p.SupplierID, Name: p.Name, Image: p.Image, Price: p.Price, Type: p.Type, Ingredients: p.Ingredients})
	h.cfg.Logger.Infof("product %d successfully created", p.ID)
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, req *http.Request) {
	productID, _ := strconv.Atoi(req.URL.Query().Get("id"))

	p, err := h.services.GetByID(productID)
	if err != nil {
		h.cfg.Logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&models.ProductResponse{ID: p.ID, SupplierID: p.SupplierID, Name: p.Name, Image: p.Image, Price: p.Price, Type: p.Type, Ingredients: p.Ingredients})
	h.cfg.Logger.Infof("user successfully fetched product %d", p.ID)
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	p, err := h.services.GetAll()
	if err != nil {
		h.cfg.Logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	var resp []models.ProductResponse
	for _, x := range *p {
		resp = append(
			resp,
			models.ProductResponse{ID: x.ID, SupplierID: x.SupplierID, Name: x.Name, Image: x.Image, Price: x.Price, Type: x.Type, Ingredients: x.Ingredients},
		)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	h.cfg.Logger.Infof("user fetched all products")
}

func (h *ProductHandler) Update(w http.ResponseWriter, req *http.Request) {
	product := new(requests.ProductRequest)
	if err := json.NewDecoder(req.Body).Decode(&product); err != nil {
		h.cfg.Logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}

	p, err := h.services.Update(product)
	if err != nil {
		h.cfg.Logger.Error(err.Error())
		http.Error(w, "invalid data", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&models.ProductResponse{ID: p.ID, SupplierID: p.SupplierID, Name: p.Name, Image: p.Image, Price: p.Price, Type: p.Type, Ingredients: p.Ingredients})
	h.cfg.Logger.Infof("product %d successfully updated", p.ID)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, req *http.Request) {
	productID, _ := strconv.Atoi(req.URL.Query().Get("id"))

	err := h.services.Delete(productID)
	if err != nil {
		h.cfg.Logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("product successfully deleted"))
	h.cfg.Logger.Infof("product %d successfully deleted", productID)
}

func (h *ProductHandler) GetBySupplier(w http.ResponseWriter, req *http.Request) {
	supplierID, _ := strconv.Atoi(req.URL.Query().Get("id"))

	pr, err := h.services.GetBySupplier(int32(supplierID))
	if err != nil {
		h.cfg.Logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	var resp []models.ProductResponse
	for _, x := range *pr {
		resp = append(
			resp,
			models.ProductResponse{ID: x.ID, SupplierID: x.SupplierID, Name: x.Name, Image: x.Image, Price: x.Price, Type: x.Type, Ingredients: x.Ingredients},
		)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	h.cfg.Logger.Infof("user successfully fetched supplier %d products ", supplierID)
}
