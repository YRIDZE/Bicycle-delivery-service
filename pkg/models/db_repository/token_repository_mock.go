package db_repository

import (
	"database/sql"

	"github.com/YRIDZE/Bicycle-delivery-service/helper"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
)

type TokensRepositoryMock struct {
	db *sql.DB
}

func NewTokensRepositoryMock(db *sql.DB) *TokensRepositoryMock {
	return &TokensRepositoryMock{db: db}
}

func (u *TokensRepositoryMock) CreateUid(userID int32, uid models.CachedTokens) error {
	return nil
}

func (u *TokensRepositoryMock) GetUidByID(claims *helper.JwtCustomClaims) (*models.CachedTokens, error) {
	cachedT := &models.CachedTokens{AccessUID: claims.UID}
	return cachedT, nil
}

func (u *TokensRepositoryMock) UpdateUid(userID int32, uid models.CachedTokens) error {
	return nil
}
func (u *TokensRepositoryMock) DeleteUid(userID int32) error {
	return nil
}
