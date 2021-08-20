package services

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
)

type OrderService struct {
	repo db_repository.OrderRepositoryI
}

func NewOrderService(repo *db_repository.Repository) *OrderService {
	return &OrderService{repo: repo.OrderRepositoryI}
}

func (o OrderService) CreateOrder(order *models.Order) (int, error) {
	return o.repo.CreateOrder(order)
}

func (o OrderService) GetOrderByID(id int) (*models.Order, error) {
	return o.repo.GetOrderByID(id)
}

func (o OrderService) GetAllOrders(userID int32) (*[]models.Order, error) {
	return o.repo.GetAllOrders(userID)
}

func (o OrderService) UpdateOrder(order *models.Order) error {
	return o.repo.UpdateOrder(order)
}

func (o OrderService) DeleteOrder(id int) error {
	return o.repo.DeleteOrder(id)
}
