package services

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/requests"
)

type ProductService struct {
	repo db_repository.ProductRepositoryI
}

func NewProductService(repo db_repository.ProductRepositoryI) *ProductService {
	return &ProductService{repo: repo}
}

func (p ProductService) Create(product *requests.ProductRequest) (*models.Product, error) {
	return p.repo.Create(
		&models.Product{
			ID:          product.ID,
			SupplierID:  product.SupplierID,
			Name:        product.Name,
			Image:       product.Image,
			Price:       product.Price,
			Type:        product.Type,
			Ingredients: product.Ingredients,
		},
	)
}

func (p ProductService) GetByID(id int) (*models.Product, error) {
	return p.repo.GetByID(id)
}

func (p ProductService) GetAll() (*[]models.Product, error) {
	return p.repo.GetAll()
}

func (p ProductService) Update(product *requests.ProductRequest) (*models.Product, error) {
	return p.repo.Update(
		&models.Product{
			ID:          product.ID,
			SupplierID:  product.SupplierID,
			Name:        product.Name,
			Image:       product.Image,
			Price:       product.Price,
			Type:        product.Type,
			Ingredients: product.Ingredients,
		},
	)
}

func (p ProductService) Delete(id int) error {
	return p.repo.Delete(id)
}

func (p ProductService) GetByName(name string) (int32, error) {
	return p.repo.GetByName(name)
}

func (p ProductService) GetBySupplier(id int32) (*[]models.Product, error) {
	return p.repo.GetBySupplier(id)
}
