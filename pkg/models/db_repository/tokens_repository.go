package db_repository

import (
	"database/sql"
	"fmt"

	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
)

type TokensRepositoryI interface {
	CreateUid(userID int32, uid models.CachedTokens) error
	GetUidByID(userID int32) (*models.CachedTokens, error)
	UpdateUid(userID int32, uid models.CachedTokens) error
	DeleteUid(userID int32) error
}

type TokensRepository struct {
	db *sql.DB
}

func NewTokensRepository(db *sql.DB) *TokensRepository {
	return &TokensRepository{db: db}
}

func (u *TokensRepository) CreateUid(userID int32, uid models.CachedTokens) error {
	var exist int8 = 0
	query := fmt.Sprintf("select exists(select user_id from %s where user_id = ?)", CacheTokenTable)
	err := u.db.QueryRow(query, userID).Scan(&exist)
	if err != nil {
		return err
	}

	if exist == 0 {
		query := fmt.Sprintf("insert into %s (user_id, access_uid, refresh_uid) value (?, ?, ?)", CacheTokenTable)
		_, err = u.db.Exec(query, userID, uid.AccessUID, uid.RefreshUID)
		if err != nil {
			return err
		}
	} else {
		err = u.UpdateUid(userID, uid)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *TokensRepository) GetUidByID(userID int32) (*models.CachedTokens, error) {
	cachedT := new(models.CachedTokens)
	query := fmt.Sprintf("select access_uid, refresh_uid from %s where user_id = ?", CacheTokenTable)

	err := u.db.QueryRow(query, userID).Scan(&cachedT.AccessUID, &cachedT.RefreshUID)
	if err != nil {
		return nil, err
	}
	return cachedT, nil
}

func (u *TokensRepository) UpdateUid(userID int32, uid models.CachedTokens) error {
	query := fmt.Sprintf("update %s set access_uid = ?, refresh_uid = ? where user_id = ?", CacheTokenTable)
	_, err := u.db.Exec(query, uid.AccessUID, uid.RefreshUID, userID)
	if err != nil {
		return err
	}
	return nil
}
func (u *TokensRepository) DeleteUid(userID int32) error {
	query := fmt.Sprintf("delete from %s where user_id = ?", CacheTokenTable)
	_, err := u.db.Exec(query, userID)
	if err != nil {
		return err
	}
	return nil
}
