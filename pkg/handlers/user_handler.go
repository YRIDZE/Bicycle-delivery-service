package handlers

import (
	"encoding/json"
	"github.com/YRIDZE/Bicycle-delivery-service/internal"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/services"
	"github.com/gorilla/mux"
	"net/http"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(userRepo db_repository.UserRepositoryI, tokenRepo db_repository.TokensRepositoryI) *UserHandler {
	s := services.NewUserService(&userRepo, &tokenRepo)
	return &UserHandler{service: s}
}

func (h *UserHandler) RegisterRoutes(r *mux.Router, appH *AppHandlers) {
	r.HandleFunc("/login", h.Login).Methods(http.MethodPost)
	r.HandleFunc("/refresh", h.Refresh).Methods(http.MethodPost)
	r.Handle("/logout", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.Logout))).Methods(http.MethodPost)

	r.HandleFunc("/user", appH.userHandler.Create).Methods(http.MethodPost)
	r.HandleFunc("/users", appH.userHandler.GetAll).Methods(http.MethodGet)
	r.Handle("/user", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.GetProfile))).Methods(http.MethodGet)
	r.Handle("/user", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.Update))).Methods(http.MethodPut)
	r.Handle("/user", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.Delete))).Methods(http.MethodDelete)
}

func (h *UserHandler) Create(w http.ResponseWriter, req *http.Request) {
	r := new(models.User)
	if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
		internal.Log.Error(err.Error())
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return
	}

	userID, err := h.service.Create(r)
	if err != nil {
		internal.Log.Error(err.Error())
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User created"))
	internal.Log.Infof("User %d successfully created", userID)

}

func (h *UserHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	users, err := h.service.GetAll()
	if err != nil {
		internal.Log.Error(err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
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
	internal.Log.Infof("Users successfully extracted")
}

func (h *UserHandler) Update(w http.ResponseWriter, req *http.Request) {
	user := new(models.User)
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		internal.Log.Error(err.Error())
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return
	}

	user.ID = req.Context().Value("user").(*models.User).ID
	err := h.service.Update(user)
	if err != nil {
		internal.Log.Error(err.Error())
		http.Error(w, "Invalid data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User successfully updated"))
	internal.Log.Infof("User %d successfully updated", user.ID)
}

func (h *UserHandler) Delete(w http.ResponseWriter, req *http.Request) {
	userID := req.Context().Value("user").(*models.User).ID
	err := h.service.Delete(userID)
	if err != nil {
		internal.Log.Error(err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User successfully deleted"))
	internal.Log.Infof("User %d successfully deleted", userID)

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
	internal.Log.Infof("User %d successfully fetched profile", user.ID)

}
