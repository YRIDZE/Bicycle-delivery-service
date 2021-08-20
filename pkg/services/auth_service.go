package services

import (
	"errors"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"strings"
	"time"
)

type JwtCustomClaims struct {
	jwt.StandardClaims
	ID  int32  `json:"id"`
	UID string `json:"uid"`
}

type AuthService struct {
	repo db_repository.AuthorizationI
}

func NewAuthService(repo db_repository.AuthorizationI) *AuthService {
	return &AuthService{repo: repo}
}

func (h *AuthService) GenerateToken(userID int32, lifetimeMinutes int, secret string) (string, string, error) {
	uid := uuid.New().String()
	claims := &JwtCustomClaims{
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Minute * time.Duration(lifetimeMinutes)).Unix()},
		ID:             userID,
		UID:            uid,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	return uid, token, err
}

func (h *AuthService) ValidateToken(tokenString, secretString string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretString), nil
	})

	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("failed to parse token claims")
	}

	return claims, nil

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

func (h *AuthService) AddUid(userID int32, uid models.CachedTokens) error {
	return h.repo.AddUid(userID, uid)
}

func (h *AuthService) UpdateUid(userID int32, uid models.CachedTokens) error {
	return h.repo.UpdateUid(userID, uid)
}

func (h *AuthService) DeleteUid(userID int32) error {
	return h.repo.DeleteUid(userID)
}
func (h *AuthService) GetUidByID(userID int32) (*models.CachedTokens, error) {
	return h.repo.GetUidByID(userID)
}
