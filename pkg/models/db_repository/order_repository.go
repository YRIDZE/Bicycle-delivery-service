package db_repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/YRIDZE/Bicycle-delivery-service/internal"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"log"
)

type OrderDBRepository struct {
	db *sql.DB
}

func NewOrderDBRepository(db *sql.DB) *OrderDBRepository {
	return &OrderDBRepository{db: db}
}

func (o OrderDBRepository) CreateOrder(order *models.Order) (int, error) {

	ctx := context.Background()
	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	query := fmt.Sprintf("insert into %s (address, user_id) values (?, ?)", internal.OrdersTable)
	res, err := tx.ExecContext(ctx, query, order.Address, order.UserID)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	query2 := fmt.Sprintf("insert into %s (order_id, product_id, count) values (?, ?, ?)", internal.OPTable)
	for _, x := range order.Products {
		_, err := tx.ExecContext(ctx, query2, lastId, x.ProductID, x.Count)
		if err != nil {
			_ = tx.Rollback()
			return 0, err
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return int(lastId), nil
}

func (o OrderDBRepository) GetOrderByID(id int) (*models.Order, error) {
	order := new(models.Order)
	query := fmt.Sprintf("select id, user_id, address, status from %s where id = ?", internal.OrdersTable)
	err := o.db.QueryRow(query, id).Scan(&order.ID, &order.UserID, &order.Address, &order.Status)
	if err != nil {
		return nil, err
	}

	query2 := fmt.Sprintf("select order_id, product_id, count from %s where order_id = ?", internal.OPTable)
	rows, err := o.db.Query(query2, order.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		orderProduct := new(models.OrderProducts)
		err := rows.Scan(&orderProduct.OrderID, &orderProduct.ProductID, &orderProduct.Count)
		if err != nil {
			return nil, err
		}
		order.Products = append(order.Products, *orderProduct)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o OrderDBRepository) GetAllOrders(userID int32) (*[]models.Order, error) {
	var orders []models.Order
	query := fmt.Sprintf("select orders.id, orders.user_id, orders.address, orders.status, order_products.order_id,  "+
		"order_products.product_id, order_products.count from %s join %s on orders.id = order_products.order_id where orders.user_id = ?", internal.OrdersTable, internal.OPTable)
	rows, err := o.db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		or := new(models.Order)
		orP := new(models.OrderProducts)
		err := rows.Scan(&or.ID, &or.UserID, &or.Address, &or.Status, &orP.OrderID, &orP.ProductID, &orP.Count)
		if err != nil {
			return nil, err
		}

		if len(orders) != 0 {
			for i := range orders {
				if orders[i].ID == orP.OrderID {
					orders[i].Products = append(orders[i].Products, *orP)
					break
				}
				if i == len(orders)-1 {
					or.Products = append(or.Products, *orP)
					orders = append(orders, *or)
				}
			}
		} else {
			or.Products = append(or.Products, *orP)
			orders = append(orders, *or)
		}
	}
	return &orders, nil
}

func (o OrderDBRepository) UpdateOrder(order *models.Order) error {
	ctx := context.Background()
	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	query := fmt.Sprintf("update %s set address = ?, status = ? where id = ?", internal.OrdersTable)
	_, err = tx.ExecContext(ctx, query, order.Address, order.Status, order.ID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	for _, x := range order.Products {
		query2 := fmt.Sprintf("update %s set count = ? where order_id = ? and product_id = ?", internal.OPTable)
		_, err = tx.ExecContext(ctx, query2, x.Count, x.OrderID, x.ProductID)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (o OrderDBRepository) DeleteOrder(id int) error {

	ctx := context.Background()
	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	query := fmt.Sprintf("delete from %s where order_id = ?", internal.OPTable)
	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	query2 := fmt.Sprintf("delete from %s where id = ?", internal.OrdersTable)
	_, err = tx.ExecContext(ctx, query2, id)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
