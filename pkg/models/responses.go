package models

import (
	"github.com/YRIDZE/Bicycle-delivery-service/internal"
	"net/http"
)

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserResponse struct {
	ID        int32  `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

type TokenResponse struct {
	RefreshToken string `json:"refresh_token"`
}

type CachedTokens struct {
	AccessUID  string `json:"access"`
	RefreshUID string `json:"refresh"`
}

func ErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	internal.Log.Error(message)
	http.Error(w, message, statusCode)
}
