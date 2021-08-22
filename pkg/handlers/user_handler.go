package handlers

import (
	"encoding/json"
	"fmt"
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
	r.HandleFunc("/login", h.Login)
	r.HandleFunc("/refresh", h.Refresh)
	r.Handle("/logout", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.Logout)))

	r.HandleFunc("/userC", appH.userHandler.Create)
	r.HandleFunc("/users", appH.userHandler.GetAll)
	r.HandleFunc("/user/{email}", h.GetByEmail)
	r.Handle("/user", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.GetProfile)))
	r.Handle("/userU", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.Update)))
	r.Handle("/userD", appH.userHandler.AuthMiddleware(http.HandlerFunc(h.Delete)))
}

func (h *UserHandler) Create(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		r := new(models.User)
		if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err := h.service.Create(r)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User created"))

	default:
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) GetByEmail(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		user, err := h.service.GetByEmail(mux.Vars(req)["email"])
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

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

	default:
		http.Error(w, "Only GET is allowed", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		users, err := h.service.GetAll()
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
		respJ, _ := json.Marshal(users)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respJ)

	default:
		http.Error(w, "Only GET is allowed", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) Update(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "PUT":
		user := new(models.User)
		if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user.ID = req.Context().Value("user").(*models.User).ID
		err := h.service.Update(user)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User updated"))

	default:
		http.Error(w, "Only GET is allowed", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) Delete(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "DELETE":
		err := h.service.Delete(req.Context().Value("user").(*models.User).ID)
		fmt.Println()
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User successfully deleted"))

	default:
		http.Error(w, "Only GET is allowed", http.StatusMethodNotAllowed)
	}

}

func (h *UserHandler) GetProfile(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
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

	default:
		http.Error(w, "Only GET is allowed", http.StatusMethodNotAllowed)
	}
}
