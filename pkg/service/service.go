package service

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/model"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/model/db_repository"
)

type UserServiceI interface {
	Create(user *model.User) (int32, error)
	GetByEmail(email *string) (*model.User, error)
	GetAll() (*[]model.User, error)
	Update(user *model.User) error
	Delete(id int) error
}

type Service struct {
	UserServiceI
}

func NewService(repo db_repository.UserRepositoryI) *Service {
	return &Service{
		&UserService{NewUserService(repo)},
	}
}
