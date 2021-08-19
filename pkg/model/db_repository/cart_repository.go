package db_repository

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/model"
)

type CartRepositoryI interface {
	Create(cart *model.Cart) (*model.Cart, error)
	GetByID(id *int32) (*model.Cart, error)
	GetAll() (*[]model.Cart, error)
	Update(cart *model.Cart) error
	Delete(id int) error
}

type CartDBRepository struct {
}

func (c CartDBRepository) Create(cart *model.Cart) (*model.Cart, error) {
	panic("implement me")
}

func (c CartDBRepository) GetByID(id *int32) (*model.Cart, error) {
	panic("implement me")
}

func (c CartDBRepository) GetAll() (*[]model.Cart, error) {
	panic("implement me")
}

func (c CartDBRepository) Update(cart *model.Cart) error {
	panic("implement me")
}

func (c CartDBRepository) Delete(id int) error {
	panic("implement me")
}
