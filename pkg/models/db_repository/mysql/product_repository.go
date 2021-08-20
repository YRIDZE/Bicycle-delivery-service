package mysql

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
)

type ProductRepositoryI interface {
	CreateProduct(product *models.Product) (int, error)
	GetProductByID(id int) (*models.Product, error)
	GetAllProducts() (*[]models.Product, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(id int) error
}

type ProductDBRepository struct{}

func NewProductDBRepository() *ProductDBRepository {
	return &ProductDBRepository{}
}

func (p ProductDBRepository) CreateProduct(product *models.Product) (int, error) {
	panic("implement me")
}

func (p ProductDBRepository) GetProductByID(id int) (*models.Product, error) {
	panic("implement me")
}

func (p ProductDBRepository) GetAllProducts() (*[]models.Product, error) {
	panic("implement me")
}

func (p ProductDBRepository) UpdateProduct(product *models.Product) error {
	panic("implement me")
}

func (p ProductDBRepository) DeleteProduct(id int) error {
	panic("implement me")
}
