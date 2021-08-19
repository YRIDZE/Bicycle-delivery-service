package db_repository

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/model"
)

type ProductRepositoryI interface {
	Create(product *model.Product) (*model.Product, error)
	GetByID(id *int32) (*model.Product, error)
	GetAll() (*[]model.Product, error)
	Update(product *model.Product) error
	Delete(id int) error
}

type ProductDBRepository struct {
}

func NewProductDBRepository() *ProductDBRepository {
	return &ProductDBRepository{}
}

func (p ProductDBRepository) Create(product *model.Product) (*model.Product, error) {
	panic("implement me")
}

func (p ProductDBRepository) GetByID(id *int32) (*model.Product, error) {
	panic("implement me")
}

func (p ProductDBRepository) GetAll() (*[]model.Product, error) {
	panic("implement me")
}

func (p ProductDBRepository) Update(product *model.Product) error {
	panic("implement me")
}

func (p ProductDBRepository) Delete(id int) error {
	panic("implement me")
}
