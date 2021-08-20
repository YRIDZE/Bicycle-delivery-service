package db_repository

import (
	"database/sql"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
)

type UserRepositoryI interface {
	CreateUser(user *models.User) (int32, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int32) (*models.User, error)
	GetAllUsers() (*[]models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int32) error
}

type AuthorizationI interface {
	AddUid(userID int32, uid models.CachedTokens) error
	GetUidByID(userID int32) (*models.CachedTokens, error)
	UpdateUid(userID int32, uid models.CachedTokens) error
	DeleteUid(userID int32) error
}

type OrderRepositoryI interface {
	CreateOrder(order *models.Order) (int, error)
	GetOrderByID(id int) (*models.Order, error)
	GetAllOrders(userID int32) (*[]models.Order, error)
	UpdateOrder(order *models.Order) error
	DeleteOrder(id int) error
}

type Repository struct {
	UserRepositoryI
	OrderRepositoryI
	AuthorizationI
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepositoryI:  NewUserDBRepository(db),
		OrderRepositoryI: NewOrderDBRepository(db),
		AuthorizationI:   NewTokenCacheRepository(db),
	}
}
