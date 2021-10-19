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

type ProductHandler struct {
	logger   *yolo_log.Logger
	services *services.ProductService
}

func NewProductHandler(logger *yolo_log.Logger, repo db_repository.ProductRepositoryI) *ProductHandler {
	return &ProductHandler{
		logger:   logger,
		services: services.NewProductService(repo),
	}
}

func (h *ProductHandler) RegisterRoutes(r *http.ServeMux, appH *AppHandlers) {
	meth := appH.MethodDispatcher

	r.Handle("/createProduct", meth(Methods{post: http.HandlerFunc(h.Create)}))
	r.Handle("/getProducts", meth(Methods{get: http.HandlerFunc(h.GetAll)}))
	r.Handle("/getProductTypes", meth(Methods{get: http.HandlerFunc(h.GetTypes)}))
	r.Handle("/getProductTypesBySupp", meth(Methods{get: http.HandlerFunc(h.GetTypesBySupplier)}))
}

func (h *ProductHandler) Create(w http.ResponseWriter, req *http.Request) {
	product := new(requests.ProductRequest)
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

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&models.ProductResponse{ID: p.ID, SupplierID: p.SupplierID, Name: p.Name, Image: p.Image, Price: p.Price, Type: p.Type, Ingredients: p.Ingredients})
	h.logger.Infof("product %d successfully created", p.ID)
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	p, err := h.services.GetAll()
	if err != nil {
		h.logger.Error(err.Error())
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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	h.logger.Infof("fetched all products")
}

func (h *ProductHandler) GetTypes(w http.ResponseWriter, req *http.Request) {
	productTypes, err := h.services.GetTypes()
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productTypes)
	h.logger.Infof("fetched product types")
}

func (h *ProductHandler) GetTypesBySupplier(w http.ResponseWriter, req *http.Request) {
	supplierID, _ := strconv.Atoi(req.URL.Query().Get("id"))
	productTypes, err := h.services.GetTypesBySupplier(int32(supplierID))
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productTypes)
	h.logger.Infof("user successfully fetched supplier %d products ", supplierID)
}
