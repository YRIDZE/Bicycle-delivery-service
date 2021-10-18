package db_repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
)

type CartRepositoryI interface {
	Create(cart *models.Cart) (*models.Cart, error)
	CreateProduct(cart *models.Cart) (*models.Cart, error)
	GetCartByUserID(userID int32) (*models.Cart, error)
	GetCart(userID int32) (int, error)
	GetAllProductsFromCart(userID int32) (*[]models.Cart, error)
	GetCartProductsByID(id int) (cartProducts []models.CartProducts, err error)
	Update(cart *models.Cart) (*models.Cart, error)
	DeleteProductFromCart(userID int32, productID int) error
	DeleteAllProductFromCart(userID int32) error
}

type CartRepository struct {
	db *sql.DB
}

func (c CartRepository) GetCart(userID int32) (int, error) {
	var exist int
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE user_id = ? and deleted is null)", CartTable)
	err := c.db.QueryRow(query, userID).Scan(&exist)
	if err != nil {
		return 0, err
	}

	return exist, nil
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (c CartRepository) CreateProduct(cart *models.Cart) (*models.Cart, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var exist int8 = 0

	existQuery := fmt.Sprintf("select exists(select 1 from %s where product_id = ?)", CartProductsTable)
	query := fmt.Sprintf("insert into %s (cart_id, product_id, quantity, price) value (?, ?, ?, ?)", CartProductsTable)

	for _, x := range cart.Products {

		err := c.db.QueryRow(existQuery, x.ProductID).Scan(&exist)
		if err != nil {
			return nil, err
		}
		if exist == 0 {
			_, err = c.db.ExecContext(ctx, query, x.CartID, x.ProductID, x.Quantity, x.Price)
			if err != nil {
				return nil, err
			}
		} else {
			cart, err = c.Update(cart)
			if err != nil {
				return nil, err
			}
		}
	}
	return cart, nil
}

func (c CartRepository) Create(cart *models.Cart) (*models.Cart, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := fmt.Sprintf("insert into %s (user_id) value (?)", CartTable)
	res, err := c.db.ExecContext(ctx, query, cart.UserID)
	if err != nil {
		return nil, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	cart.ID = int(lastID)

	return cart, nil
}

func (c CartRepository) GetCartByUserID(userID int32) (*models.Cart, error) {
	cart := new(models.Cart)
	query := fmt.Sprintf("select id, user_id from %s where user_id = ? and deleted is null", CartTable)
	err := c.db.QueryRow(query, userID).Scan(&cart.ID, &cart.UserID)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (c CartRepository) GetAllProductsFromCart(userID int32) (*[]models.Cart, error) {
	var carts []models.Cart
	query := fmt.Sprintf("select id, user_id from %s where user_id=? and deleted is null", CartTable)
	pr, err := c.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := pr.Query(userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cart models.Cart
		err := rows.Scan(&cart.ID, &cart.UserID)
		if err != nil {
			return nil, err
		}

		orderProducts, err := c.GetCartProductsByID(cart.ID)
		if err != nil {
			return nil, err
		}
		cart.Products = orderProducts

		carts = append(carts, cart)
	}

	return &carts, nil
}

func (c CartRepository) GetCartProductsByID(cartID int) (cartProducts []models.CartProducts, err error) {
	query := fmt.Sprintf("select cart_id, product_id, quantity, price from %s where cart_id = ?", CartProductsTable)
	pr, err := c.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := pr.Query(cartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		cartProduct := new(models.CartProducts)
		err := rows.Scan(&cartProduct.CartID, &cartProduct.ProductID, &cartProduct.Quantity, &cartProduct.Price)
		if err != nil {
			return nil, err
		}
		cartProducts = append(cartProducts, *cartProduct)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return
}

func (c CartRepository) Update(cart *models.Cart) (*models.Cart, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := fmt.Sprintf("update %s set quantity = ? where cart_id = ? and product_id = ?", CartProductsTable)
	for _, x := range cart.Products {
		_, err := c.db.ExecContext(ctx, query, x.Quantity, x.CartID, x.ProductID)
		if err != nil {
			return nil, err
		}
	}

	return cart, nil
}

func (c CartRepository) DeleteProductFromCart(userID int32, productID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	cart, err := c.GetCartByUserID(userID)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("delete from %s where cart_id = ? and product_id = ?", CartProductsTable)
	_, err = c.db.ExecContext(ctx, query, cart.ID, productID)
	if err != nil {
		return err
	}

	return nil
}

func (c CartRepository) DeleteAllProductFromCart(userID int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	cart, err := c.GetCartByUserID(userID)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("delete from %s where cart_id = ?", CartProductsTable)
	_, err = c.db.ExecContext(ctx, query, cart.ID)
	if err != nil {
		return err
	}

	return nil
}
