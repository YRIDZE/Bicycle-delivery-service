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
	yolo_log "github.com/YRIDZE/yolo-log"
)

type CartHandler struct {
	cfg     *conf.ConfigToken
	logger  *yolo_log.Logger
	service *services.CartService
}

func NewCartHandler(cfg *conf.ConfigToken, logger *yolo_log.Logger, repo db_repository.CartRepositoryI) *CartHandler {
	return &CartHandler{
		cfg:     cfg,
		logger:  logger,
		service: services.NewCartService(repo),
	}
}

func (h *CartHandler) RegisterRoutes(r *http.ServeMux, appH *AppHandlers) {
	auth := appH.UserHandler.AuthMiddleware
	meth := appH.MethodDispatcher

	r.Handle("/createCart", meth(Methods{post: http.HandlerFunc(h.Create)}))
	r.Handle("/createCartProduct", auth(meth(Methods{post: http.HandlerFunc(h.CreateProduct)})))
	r.Handle("/getCartProducts", auth(meth(Methods{get: http.HandlerFunc(h.GetAll)})))
	r.Handle("/updateCart", auth(meth(Methods{put: http.HandlerFunc(h.Update)})))
	r.Handle("/deleteAllCartProducts", auth(meth(Methods{delete: http.HandlerFunc(h.DeleteAll)})))
	r.Handle("/deleteCartProduct", auth(meth(Methods{delete: http.HandlerFunc(h.DeleteProduct)})))
}

func (h *CartHandler) Create(w http.ResponseWriter, req *http.Request) {
	cart := new(requests.CartRequest)

	if err := json.NewDecoder(req.Body).Decode(&cart); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}
	c, err := h.service.Create(cart)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "invalid data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&models.CartResponse{ID: c.ID, UserID: c.UserID})
	h.logger.Infof("cart %d successfully created by user %d", c.ID, c.UserID)
}

func (h *CartHandler) CreateProduct(w http.ResponseWriter, req *http.Request) {
	userID := req.Context().Value("user").(*models.User).ID

	cartRequest := new(requests.CartProductRequest)
	cart := new(models.Cart)

	if err := json.NewDecoder(req.Body).Decode(&cartRequest); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}

	if err := cartRequest.Validate(); err != nil {
		h.logger.Error(err.Error())
		requests.ValidationErrorResponse(w, err)
		return
	}

	exist, err := h.service.GetCart(userID)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}

	switch exist {
	case 1:
		cart, err = h.service.GetCartByUserID(userID)
		if err != nil {
			h.logger.Error(err.Error())
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		for _, x := range cartRequest.Products {
			p := models.CartProducts{CartID: cart.ID, ProductID: x.ProductID, Quantity: x.Quantity, Price: x.Price}
			cart.Products = append(cart.Products, p)
		}
	case 0:
		c := new(requests.CartRequest)
		c.UserID = userID
		cart, err = h.service.Create(c)
		if err != nil {
			h.logger.Error(err.Error())
			http.Error(w, "invalid data", http.StatusUnauthorized)
			return
		}
	}

	c, err := h.service.CreateProduct(cart)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "invalid data", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&models.CartProductResponse{ID: c.ID, UserID: c.UserID, Products: c.Products})
	h.logger.Infof("product %d successfully created by user %d", cartRequest.Products[0].ProductID, c.UserID)
}

func (h *CartHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	user := req.Context().Value("user").(*models.User)
	c, err := h.service.GetAllProductsFromCart(user.ID)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "invalid data", http.StatusInternalServerError)
		return
	}
	var resp []models.CartProductResponse
	for _, x := range *c {
		resp = append(resp, models.CartProductResponse{ID: x.ID, UserID: x.UserID, Products: x.Products})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	h.logger.Infof("user %d successfully fetched cart", user.ID)
}

func (h *CartHandler) Update(w http.ResponseWriter, req *http.Request) {
	cartProduct := new(requests.CartProductRequest)
	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(&cartProduct); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}
	cartProduct.UserID = req.Context().Value("user").(*models.User).ID
	if err := cartProduct.Validate(); err != nil {
		h.logger.Error(err.Error())
		requests.ValidationErrorResponse(w, err)
		return
	}

	p, err := h.service.Update(cartProduct)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "Invalid data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&models.CartProductResponse{ID: p.ID, UserID: p.UserID, Products: p.Products})
	h.logger.Infof("cart %d product successfully updated", p.ID)
}

func (h *CartHandler) DeleteProduct(w http.ResponseWriter, req *http.Request) {
	userID := req.Context().Value("user").(*models.User).ID
	productID, err := strconv.Atoi(req.URL.Query().Get("productId"))
	if err != nil || productID < 1 {
		h.logger.Error(errors.New("invalid product id parameter"))
		http.Error(w, "invalid product id parameter", http.StatusNotFound)
		return
	}

	err = h.service.DeleteProductFromCart(userID, productID)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("cart product successfully deleted"))
	h.logger.Infof("cart product %d successfully deleted", productID)
}

func (h *CartHandler) DeleteAll(w http.ResponseWriter, req *http.Request) {
	userID := req.Context().Value("user").(*models.User).ID

	err := h.service.DeleteAllProductFromCart(userID)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("all cart products successfully deleted"))
	h.logger.Infof("user %d cart products successfully deleted", userID)
}
