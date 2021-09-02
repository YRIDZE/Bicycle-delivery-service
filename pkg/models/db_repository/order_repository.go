package db_repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
)

type OrderRepositoryI interface {
	Create(order *models.Order) (int, error)
	GetByID(id int) (*models.Order, error)
	GetAll(userID int32) (*[]models.Order, error)
	GetOrderProductsByID(id int32) (orderProducts []models.OrderProducts, err error)
	Update(order *models.Order) error
	Delete(id int) error
}

type OrderDBRepository struct {
	db *sql.DB
}

func NewOrderDBRepository(db *sql.DB) *OrderDBRepository {
	return &OrderDBRepository{db: db}
}

func (o OrderDBRepository) Create(order *models.Order) (int, error) {
	ctx := context.Background()
	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("insert into %s (address, user_id) values (?, ?)", OrdersTable)
	res, err := tx.ExecContext(ctx, query, order.Address, order.UserID)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	query2 := fmt.Sprintf("insert into %s (order_id, product_id, quantity) values (?, ?, ?)", OPTable)
	for _, x := range order.Products {
		_, err := tx.ExecContext(ctx, query2, lastId, x.ProductID, x.Quantity)
		if err != nil {
			_ = tx.Rollback()
			return 0, err
		}
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return int(lastId), nil
}

func (o OrderDBRepository) GetByID(id int) (*models.Order, error) {
	order := new(models.Order)
	query := fmt.Sprintf("select id, user_id, address, status from %s where id = ?", OrdersTable)
	err := o.db.QueryRow(query, id).Scan(&order.ID, &order.UserID, &order.Address, &order.Status)
	if err != nil {
		return nil, err
	}

	query2 := fmt.Sprintf("select order_id, product_id, quantity from %s where order_id = ?", OPTable)
	pr, err := o.db.Prepare(query2)
	if err != nil {
		return nil, err
	}

	rows, err := pr.Query(order.ID)
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
		order.Products = append(order.Products, *orderProduct)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return order, nil
}

func (o OrderDBRepository) GetAll(userID int32) (*[]models.Order, error) {
	var orders []models.Order
	query := fmt.Sprintf("select id, user_id, address, status from %s where user_id=?", OrdersTable)
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
		err := rows.Scan(&order.ID, &order.UserID, &order.Address, &order.Status)
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
	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		return nil, errors.New("sql: no rows in result set")
	}

	return &orders, nil
}

func (o OrderDBRepository) GetOrderProductsByID(id int32) (orderProducts []models.OrderProducts, err error) {
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

func (o OrderDBRepository) Update(order *models.Order) error {
	ctx := context.Background()
	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("update %s set address = ?, status = ? where id = ?", OrdersTable)
	_, err = tx.ExecContext(ctx, query, order.Address, order.Status, order.ID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	for _, x := range order.Products {
		query := fmt.Sprintf("update %s set quantity = ? where order_id = ? and product_id = ?", OPTable)
		_, err = tx.ExecContext(ctx, query, x.Quantity, x.OrderID, x.ProductID)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (o OrderDBRepository) Delete(id int) error {

	ctx := context.Background()
	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("delete from %s where order_id = ?", OPTable)
	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	query2 := fmt.Sprintf("delete from %s where id = ?", OrdersTable)
	_, err = tx.ExecContext(ctx, query2, id)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
