package services

import (
	"errors"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

type JwtCustomClaims struct {
	jwt.StandardClaims
	ID int32 `json:"id"`
}

type AuthService struct {
	repo db_repository.AuthorizationI
}

func NewAuthService(repo *db_repository.Repository) *AuthService {
	return &AuthService{repo: repo.AuthorizationI}
}

func (h *AuthService) GenerateToken(userID int32, lifetimeMinutes int, secret string) (string, error) {
	claims := &JwtCustomClaims{
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Minute * time.Duration(lifetimeMinutes)).Unix()},
		ID:             userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func (h *AuthService) ValidateToken(tokenString, secretString string) (int32, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretString), nil
	})

	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok || !token.Valid {
		return 0, errors.New("failed to parse token claims")
	}

	return claims.ID, nil

}

func (h *AuthService) GetTokenFromBearerString(input string) string {
	if input == "" {
		return ""
	}

	parts := strings.Split(input, "Bearer")
	if len(parts) != 2 {
		return ""
	}

	token := strings.TrimSpace(parts[1])
	if len(token) < 1 {
		return ""
	}

	return token
}
