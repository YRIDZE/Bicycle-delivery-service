package services

import (
	"errors"
	"strings"
	"time"

	"github.com/YRIDZE/Bicycle-delivery-service/conf"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	yolo_log "github.com/YRIDZE/yolo-log"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type JwtCustomClaims struct {
	jwt.StandardClaims
	ID  int32  `json:"id"`
	UID string `json:"uid"`
}

type TokenService struct {
	cfg       *conf.ConfigToken
	logger    *yolo_log.Logger
	tokenRepo db_repository.TokensRepositoryI
}

func NewTokenService(cfg *conf.ConfigToken, logger *yolo_log.Logger, tokenRepo *db_repository.TokensRepositoryI) *TokenService {
	return &TokenService{
		cfg:       cfg,
		logger:    logger,
		tokenRepo: *tokenRepo,
	}
}

func (s *TokenService) GenerateAccessToken(userID int32) (string, string, error) {
	return s.generateToken(userID, s.cfg.AccessLifetimeMinutes, s.cfg.AccessSecret)
}

func (s *TokenService) GenerateRefreshToken(userID int32) (string, string, error) {
	return s.generateToken(userID, s.cfg.RefreshLifetimeMinutes, s.cfg.RefreshSecret)
}

func (u *TokenService) generateToken(userID int32, lifetimeMinutes int, secret string) (string, string, error) {
	uid := uuid.New().String()
	claims := &JwtCustomClaims{
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Minute * time.Duration(lifetimeMinutes)).Unix()},
		ID:             userID,
		UID:            uid,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	return uid, token, err
}

func (s *TokenService) ValidateAccessToken(tokenString string) (*JwtCustomClaims, error) {
	return s.validateToken(tokenString, s.cfg.AccessSecret)
}

func (s *TokenService) ValidateRefreshToken(tokenString string) (*JwtCustomClaims, error) {
	return s.validateToken(tokenString, s.cfg.RefreshSecret)
}

func (u *TokenService) validateToken(tokenString, secretString string) (*JwtCustomClaims, error) {
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

func (u *TokenService) GetTokenFromBearerString(input string) (string, error) {

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

func (u *TokenService) CreateUid(userID int32, uid models.CachedTokens) error {
	return u.tokenRepo.CreateUid(userID, uid)
}

func (u *TokenService) GetUidByID(userID int32) (*models.CachedTokens, error) {
	return u.tokenRepo.GetUidByID(userID)
}

func (u *TokenService) UpdateUid(userID int32, uid models.CachedTokens) error {
	return u.tokenRepo.UpdateUid(userID, uid)
}

func (u *TokenService) DeleteUid(userID int32) error {
	return u.tokenRepo.DeleteUid(userID)
}
