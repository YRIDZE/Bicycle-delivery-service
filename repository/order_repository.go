package repository

import "github.com/YRIDZE/Bicycle-delivery-service/model"

type OrderRepositoryI interface {
	Create(order *model.Order) (*model.Order, error)
	Get(id *int32) (*model.Order, error)
	GetAll() (*[]model.Order, error)
	Update(order *model.Order) (*model.Order, error)
	Delete(id int) error
}

type OrderDBRepository struct {
}

func NewOrderDBRepository() *OrderDBRepository {
	return &OrderDBRepository{}
}

func (o OrderDBRepository) Create(order *model.Order) (*model.Order, error) {
	panic("implement me")
}

func (o OrderDBRepository) Get(id *int32) (*model.Order, error) {
	panic("implement me")
}

func (o OrderDBRepository) GetAll() (*[]model.Order, error) {
	panic("implement me")
}

func (o OrderDBRepository) Update(order *model.Order) (*model.Order, error) {
	panic("implement me")
}

func (o OrderDBRepository) Delete(id int) error {
	panic("implement me")
}
