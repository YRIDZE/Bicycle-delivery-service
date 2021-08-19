package services

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
)

type UserServiceI interface {
	Create(user *models.User) (int32, error)
	GetByEmail(email string) (*models.User, error)
	GetByID(id int32) (*models.User, error)
	GetAll() (*[]models.User, error)
	Update(user *models.User) error
	Delete(id int) error
}

type Authorization interface {
	GenerateToken(userID int32, lifetimeMinutes int, secret string) (string, error)
	ValidateToken(tokenString, secretString string) (int32, error)
	GetTokenFromBearerString(input string) string
}

type Service struct {
	UserServiceI
	Authorization
}

func NewService(repo db_repository.UserRepositoryI) *Service {
	return &Service{
		&UserService{NewUserService(repo)},
		&AuthService{NewAuthService(repo)},
	}
}
