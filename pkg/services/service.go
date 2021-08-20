package services

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
)

type UserServiceI interface {
	CreateUser(user *models.User) (int32, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int32) (*models.User, error)
	GetAllUsers() (*[]models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int32) error
}

type OrderServiceI interface {
	CreateOrder(order *models.Order) (int, error)
	GetOrderByID(id int) (*models.Order, error)
	GetAllOrders(userID int32) (*[]models.Order, error)
	UpdateOrder(order *models.Order) error
	DeleteOrder(id int) error
}

type AuthorizationI interface {
	GenerateToken(userID int32, lifetimeMinutes int, secret string) (string, string, error)
	ValidateToken(tokenString, secretString string) (*JwtCustomClaims, error)
	GetTokenFromBearerString(input string) string

	AddUid(userID int32, uid models.CachedTokens) error
	GetUidByID(userID int32) (*models.CachedTokens, error)
	UpdateUid(userID int32, uid models.CachedTokens) error
	DeleteUid(userID int32) error
}

type Service struct {
	UserServiceI
	OrderServiceI
	AuthorizationI
}

func NewService(repo *db_repository.Repository) *Service {
	return &Service{
		UserServiceI:   NewUserService(repo.UserRepositoryI),
		OrderServiceI:  NewOrderService(repo.OrderRepositoryI),
		AuthorizationI: NewAuthService(repo.AuthorizationI),
	}
}
