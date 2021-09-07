package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/YRIDZE/Bicycle-delivery-service/conf"
	"github.com/YRIDZE/Bicycle-delivery-service/internal"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func (h *UserHandler) Logout(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		userID := req.Context().Value("user").(*models.User).ID
		err := h.service.DeleteUid(userID)
		if err != nil {
			internal.Log.Error(err.Error())
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		http.SetCookie(
			w, &http.Cookie{
				Name:     "refresh-token",
				Value:    "",
				Path:     "/refresh",
				MaxAge:   0,
				HttpOnly: true,
			},
		)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("successfully logged out"))
		internal.Log.Infof("user %d successfully logged out", userID)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

}

func (h *UserHandler) Refresh(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		c, err := req.Cookie("refresh-token")
		if err != nil {
			internal.Log.Error(err.Error())
			http.Error(w, "invalid credentials", http.StatusBadRequest)
			return
		}

		claims, err := h.service.ValidateToken(c.Value, conf.RefreshSecret)
		if err != nil {
			internal.Log.Error(err.Error())
			http.Error(w, "something went wrong", http.StatusUnauthorized)
			return
		}

		accessUID, accessString, err := h.service.GenerateToken(claims.ID, conf.AccessLifetimeMinutes, conf.AccessSecret)
		if err != nil {
			internal.Log.Error(err.Error())
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		refreshUID, refreshString, err := h.service.GenerateToken(claims.ID, conf.RefreshLifetimeMinutes, conf.RefreshSecret)
		if err != nil {
			internal.Log.Error(err.Error())
			http.Error(w, "something went wrong", http.StatusInternalServerError)
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

		http.SetCookie(
			w, &http.Cookie{
				Name:     "refresh-token",
				Value:    refreshString,
				Path:     "/refresh",
				HttpOnly: true,
			},
		)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respJ)
		internal.Log.Infof("user %d token successfully refreshed", claims.ID)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

}

func (h *UserHandler) Login(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		r := new(models.LoginRequest)
		if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
			internal.Log.Error(err.Error())
			http.Error(w, "something went wrong", http.StatusBadRequest)
			return
		}

		user, err := h.service.GetByEmail(r.Email)
		if err != nil {
			internal.Log.Error(err.Error())
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password)); err != nil {
			internal.Log.Error(err.Error())
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}

		accessUID, accessString, err := h.service.GenerateToken(user.ID, conf.AccessLifetimeMinutes, conf.AccessSecret)
		if err != nil {
			internal.Log.Error(err.Error())
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		refreshUID, refreshString, err := h.service.GenerateToken(user.ID, conf.RefreshLifetimeMinutes, conf.RefreshSecret)
		if err != nil {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		cachedTokens := models.CachedTokens{
			AccessUID:  accessUID,
			RefreshUID: refreshUID,
		}
		err = h.service.CreateUid(user.ID, cachedTokens)
		if err != nil {
			internal.Log.Error(err.Error())
			http.Error(w, "invalid token", http.StatusInternalServerError)
			return
		}

		resp := models.LoginResponse{
			AccessToken:  accessString,
			RefreshToken: refreshString,
		}
		respJ, _ := json.Marshal(resp)

		http.SetCookie(
			w, &http.Cookie{
				Name:     "refresh-token",
				Value:    refreshString,
				Path:     "/refresh",
				HttpOnly: true,
			},
		)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respJ)
		internal.Log.Infof("user %d successfully logged in", user.ID)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
