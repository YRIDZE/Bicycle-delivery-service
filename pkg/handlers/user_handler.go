package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handler) CreateUser(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		r := new(models.User)
		if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err := h.services.CreateUser(r)
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

func (h *Handler) GetUserByEmail(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		user, err := h.services.GetUserByEmail(mux.Vars(req)["email"])
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

func (h *Handler) GetAllUsers(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		users, err := h.services.GetAllUsers()
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

func (h *Handler) UpdateUser(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "PUT":
		user := new(models.User)
		if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user.ID = req.Context().Value("user").(*models.User).ID
		err := h.services.UpdateUser(user)
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

func (h *Handler) DeleteUser(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "DELETE":
		err := h.services.DeleteUser(req.Context().Value("user").(*models.User).ID)
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
