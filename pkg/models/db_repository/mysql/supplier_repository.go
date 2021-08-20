package mysql

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
)

type SupplierRepositoryI interface {
	CreateSupplier(supplier *models.Supplier) (int, error)
	GetSupplierByID(id int) (*models.Supplier, error)
	GetAllSuppliers() (*[]models.Supplier, error)
	UpdateSupplier(supplier *models.Supplier) error
	DeleteSupplier(id int) error
}

type SupplierDBRepository struct{}

func NewSupplierDBRepository() *SupplierDBRepository {
	return &SupplierDBRepository{}
}

func (s SupplierDBRepository) CreateSupplier(supplier *models.Supplier) (int, error) {
	panic("implement me")
}

func (s SupplierDBRepository) GetSupplierByID(id int) (*models.Supplier, error) {
	panic("implement me")
}

func (s SupplierDBRepository) GetAllSuppliers() (*[]models.Supplier, error) {
	panic("implement me")
}

func (s SupplierDBRepository) UpdateSupplier(supplier *models.Supplier) error {
	panic("implement me")
}

func (s SupplierDBRepository) DeleteSupplier(id int) error {
	panic("implement me")
}
