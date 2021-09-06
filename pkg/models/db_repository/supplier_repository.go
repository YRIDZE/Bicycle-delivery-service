package db_repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
)

type SupplierRepositoryI interface {
	Create(supplier *models.Supplier) (int32, error)
	GetByID(id int) (*models.Supplier, error)
	GetAll() (*[]models.Supplier, error)
	Update(supplier *models.Supplier) error
	Delete(id int32) error
	GetByName(name string) (int32, error)
}

type SupplierRepository struct {
	db *sql.DB
}

func NewSupplierRepository(db *sql.DB) *SupplierRepository {
	return &SupplierRepository{db: db}
}

func (s SupplierRepository) Create(supplier *models.Supplier) (int32, error) {
	query := fmt.Sprintf("insert into %s (name, image) value (?, ?)", SuppliersTable)

	res, err := s.db.Exec(query, supplier.Name, supplier.Image)
	if err != nil {
		return 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int32(lastId), nil
}

func (s SupplierRepository) GetByID(id int) (*models.Supplier, error) {
	supplier := new(models.Supplier)

	query := fmt.Sprintf("select id, name, image from %s where id = ?", SuppliersTable)
	err := s.db.QueryRow(query, id).Scan(&supplier.ID, &supplier.Name, &supplier.Image)
	if err != nil {
		return nil, err
	}

	return supplier, nil
}

func (s SupplierRepository) GetByName(name string) (int32, error) {
	var id int32
	query := fmt.Sprintf("select id from %s where name = ? and deleted is null", SuppliersTable)
	err := s.db.QueryRow(query, name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s SupplierRepository) GetAll() (*[]models.Supplier, error) {
	var suppliers []models.Supplier
	var supplier models.Supplier
	query := fmt.Sprintf("select id, name, image from %s where deleted != 0", SuppliersTable)
	pr, err := s.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := pr.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&supplier.ID, &supplier.Name, &supplier.Image)
		if err != nil {
			return nil, err
		}
		suppliers = append(suppliers, supplier)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &suppliers, nil
}

func (s SupplierRepository) Update(supplier *models.Supplier) error {
	query := fmt.Sprintf("update %s set name, image where id = ?", SuppliersTable)
	_, err := s.db.Exec(query, &supplier.Name, &supplier.Image, supplier.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s SupplierRepository) Delete(id int32) error {
	query := fmt.Sprintf("update %s set deleted = ? where id = ?", SuppliersTable)
	_, err := s.db.Exec(query, (time.Now().UTC()).Format("2006-01-02 15:04:05.999999"), id)
	if err != nil {
		return err
	}
	return nil
}
