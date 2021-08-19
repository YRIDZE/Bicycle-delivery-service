package db_repository

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
)

type ProductRepositoryI interface {
	Create(product *models.Product) (*models.Product, error)
	GetByID(id *int32) (*models.Product, error)
	GetAll() (*[]models.Product, error)
	Update(product *models.Product) error
	Delete(id int) error
}

type ProductDBRepository struct {
}

func NewProductDBRepository() *ProductDBRepository {
	return &ProductDBRepository{}
}

func (p ProductDBRepository) Create(product *models.Product) (*models.Product, error) {
	panic("implement me")
}

func (p ProductDBRepository) GetByID(id *int32) (*models.Product, error) {
	panic("implement me")
}

func (p ProductDBRepository) GetAll() (*[]models.Product, error) {
	panic("implement me")
}

func (p ProductDBRepository) Update(product *models.Product) error {
	panic("implement me")
}

func (p ProductDBRepository) Delete(id int) error {
	panic("implement me")
}
