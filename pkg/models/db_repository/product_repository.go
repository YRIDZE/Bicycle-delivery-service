package db_repository

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
)

type ProductRepositoryI interface {
	Create(product *models.Product) (*models.Product, error)
	GetByID(id int) (*models.Product, error)
	GetAll() (*[]models.Product, error)
	Update(product *models.Product) (*models.Product, error)
	GetByName(name string) (int32, error)
	GetBySupplier(id int32) (*[]models.Product, error)
}

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p ProductRepository) Create(product *models.Product) (*models.Product, error) {

	productQuery := fmt.Sprintf(
		"insert into %s (supplier_id, name, price, type, image, ingredients) value (?, ?, ?, ?, ?, ?)", ProductsTable,
	)
	ingredientsJson, _ := json.Marshal(product.Ingredients)
	res, err := p.db.Exec(
		productQuery, product.SupplierID, product.Name, product.Price, product.Type, product.Image, ingredientsJson,
	)
	if err != nil {
		return nil, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	product.ID = int(lastID)

	return product, nil
}

func (p ProductRepository) GetByID(id int) (*models.Product, error) {
	product := new(models.Product)

	query := fmt.Sprintf(
		"select id, name, price, type, ingredients, image, supplier_id from %s where id = ?",
		ProductsTable,
	)
	var ingredientsJson string
	err := p.db.QueryRow(query, id).Scan(
		&product.ID, &product.Name, &product.Price, &product.Type, &ingredientsJson, &product.Image, &product.SupplierID,
	)
	if err != nil {
		return nil, err
	}

	_ = json.Unmarshal([]byte(ingredientsJson), &product.Ingredients)

	return product, nil
}

func (p ProductRepository) GetBySupplier(supplierID int32) (*[]models.Product, error) {
	var products []models.Product
	var product models.Product
	var ingredientsJson string

	query := fmt.Sprintf(
		"select id, supplier_id, name, price, type, ingredients, image from %s where supplier_id = ? and deleted is null ",
		ProductsTable,
	)
	pr, err := p.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := pr.Query(supplierID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&product.ID, &product.SupplierID, &product.Name, &product.Price, &product.Type, &ingredientsJson, &product.Image,
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

func (p ProductRepository) GetByName(name string) (int32, error) {
	var id int32
	query := fmt.Sprintf("select id from %s where name = ? and deleted is null", ProductsTable)
	err := p.db.QueryRow(query, name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (p ProductRepository) GetAll() (*[]models.Product, error) {
	var products []models.Product
	var product models.Product
	var ingredientsJson string

	query := fmt.Sprintf(
		"select id, supplier_id, name, price, type, ingredients, image from %s where deleted is null ", ProductsTable,
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
			&product.ID, &product.SupplierID, &product.Name, &product.Price, &product.Type, &ingredientsJson, &product.Image,
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

func (p ProductRepository) Update(product *models.Product) (*models.Product, error) {
	productQuery := fmt.Sprintf("update %s set price = ? where id = ?", ProductsTable)
	_, err := p.db.Exec(productQuery, product.Price, product.ID)
	if err != nil {
		return nil, err
	}

	return product, nil
}
