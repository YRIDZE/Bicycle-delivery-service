package repository

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/model"
)

type SupplierRepositoryI interface {
	Create(supplier *model.Supplier) (*model.Supplier, error)
	Get(id *int32) (*model.Supplier, error)
	GetAll() (*[]model.Supplier, error)
	Update(supplier *model.Supplier) (*model.Supplier, error)
	Delete(id int) error
}

type SupplierDBRepository struct {
}

func NewSupplierDBRepository() *SupplierDBRepository {
	return &SupplierDBRepository{}
}

func (s SupplierDBRepository) Create(supplier *model.Supplier) (*model.Supplier, error) {
	panic("implement me")
}

func (s SupplierDBRepository) Get(id *int32) (*model.Supplier, error) {
	panic("implement me")
}

func (s SupplierDBRepository) GetAll() (*[]model.Supplier, error) {
	panic("implement me")
}

func (s SupplierDBRepository) Update(supplier *model.Supplier) (*model.Supplier, error) {
	panic("implement me")
}

func (s SupplierDBRepository) Delete(id int) error {
	panic("implement me")
}
