package handlers

import (
	"context"
	"fmt"
	"net/http"
)

func (h *UserHandler) AuthMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			bearerString := req.Header.Get("Authorization")
			tokenString, err := h.service.GetTokenFromBearerString(bearerString)
			if err != nil {
				h.cfg.Logger.Error(err.Error())
				http.Error(w, fmt.Sprint("bad token: ", err.Error()), http.StatusUnauthorized)
				return
			}

			claims, err := h.service.ValidateToken(tokenString, h.cfg.AccessSecret)
			if err != nil {
				h.cfg.Logger.Error(err.Error())
				http.Error(w, fmt.Sprint("bad token: ", err.Error()), http.StatusUnauthorized)
				return
			}

			if cachedTokens, ok := h.service.GetUidByID(claims.ID); ok != nil || cachedTokens.AccessUID != claims.UID {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			user, err := h.service.GetByID(claims.ID)
			if err != nil {
				h.cfg.Logger.Error(err.Error())
				http.Error(w, "invalid credentials", http.StatusBadRequest)
				return
			}

			req = req.WithContext(context.WithValue(req.Context(), "accessToken", tokenString))
			handler.ServeHTTP(w, req.WithContext(context.WithValue(req.Context(), "user", user)))
		},
	)
}
