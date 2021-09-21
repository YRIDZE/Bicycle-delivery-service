package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/YRIDZE/Bicycle-delivery-service/conf"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/requests"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/services"
	yolo_log "github.com/YRIDZE/yolo-log"
)

type UserHandler struct {
	cfg          *conf.ConfigToken
	logger       *yolo_log.Logger
	userService  *services.UserService
	tokenService *services.TokenService
}

func NewUserHandler(cfg *conf.ConfigToken, logger *yolo_log.Logger, userRepo db_repository.UserRepositoryI, tokenRepo db_repository.TokensRepositoryI) *UserHandler {
	return &UserHandler{
		cfg:          cfg,
		logger:       logger,
		userService:  services.NewUserService(cfg, logger, userRepo),
		tokenService: services.NewTokenService(cfg, logger, tokenRepo),
	}
}

func (h *UserHandler) RegisterRoutes(r *http.ServeMux, appH *AppHandlers) {
	r.HandleFunc("/login", h.Login)
	r.HandleFunc("/refresh", h.Refresh)
	r.Handle("/logout", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.Logout)))

	r.HandleFunc("/createUser", appH.userHandler.Create)
	r.HandleFunc("/getUsers", appH.userHandler.GetAll)
	r.Handle("/getUser", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.GetProfile)))
	r.Handle("/updateUser", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.Update)))
	r.Handle("/deleteUser", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.Delete)))
}

func (h *UserHandler) Create(w http.ResponseWriter, req *http.Request) {
	user := new(requests.UserRequest)

	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}

	if err := user.Validate(); err != nil {
		h.logger.Error(err.Error())
		requests.ValidationErrorResponse(w, err)
		return
	}

	u, err := h.userService.Create(user)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&models.UserResponse{ID: u.ID, FirstName: u.FirstName, LastName: u.LastName, Email: u.Email})
	h.logger.Infof("user %d successfully created", u.ID)
}

func (h *UserHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	u, err := h.userService.GetAll()
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	var resp []models.UserResponse
	for _, u := range *u {
		resp = append(resp, models.UserResponse{ID: u.ID, FirstName: u.FirstName, LastName: u.LastName, Email: u.Email})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	h.logger.Infof("users successfully extracted")
}

func (h *UserHandler) Update(w http.ResponseWriter, req *http.Request) {
	user := new(requests.UserRequest)

	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}

	if err := user.Validate(); err != nil {
		h.logger.Error(err.Error())
		requests.ValidationErrorResponse(w, err)
		return
	}

	userID := req.Context().Value("user").(*models.User).ID
	u, err := h.userService.Update(user)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "Invalid data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&models.UserResponse{ID: u.ID, FirstName: u.FirstName, LastName: u.LastName, Email: u.Email})
	h.logger.Infof("user %d successfully updated", userID)
}

func (h *UserHandler) Delete(w http.ResponseWriter, req *http.Request) {
	userID := req.Context().Value("user").(*models.User).ID

	err := h.userService.Delete(userID)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user successfully deleted"))
	h.logger.Infof("user %d successfully deleted", userID)
}

func (h *UserHandler) GetProfile(w http.ResponseWriter, req *http.Request) {
	user := req.Context().Value("user").(*models.User)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&models.UserResponse{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName, Email: user.Email})
	h.logger.Infof("user %d successfully fetched profile", user.ID)
}
