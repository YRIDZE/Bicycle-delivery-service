package services

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
)

type SupplierService struct {
	repo db_repository.SupplierRepositoryI
}

func NewSupplierService(repo db_repository.SupplierRepositoryI) *SupplierService {
	return &SupplierService{repo: repo}
}

func (s SupplierService) Create(supplier *models.SupplierResponse) (*models.SupplierResponse, error) {
	return s.repo.Create(supplier)
}

func (s SupplierService) GetByID(id int) (*models.SupplierResponse, error) {
	return s.repo.GetByID(id)
}

func (s SupplierService) GetByName(name string) (int32, error) {
	return s.repo.GetByName(name)
}

func (s SupplierService) GetAll() (*[]models.SupplierResponse, error) {
	return s.repo.GetAll()
}

func (s SupplierService) Update(supplier *models.SupplierResponse) (*models.SupplierResponse, error) {
	return s.repo.Update(supplier)
}

func (s SupplierService) Delete(id int32) error {
	return s.repo.Delete(id)
}
