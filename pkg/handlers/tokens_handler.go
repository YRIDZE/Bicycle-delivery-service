package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/requests"
	"golang.org/x/crypto/bcrypt"
)

func (h *UserHandler) Logout(w http.ResponseWriter, req *http.Request) {
	userID := req.Context().Value("user").(*models.User).ID
	err := h.tokenService.DeleteUid(userID)
	if err != nil {
		h.logger.Error(err.Error())
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
	h.logger.Infof("user %d successfully logged out", userID)
}

func (h *UserHandler) Refresh(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("refresh-token")
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "invalid credentials", http.StatusBadRequest)
		return
	}

	claims, err := h.tokenService.ValidateRefreshToken(c.Value)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusUnauthorized)
		return
	}

	accessUID, accessString, err := h.tokenService.GenerateAccessToken(claims.ID)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	refreshUID, refreshString, err := h.tokenService.GenerateRefreshToken(claims.ID)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	cachedTokens := models.CachedTokens{
		AccessUID:  accessUID,
		RefreshUID: refreshUID,
	}
	err = h.tokenService.UpdateUid(claims.ID, cachedTokens)

	resp := models.LoginResponse{
		UserID:       claims.ID,
		AccessToken:  accessString,
		RefreshToken: refreshString,
	}

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
	json.NewEncoder(w).Encode(resp)
	h.logger.Infof("user %d token successfully refreshed", claims.ID)
}

func (h *UserHandler) Login(w http.ResponseWriter, req *http.Request) {
	r := new(requests.LoginRequest)
	if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}

	if err := r.Validate(); err != nil {
		h.logger.Error(err)
		requests.ValidationErrorResponse(w, err)
		return
	}

	user, err := h.userService.GetByEmail(r.Email)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password)); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	accessUID, accessString, err := h.tokenService.GenerateAccessToken(user.ID)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	refreshUID, refreshString, err := h.tokenService.GenerateRefreshToken(user.ID)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	cachedTokens := models.CachedTokens{
		AccessUID:  accessUID,
		RefreshUID: refreshUID,
	}
	err = h.tokenService.CreateUid(user.ID, cachedTokens)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "invalid token", http.StatusInternalServerError)
		return
	}

	resp := models.LoginResponse{
		UserID:       user.ID,
		AccessToken:  accessString,
		RefreshToken: refreshString,
	}

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
	json.NewEncoder(w).Encode(resp)
	h.logger.Infof("user %d successfully logged in", user.ID)
}

func (h *UserHandler) IsValid(w http.ResponseWriter, req *http.Request) {
	var r struct {
		AccessToken string `json:"access_token"`
	}

	if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}

	claims, err := h.tokenService.ValidateAccessToken(r.AccessToken)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, fmt.Sprint("bad token: ", err.Error()), http.StatusUnauthorized)
		return
	}

	if cachedTokens, ok := h.tokenService.GetUidByID(claims); ok != nil || cachedTokens.AccessUID != claims.UID {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		fmt.Println("tyt2")
		return
	}
	w.WriteHeader(http.StatusOK)
}
