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
	auth := appH.UserHandler.AuthMiddleware
	meth := appH.MethodDispatcher

	r.Handle("/login", meth(Methods{post: http.HandlerFunc(h.Login)}))
	r.Handle("/refresh", meth(Methods{get: http.HandlerFunc(h.Refresh)}))
	r.Handle("/logout", auth(meth(Methods{get: http.HandlerFunc(h.Logout)})))

	r.Handle("/createUser", meth(Methods{get: http.HandlerFunc(h.Create)}))
	r.Handle("/getUser", auth(meth(Methods{get: http.HandlerFunc(h.GetProfile)})))
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

func (h *UserHandler) GetProfile(w http.ResponseWriter, req *http.Request) {

	user := req.Context().Value("user").(*models.User)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&models.UserResponse{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName, Email: user.Email})
	h.logger.Infof("user %d successfully fetched profile", user.ID)
}
