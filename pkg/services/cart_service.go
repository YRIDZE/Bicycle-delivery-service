package services

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/requests"
)

type CartService struct {
	repo db_repository.CartRepositoryI
}

func NewCartService(repo db_repository.CartRepositoryI) *CartService {
	return &CartService{repo: repo}
}

func (c CartService) Create(cart *requests.CartRequest) (*models.Cart, error) {
	return c.repo.Create(
		&models.Cart{
			ID:     cart.ID,
			UserID: cart.UserID,
		},
	)
}

func (c CartService) CreateProduct(cart *models.Cart) (*models.Cart, error) {
	return c.repo.CreateProduct(cart)
}

func (c CartService) GetAllProductsFromCart(id int32) (*[]models.Cart, error) {
	return c.repo.GetAllProductsFromCart(id)
}

func (c CartService) GetCartByUserID(userID int32) (*models.Cart, error) {
	return c.repo.GetCartByUserID(userID)
}
func (c CartService) GetCart(userID int32) (int, error) {
	return c.repo.GetCart(userID)
}

func (c CartService) Update(cart *requests.CartProductRequest) (*models.Cart, error) {
	return c.repo.Update(
		&models.Cart{
			ID:       cart.ID,
			UserID:   cart.UserID,
			Products: cart.Products,
		},
	)
}

func (c CartService) Delete(id int) error {
	return c.repo.Delete(id)
}

func (c CartService) DeleteProductFromCart(userID int32, productID int) error {
	return c.repo.DeleteProductFromCart(userID, productID)
}

func (c CartService) DeleteAllProductFromCart(userID int32) error {
	return c.repo.DeleteAllProductFromCart(userID)
}
