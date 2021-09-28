package services

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/requests"
)

type SupplierService struct {
	repo db_repository.SupplierRepositoryI
}

func NewSupplierService(repo db_repository.SupplierRepositoryI) *SupplierService {
	return &SupplierService{repo: repo}
}

func (s SupplierService) Create(supplier *requests.SupplierRequest) (*models.Supplier, error) {
	return s.repo.Create(
		&models.Supplier{
			ID:    supplier.ID,
			Name:  supplier.Name,
			Image: supplier.Image,
		},
	)
}

func (s SupplierService) GetByID(id int) (*models.Supplier, error) {
	return s.repo.GetByID(id)
}

func (s SupplierService) GetByName(name string) (int32, error) {
	return s.repo.GetByName(name)
}

func (s SupplierService) GetAll() (*[]models.Supplier, error) {
	return s.repo.GetAll()
}
