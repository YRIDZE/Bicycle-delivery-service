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

type ProductHandler struct {
	logger   *yolo_log.Logger
	services *services.ProductService
}

func NewProductHandler(logger *yolo_log.Logger, repo db_repository.ProductRepositoryI) *ProductHandler {
	s := services.NewProductService(repo)
	return &ProductHandler{logger: logger, services: s}
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
	product := new(models.Product)
	if err := json.NewDecoder(req.Body).Decode(&product); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}

	p, err := h.services.Create(product)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "invalid data", http.StatusUnauthorized)
		return
	}
	respJ, _ := json.Marshal(p)

	w.WriteHeader(http.StatusCreated)
	w.Write(respJ)
	h.logger.Infof("product %d successfully created", p.ID)
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, req *http.Request) {
	productID, _ := strconv.Atoi(req.URL.Query().Get("id"))

	p, err := h.services.GetByID(productID)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	respJ, _ := json.Marshal(p)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJ)
	h.logger.Infof("user successfully fetched product %d", p.ID)
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	product, err := h.services.GetAll()
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	respJ, _ := json.Marshal(product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJ)
	h.logger.Infof("user fetched all products")
}

func (h *ProductHandler) Update(w http.ResponseWriter, req *http.Request) {
	product := new(models.Product)
	if err := json.NewDecoder(req.Body).Decode(&product); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}

	p, err := h.services.Update(product)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "invalid data", http.StatusUnauthorized)
		return
	}
	respJ, _ := json.Marshal(p)

	w.WriteHeader(http.StatusOK)
	w.Write(respJ)
	h.logger.Infof("product %d successfully updated", product.ID)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, req *http.Request) {
	productID, _ := strconv.Atoi(req.URL.Query().Get("id"))

	err := h.services.Delete(productID)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("product successfully deleted"))
	h.logger.Infof("product %d successfully deleted", productID)
}

func (h *ProductHandler) GetBySupplier(w http.ResponseWriter, req *http.Request) {
	supplierID, _ := strconv.Atoi(req.URL.Query().Get("id"))

	order, err := h.services.GetBySupplier(int32(supplierID))
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	respJ, _ := json.Marshal(order)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJ)
	h.logger.Infof("user successfully fetched supplier %d products ", supplierID)
}
