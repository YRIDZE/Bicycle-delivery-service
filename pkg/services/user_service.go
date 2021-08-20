package services

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
)

type UserService struct {
	repo db_repository.UserRepositoryI
}

func NewUserService(repo *db_repository.Repository) *UserService {
	return &UserService{repo: repo.UserRepositoryI}
}

func (u UserService) CreateUser(user *models.User) (int32, error) {
	return u.repo.CreateUser(user)
}

func (u UserService) GetUserByEmail(email string) (*models.User, error) {
	return u.repo.GetUserByEmail(email)
}

func (u UserService) GetUserByID(id int32) (*models.User, error) {
	return u.repo.GetUserByID(id)
}

func (u UserService) GetAllUsers() (*[]models.User, error) {
	return u.repo.GetAllUsers()
}

func (u UserService) UpdateUser(user *models.User) error {
	return u.repo.UpdateUser(user)
}

func (u UserService) DeleteUser(id int32) error {
	return u.repo.DeleteUser(id)
}
