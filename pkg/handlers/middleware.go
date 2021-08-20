package handlers

import (
	"context"
	"github.com/YRIDZE/Bicycle-delivery-service/conf"
	"net/http"
)

func (h *Handler) AuthMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		bearerString := req.Header.Get("Authorization")
		tokenString := h.services.GetTokenFromBearerString(bearerString)
		claims, err := h.services.ValidateToken(tokenString, conf.AccessSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if cachedTokens, ok := h.services.GetUidByID(claims.ID); ok != nil || cachedTokens.AccessUID != claims.UID {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		user, err := h.services.GetUserByID(claims.ID)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusBadRequest)
			return
		}

		req = req.WithContext(context.WithValue(req.Context(), "accessToken", tokenString))
		handler.ServeHTTP(w, req.WithContext(context.WithValue(req.Context(), "user", user)))
	})
}