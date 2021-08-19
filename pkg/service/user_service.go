package service

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/model"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/model/db_repository"
)

type UserService struct {
	repo db_repository.UserRepositoryI
}

func NewUserService(repo db_repository.UserRepositoryI) *UserService {
	return &UserService{repo: repo}
}

func (u UserService) Create(user *model.User) (int32, error) {
	return u.repo.Create(user)
}

func (u UserService) GetByEmail(email *string) (*model.User, error) {
	return u.repo.GetByEmail(email)
}

func (u UserService) GetAll() (*[]model.User, error) {
	return u.repo.GetAll()
}

func (u UserService) Update(user *model.User) error {
	return u.repo.Update(user)
}

func (u UserService) Delete(id int) error {
	return u.repo.Delete(id)
}
