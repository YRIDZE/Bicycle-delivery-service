package services

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/requests"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo  db_repository.UserRepositoryI
	tokenRepo db_repository.TokensRepositoryI
}

func NewUserService(userRepo *db_repository.UserRepositoryI, tokenRepo *db_repository.TokensRepositoryI) *UserService {
	return &UserService{
		userRepo:  *userRepo,
		tokenRepo: *tokenRepo,
	}
}

func (u UserService) Create(user *requests.UserRequest) (*models.User, error) {
	return u.userRepo.Create(
		&models.User{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Password:  generatePasswordHash(user.Password),
		},
	)
}

func (u UserService) GetByEmail(email string) (*models.User, error) {
	return u.userRepo.GetByEmail(email)
}

func (u UserService) GetByID(id int32) (*models.User, error) {
	return u.userRepo.GetByID(id)
}

func (u UserService) GetAll() (*[]models.User, error) {
	return u.userRepo.GetAll()
}

func (u UserService) Update(user *requests.UserRequest) (*models.User, error) {
	return u.userRepo.Create(
		&models.User{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Password:  generatePasswordHash(user.Password),
		},
	)
}

func (u UserService) Delete(id int32) error {
	return u.userRepo.Delete(id)
}

func generatePasswordHash(password string) string {
	p, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(p)
}
