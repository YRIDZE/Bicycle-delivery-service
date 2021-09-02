package db_repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
)

type ProductRepositoryI interface {
	Create(product *models.Product) (int, error)
	GetByID(id int) (*models.Product, error)
	GetAll() (*[]models.Product, error)
	Update(product *models.Product) error
	Delete(id int) error
}

type ProductDBRepository struct {
	db *sql.DB
}

func NewProductDBRepository(db *sql.DB) *ProductDBRepository {
	return &ProductDBRepository{db: db}
}

func (p ProductDBRepository) Create(product *models.Product) (int, error) {

	ctx := context.Background()
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	productQuery := fmt.Sprintf(
		"insert into %s (supplier_id, name, price, logo, ingredients) value (?, ?, ?, ?, ?)", ProductsTable,
	)
	ingredientsJson, _ := json.Marshal(product.Ingredients)
	res, err := tx.ExecContext(ctx, productQuery, product.SupplierID, product.Name, product.Price, product.Logo, ingredientsJson)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	typeQuery := fmt.Sprintf("insert into %s (name, product_id) value (?, ?)", ProductTypeTable)
	_, err = tx.ExecContext(ctx, typeQuery, product.Type, lastId)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return int(lastId), nil
}

func (p ProductDBRepository) GetByID(id int) (*models.Product, error) {
	product := new(models.Product)

	query := fmt.Sprintf(
		"select p.id, p.name, p.price, p.ingredients, p.logo, pt.id, pt.name as type from %s as p inner join %s as pt on p.id = pt.product_id and p.id = ?",
		ProductsTable, ProductTypeTable,
	)
	var ingredientsJson string
	err := p.db.QueryRow(query, id).Scan(
		&product.ID, &product.Name, &product.Price, &ingredientsJson, &product.Logo, &product.SupplierID, &product.Type,
	)
	if err != nil {
		return nil, err
	}

	_ = json.Unmarshal([]byte(ingredientsJson), &product.Ingredients)

	return product, nil
}

func (p ProductDBRepository) GetAll() (*[]models.Product, error) {
	var products []models.Product
	var product models.Product
	var ingredientsJson string

	query := fmt.Sprintf(
		"select p.id, p.name, p.price, p.ingredients, p.logo, pt.id, pt.name as type from %s as p inner join %s as pt on p.id = pt.product_id ",
		ProductsTable, ProductTypeTable,
	)
	pr, err := p.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := pr.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&product.ID, &product.Name, &product.Price, &ingredientsJson, &product.Logo, &product.SupplierID, &product.Type,
		)
		if err != nil {
			return nil, err
		}

		_ = json.Unmarshal([]byte(ingredientsJson), &product.Ingredients)
		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &products, nil
}

func (p ProductDBRepository) Update(product *models.Product) error {
	ctx := context.Background()
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	typeQuery := fmt.Sprintf("update %s set name = ? where product_id = ?", ProductTypeTable)
	_, err = tx.ExecContext(ctx, typeQuery, product.Type, product.ID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	ingredientsJson, _ := json.Marshal(product.Ingredients)
	productQuery := fmt.Sprintf("update %s set name = ?, price = ?, ingredients = ?, logo = ? where id = ?", ProductsTable)
	_, err = tx.ExecContext(ctx, productQuery, product.Name, product.Price, ingredientsJson, product.Logo, product.ID)
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

func (p ProductDBRepository) Delete(id int) error {
	query := fmt.Sprintf("delete from %s where id = ?", ProductsTable)
	_, err := p.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
