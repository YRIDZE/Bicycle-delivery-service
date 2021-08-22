package handlers

import (
	"encoding/json"
	"github.com/YRIDZE/Bicycle-delivery-service/conf"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (h *UserHandler) Logout(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		err := h.service.DeleteUid(req.Context().Value("user").(*models.User).ID)
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

func (h *UserHandler) Refresh(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		c, err := req.Cookie("refresh-token")
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusBadRequest)
			return
		}

		claims, err := h.service.ValidateToken(c.Value, conf.RefreshSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		accessUID, accessString, err := h.service.GenerateToken(claims.ID, conf.AccessLifetimeMinutes, conf.AccessSecret)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		refreshUID, refreshString, err := h.service.GenerateToken(claims.ID, conf.RefreshLifetimeMinutes, conf.RefreshSecret)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		cachedTokens := models.CachedTokens{
			AccessUID:  accessUID,
			RefreshUID: refreshUID,
		}
		err = h.service.UpdateUid(claims.ID, cachedTokens)

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

func (h *UserHandler) Login(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		r := new(models.LoginRequest)
		if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := h.service.GetByEmail(r.Email)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password)); err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		accessUID, accessString, err := h.service.GenerateToken(user.ID, conf.AccessLifetimeMinutes, conf.AccessSecret)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		refreshUID, refreshString, err := h.service.GenerateToken(user.ID, conf.RefreshLifetimeMinutes, conf.RefreshSecret)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		cachedTokens := models.CachedTokens{
			AccessUID:  accessUID,
			RefreshUID: refreshUID,
		}
		err = h.service.CreateUid(user.ID, cachedTokens)
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
