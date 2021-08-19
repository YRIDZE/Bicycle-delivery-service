package db_repository

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
)

type SupplierRepositoryI interface {
	Create(supplier *models.Supplier) (*models.Supplier, error)
	GetByID(id *int32) (*models.Supplier, error)
	GetAll() (*[]models.Supplier, error)
	Update(supplier *models.Supplier) error
	Delete(id int) error
}

type SupplierDBRepository struct {
}

func NewSupplierDBRepository() *SupplierDBRepository {
	return &SupplierDBRepository{}
}

func (s SupplierDBRepository) Create(supplier *models.Supplier) (*models.Supplier, error) {
	panic("implement me")
}

func (s SupplierDBRepository) GetByID(id *int32) (*models.Supplier, error) {
	panic("implement me")
}

func (s SupplierDBRepository) GetAll() (*[]models.Supplier, error) {
	panic("implement me")
}

func (s SupplierDBRepository) Update(supplier *models.Supplier) error {
	panic("implement me")
}

func (s SupplierDBRepository) Delete(id int) error {
	panic("implement me")
}
