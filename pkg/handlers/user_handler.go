package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/services"
	yolo_log "github.com/YRIDZE/yolo-log"
)

type UserHandler struct {
	logger  *yolo_log.Logger
	service *services.UserService
}

func NewUserHandler(logger *yolo_log.Logger, userRepo db_repository.UserRepositoryI, tokenRepo db_repository.TokensRepositoryI) *UserHandler {
	s := services.NewUserService(&userRepo, &tokenRepo)
	return &UserHandler{logger: logger, service: s}
}

func (h *UserHandler) RegisterRoutes(r *http.ServeMux, appH *AppHandlers) {
	r.HandleFunc("/login", h.Login)
	r.HandleFunc("/refreshTokens", h.Refresh)
	r.Handle("/logout", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.Logout)))

	r.HandleFunc("/createUser", appH.userHandler.Create)
	r.HandleFunc("/getUsers", appH.userHandler.GetAll)
	r.Handle("/getUser", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.GetProfile)))
	r.Handle("/updateUser", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.Update)))
	r.Handle("/deleteUser", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.Delete)))
}

func (h *UserHandler) Create(w http.ResponseWriter, req *http.Request) {
	user := new(models.User)
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}

	u, err := h.service.Create(user)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}
	respJ, _ := json.Marshal(u)

	w.WriteHeader(http.StatusCreated)
	w.Write(respJ)
	h.logger.Infof("user %d successfully created", u.ID)
}

func (h *UserHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	users, err := h.service.GetAll()
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	var resp []models.UserResponse
	for _, u := range *users {
		us := &models.UserResponse{
			ID:        u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
		}
		resp = append(resp, *us)
	}
	respJ, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJ)
	h.logger.Infof("users successfully extracted")
}

func (h *UserHandler) Update(w http.ResponseWriter, req *http.Request) {
	user := new(models.User)
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}

	user.ID = req.Context().Value("user").(*models.User).ID
	u, err := h.service.Update(user)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "Invalid data", http.StatusInternalServerError)
		return
	}
	respJ, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Write(respJ)
	h.logger.Infof("user %d successfully updated", user.ID)
}

func (h *UserHandler) Delete(w http.ResponseWriter, req *http.Request) {
	userID := req.Context().Value("user").(*models.User).ID
	err := h.service.Delete(userID)
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
	resp := &models.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
	respJ, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJ)
	h.logger.Infof("user %d successfully fetched profile", user.ID)
}
