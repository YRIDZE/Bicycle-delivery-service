package db_repository

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
)

type OrderRepositoryI interface {
	Create(order *models.Order) (*models.Order, error)
	GetByID(id *int32) (*models.Order, error)
	GetAll() (*[]models.Order, error)
	Update(order *models.Order) error
	Delete(id int) error
}

type OrderDBRepository struct {
}

func NewOrderDBRepository() *OrderDBRepository {
	return &OrderDBRepository{}
}

func (o OrderDBRepository) Create(order *models.Order) (*models.Order, error) {
	panic("implement me")
}

func (o OrderDBRepository) GetByID(id *int32) (*models.Order, error) {
	panic("implement me")
}

func (o OrderDBRepository) GetAll() (*[]models.Order, error) {
	panic("implement me")
}

func (o OrderDBRepository) Update(order *models.Order) error {
	panic("implement me")
}

func (o OrderDBRepository) Delete(id int) error {
	panic("implement me")
}
