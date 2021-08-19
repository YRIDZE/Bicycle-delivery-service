package db_repository

import (
	"database/sql"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/model"
)

type UserRepositoryI interface {
	Create(user *model.User) (int32, error)
	GetByEmail(email *string) (*model.User, error)
	GetAll() (*[]model.User, error)
	Update(user *model.User) error
	Delete(id int) error
}

type Repository struct {
	UserRepositoryI
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepositoryI: NewUserDBRepository(db),
	}
}
