package db_repository

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
)

type CartRepositoryI interface {
	Create(cart *models.Cart) (*models.Cart, error)
	GetByID(id *int32) (*models.Cart, error)
	GetAll() (*[]models.Cart, error)
	Update(cart *models.Cart) error
	Delete(id int) error
}

type CartDBRepository struct {
}

func (c CartDBRepository) Create(cart *models.Cart) (*models.Cart, error) {
	panic("implement me")
}

func (c CartDBRepository) GetByID(id *int32) (*models.Cart, error) {
	panic("implement me")
}

func (c CartDBRepository) GetAll() (*[]models.Cart, error) {
	panic("implement me")
}

func (c CartDBRepository) Update(cart *models.Cart) error {
	panic("implement me")
}

func (c CartDBRepository) Delete(id int) error {
	panic("implement me")
}
