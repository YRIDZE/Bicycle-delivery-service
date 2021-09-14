package services

import (
	"errors"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
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

func (u *UserService) GenerateToken(userID int32, lifetimeMinutes int, secret string) (string, string, error) {
	uid := uuid.New().String()
	claims := &JwtCustomClaims{
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Minute * time.Duration(lifetimeMinutes)).Unix()},
		ID:             userID,
		UID:            uid,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	return uid, token, err
}

func (u *UserService) ValidateToken(tokenString, secretString string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretString), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("failed to parse token claims")
	}

	return claims, nil

}

func (u *UserService) GetTokenFromBearerString(input string) (string, error) {

	if input == "" {
		return "", errors.New("no token received")
	}

	parts := strings.Split(input, "Bearer")
	if len(parts) != 2 {
		return "", errors.New("bearer token not in proper format")
	}

	token := strings.TrimSpace(parts[1])
	if len(token) < 1 {
		return "", errors.New("token string is empty")
	}

	return token, nil
}

func (u *UserService) CreateUid(userID int32, uid models.CachedTokens) error {
	return u.tokenRepo.CreateUid(userID, uid)
}

func (u *UserService) UpdateUid(userID int32, uid models.CachedTokens) error {
	return u.tokenRepo.UpdateUid(userID, uid)
}

func (u *UserService) DeleteUid(userID int32) error {
	return u.tokenRepo.DeleteUid(userID)
}
func (u *UserService) GetUidByID(userID int32) (*models.CachedTokens, error) {
	return u.tokenRepo.GetUidByID(userID)
}
