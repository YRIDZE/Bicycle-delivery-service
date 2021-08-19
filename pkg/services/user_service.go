package services

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
)

type UserService struct {
	repo db_repository.UserRepositoryI
}

func NewUserService(repo db_repository.UserRepositoryI) *UserService {
	return &UserService{repo: repo}
}

func (u UserService) Create(user *models.User) (int32, error) {
	return u.repo.Create(user)
}

func (u UserService) GetByEmail(email string) (*models.User, error) {
	return u.repo.GetByEmail(email)
}

func (u UserService) GetByID(id int32) (*models.User, error) {
	return u.repo.GetByID(id)
}

func (u UserService) GetAll() (*[]models.User, error) {
	return u.repo.GetAll()
}

func (u UserService) Update(user *models.User) error {
	return u.repo.Update(user)
}

func (u UserService) Delete(id int) error {
	return u.repo.Delete(id)
}
