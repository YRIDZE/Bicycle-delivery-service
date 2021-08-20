package handlers

import (
	"encoding/json"
	"github.com/YRIDZE/Bicycle-delivery-service/conf"
	"github.com/YRIDZE/Bicycle-delivery-service/internal"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (h *Handler) Logout(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		delete(internal.WhiteList, req.Context().Value("user").(*models.User).ID)
		err := h.services.DeleteUid(req.Context().Value("user").(*models.User).ID)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "refresh-token",
			Value:    "",
			Path:     "/refresh",
			MaxAge:   0,
			HttpOnly: true,
		})

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Successfully logged out"))

	default:
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
	}

}

func (h *Handler) Refresh(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		c, err := req.Cookie("refresh-token")
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusBadRequest)
			return
		}

		claims, err := h.services.ValidateToken(c.Value, conf.RefreshSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		accessUID, accessString, err := h.services.GenerateToken(claims.ID, conf.AccessLifetimeMinutes, conf.AccessSecret)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		refreshUID, refreshString, err := h.services.GenerateToken(claims.ID, conf.RefreshLifetimeMinutes, conf.RefreshSecret)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		cachedTokens := models.CachedTokens{
			AccessUID:  accessUID,
			RefreshUID: refreshUID,
		}
		err = h.services.UpdateUid(claims.ID, cachedTokens)

		resp := models.LoginResponse{
			AccessToken:  accessString,
			RefreshToken: refreshString,
		}
		respJ, _ := json.Marshal(resp)

		http.SetCookie(w, &http.Cookie{
			Name:     "refresh-token",
			Value:    refreshString,
			Path:     "/refresh",
			HttpOnly: true,
		})

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respJ)

	default:
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) GetProfile(w http.ResponseWriter, req *http.Request) {
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

func (h *Handler) Login(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		r := new(models.LoginRequest)
		if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := h.services.GetUserByEmail(r.Email)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password)); err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		accessUID, accessString, err := h.services.GenerateToken(user.ID, conf.AccessLifetimeMinutes, conf.AccessSecret)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		refreshUID, refreshString, err := h.services.GenerateToken(user.ID, conf.RefreshLifetimeMinutes, conf.RefreshSecret)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		cachedTokens := models.CachedTokens{
			AccessUID:  accessUID,
			RefreshUID: refreshUID,
		}
		err = h.services.AddUid(user.ID, cachedTokens)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		resp := models.LoginResponse{
			AccessToken:  accessString,
			RefreshToken: refreshString,
		}
		respJ, _ := json.Marshal(resp)

		http.SetCookie(w, &http.Cookie{
			Name:     "refresh-token",
			Value:    refreshString,
			Path:     "/refresh",
			HttpOnly: true,
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respJ)

	default:
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
	}
}
