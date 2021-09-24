package handlers

import (
	"net/http"

	"github.com/YRIDZE/Bicycle-delivery-service/conf"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/services"
	log "github.com/YRIDZE/yolo-log"
)

type UserHandlerMock struct {
	cfg          *conf.ConfigToken
	logger       *log.Logger
	userService  *services.UserService
	tokenService *services.TokenService
}

func NewUserHandlerMock(cfg *conf.ConfigToken, logger *log.Logger, userRepo db_repository.UserRepositoryI, tokenRepo db_repository.TokensRepositoryI) *UserHandlerMock {
	return &UserHandlerMock{
		cfg:          cfg,
		logger:       logger,
		userService:  services.NewUserService(cfg, logger, userRepo),
		tokenService: services.NewTokenService(cfg, logger, tokenRepo),
	}
}

func (h *UserHandlerMock) RegisterRoutes(r *http.ServeMux, appH *AppHandlers) {
	r.Handle("/getUser", appH.UserHandler.AuthMiddleware(http.HandlerFunc(h.GetProfile)))
}

func (h *UserHandlerMock) GetProfile(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}
