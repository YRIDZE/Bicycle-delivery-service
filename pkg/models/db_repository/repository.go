package db_repository

import (
	"database/sql"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
)

type UserRepositoryI interface {
	Create(user *models.User) (int32, error)
	GetByEmail(email string) (*models.User, error)
	GetByID(id int32) (*models.User, error)
	GetAll() (*[]models.User, error)
	Update(user *models.User) error
	Delete(id int) error
}

type Authorization interface {
}

type Repository struct {
	UserRepositoryI
	Authorization
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepositoryI: NewUserDBRepository(db),
	}
}
