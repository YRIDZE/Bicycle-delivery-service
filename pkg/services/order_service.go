package services

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/requests"
)

type OrderService struct {
	repo db_repository.OrderRepositoryI
}

func NewOrderService(repo db_repository.OrderRepositoryI) *OrderService {
	return &OrderService{repo: repo}
}

func (o OrderService) Create(order *requests.OrderRequest) (*models.Order, error) {
	return o.repo.Create(
		&models.Order{
			ID:       order.ID,
			UserID:   order.UserID,
			Address:  order.Address,
			Status:   order.Status,
			Products: order.Products,
		},
	)
}

func (o OrderService) GetByID(id int) (*models.Order, error) {
	return o.repo.GetByID(id)
}

func (o OrderService) GetAll(userID int32) (*[]models.Order, error) {
	return o.repo.GetAll(userID)
}

func (o OrderService) Update(order *requests.OrderRequest) (*models.Order, error) {
	return o.repo.Update(
		&models.Order{
			ID:       order.ID,
			UserID:   order.UserID,
			Address:  order.Address,
			Status:   order.Status,
			Products: order.Products,
		},
	)
}

func (o OrderService) Delete(id int) error {
	return o.repo.Delete(id)
}
