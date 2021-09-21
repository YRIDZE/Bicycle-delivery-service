package services

import (
	"errors"
	"strings"
	"time"

	"github.com/YRIDZE/Bicycle-delivery-service/conf"
	"github.com/YRIDZE/Bicycle-delivery-service/helper"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	yolo_log "github.com/YRIDZE/yolo-log"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type TokenService struct {
	cfg       *conf.ConfigToken
	logger    *yolo_log.Logger
	tokenRepo db_repository.TokensRepositoryI
}

func NewTokenService(cfg *conf.ConfigToken, logger *yolo_log.Logger, tokenRepo db_repository.TokensRepositoryI) *TokenService {
	return &TokenService{
		cfg:       cfg,
		logger:    logger,
		tokenRepo: tokenRepo,
	}
}

func (t *TokenService) GenerateAccessToken(userID int32) (string, string, error) {
	return t.generateToken(userID, t.cfg.AccessLifetimeMinutes, t.cfg.AccessSecret)
}

func (t *TokenService) GenerateRefreshToken(userID int32) (string, string, error) {
	return t.generateToken(userID, t.cfg.RefreshLifetimeMinutes, t.cfg.RefreshSecret)
}

func (t *TokenService) generateToken(userID int32, lifetimeMinutes int, secret string) (string, string, error) {
	uid := uuid.New().String()
	claims := &helper.JwtCustomClaims{
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Minute * time.Duration(lifetimeMinutes)).Unix()},
		ID:             userID,
		UID:            uid,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	return uid, token, err
}

func (t *TokenService) ValidateAccessToken(tokenString string) (*helper.JwtCustomClaims, error) {
	return t.validateToken(tokenString, t.cfg.AccessSecret)
}

func (t *TokenService) ValidateRefreshToken(tokenString string) (*helper.JwtCustomClaims, error) {
	return t.validateToken(tokenString, t.cfg.RefreshSecret)
}

func (t *TokenService) validateToken(tokenString, secretString string) (*helper.JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString, &helper.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretString), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*helper.JwtCustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("failed to parse token claims")
	}

	return claims, nil

}

func (t *TokenService) GetTokenFromBearerString(input string) (string, error) {

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

func (t *TokenService) CreateUid(userID int32, uid models.CachedTokens) error {
	return t.tokenRepo.CreateUid(userID, uid)
}

func (t *TokenService) GetUidByID(userID int32) (*models.CachedTokens, error) {
	return t.tokenRepo.GetUidByID(userID)
}

func (t *TokenService) UpdateUid(userID int32, uid models.CachedTokens) error {
	return t.tokenRepo.UpdateUid(userID, uid)
}

func (t *TokenService) DeleteUid(userID int32) error {
	return t.tokenRepo.DeleteUid(userID)
}
