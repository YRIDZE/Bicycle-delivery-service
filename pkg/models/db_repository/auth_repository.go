package db_repository

import (
	"database/sql"
	"fmt"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"log"
)

type TokenCacheRepository struct {
	db *sql.DB
}

func NewTokenCacheRepository(db *sql.DB) *TokenCacheRepository {
	return &TokenCacheRepository{db: db}
}

func (t *TokenCacheRepository) AddUid(userID int32, uid models.CachedTokens) error {
	var exist int8 = 0
	query := fmt.Sprintf("select exists(select user_id from %s where user_id = ?)", CacheTokenTable)
	err := t.db.QueryRow(query, userID).Scan(&exist)
	if err != nil {
		return err
	}
	if exist == 0 {
		query2 := fmt.Sprintf("insert into %s (user_id, access_uid, refresh_uid) value (?, ?, ?)", CacheTokenTable)
		us, err := t.db.Prepare(query2)
		if err != nil {
			log.Fatal(err)
		}

		_, err = us.Exec(userID, uid.AccessUID, uid.RefreshUID)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = t.UpdateUid(userID, uid)
	}
	return nil
}
func (t *TokenCacheRepository) UpdateUid(userID int32, uid models.CachedTokens) error {
	query := fmt.Sprintf("update %s set access_uid = ?, refresh_uid = ? where user_id = ?", CacheTokenTable)
	_, err := t.db.Exec(query, uid.AccessUID, uid.RefreshUID, userID)
	if err != nil {
		return err
	}
	return nil

}

func (t *TokenCacheRepository) DeleteUid(userID int32) error {
	query := fmt.Sprintf("delete from %s where user_id = ?", CacheTokenTable)
	_, err := t.db.Exec(query, userID)
	if err != nil {
		return err
	}
	return nil
}
func (t *TokenCacheRepository) GetUidByID(userID int32) (*models.CachedTokens, error) {
	cachedT := new(models.CachedTokens)
	query := fmt.Sprintf("select access_uid, refresh_uid from %s where user_id = ?", CacheTokenTable)

	err := t.db.QueryRow(query, userID).Scan(&cachedT.AccessUID, &cachedT.RefreshUID)
	if err != nil {
		return nil, err
	}
	return cachedT, nil
}
