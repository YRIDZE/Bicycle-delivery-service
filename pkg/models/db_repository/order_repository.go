package db_repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
)

type OrderRepositoryI interface {
	Create(order *models.Order) (*models.Order, error)
	GetAll(userID int32) (*[]models.Order, error)
	GetOrderProductsByID(id int32) (orderProducts []models.OrderProducts, err error)
}

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (o OrderRepository) Create(order *models.Order) (*models.Order, error) {
	ctx := context.Background()
	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(
		"insert into %s (address, user_id, phone_number, customer_name, customer_lastname, payment_method) values (?, ?, ?, ?, ?, ?)",
		OrdersTable,
	)
	res, err := tx.ExecContext(
		ctx, query, order.Address, order.UserID, order.PhoneNumber, order.CustomerName, order.CustomerLastname, order.PaymentMethod,
	)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	order.ID = int32(lastID)

	query2 := fmt.Sprintf("insert into %s (order_id, product_id, quantity) values (?, ?, ?)", OPTable)
	for i, x := range order.Products {
		_, err := tx.ExecContext(ctx, query2, order.ID, x.ProductID, x.Quantity)
		if err != nil {
			_ = tx.Rollback()
			return nil, err
		}
		order.Products[i].OrderID = order.ID
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o OrderRepository) GetAll(userID int32) (*[]models.Order, error) {
	var orders []models.Order
	query := fmt.Sprintf(
		"select id, user_id, address, phone_number, customer_name, customer_lastname, status, created_at from %s where user_id=? and deleted is null",
		OrdersTable,
	)
	pr, err := o.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := pr.Query(userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.ID, &order.UserID, &order.Address, &order.PhoneNumber, &order.CustomerName, &order.CustomerLastname, &order.Status, &order.CreatedAt)
		if err != nil {
			return nil, err
		}

		orderProducts, err := o.GetOrderProductsByID(order.ID)
		if err != nil {
			return nil, err
		}
		order.Products = orderProducts

		orders = append(orders, order)
	}

	return &orders, nil
}

func (o OrderRepository) GetOrderProductsByID(id int32) (orderProducts []models.OrderProducts, err error) {
	query := fmt.Sprintf("select order_id, product_id, quantity from %s where order_id = ?", OPTable)
	pr, err := o.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := pr.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		orderProduct := new(models.OrderProducts)
		err := rows.Scan(&orderProduct.OrderID, &orderProduct.ProductID, &orderProduct.Quantity)
		if err != nil {
			return nil, err
		}
		orderProducts = append(orderProducts, *orderProduct)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return
}
