package db_repository

import (
	"database/sql"

	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
)

type SupplierRepositoryMock struct {
	db *sql.DB
}

func NewSupplierRepositoryMock(db *sql.DB) *SupplierRepositoryMock {
	return &SupplierRepositoryMock{db: db}
}

func (s SupplierRepositoryMock) Create(supplier *models.Supplier) (*models.Supplier, error) {
	return supplier, nil
}

func (s SupplierRepositoryMock) GetByID(id int) (*models.Supplier, error) {
	supplier := &models.Supplier{
		ID:    int32(id),
		Name:  "name",
		Type:  "type",
		Image: "image",
		WorkHours: models.WorkingHours{
			Opening: "open",
			Closing: "close",
		},
	}
	return supplier, nil
}

func (s SupplierRepositoryMock) GetByName(name string) (int32, error) {
	return 1, nil
}

func (s SupplierRepositoryMock) GetAll() (*[]models.Supplier, error) {
	suppliers := &[]models.Supplier{
		{
			ID:    1,
			Name:  "name",
			Type:  "type",
			Image: "image",
			WorkHours: models.WorkingHours{
				Opening: "open",
				Closing: "close",
			},
		},
	}

	return suppliers, nil
}

func (s SupplierRepositoryMock) Update(supplier *models.Supplier) (*models.Supplier, error) {
	return supplier, nil
}

func (s SupplierRepositoryMock) DeleteUnnecessary(period int) error {
	return nil
}
