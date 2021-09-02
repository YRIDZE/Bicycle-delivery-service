package db_repository

import (
	"database/sql"
	"fmt"

	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
)

type SupplierRepositoryI interface {
	Create(supplier *models.Supplier) (int32, error)
	GetByID(id int) (*models.Supplier, error)
	GetAll() (*[]models.Supplier, error)
	Update(supplier *models.Supplier) error
	Delete(id int32) error
}

type SupplierDBRepository struct {
	db *sql.DB
}

func NewSupplierDBRepository(db *sql.DB) *SupplierDBRepository {
	return &SupplierDBRepository{db: db}
}

func (s SupplierDBRepository) Create(supplier *models.Supplier) (int32, error) {
	query := fmt.Sprintf("insert into %s (id, name, logo) value (?, ?, ?)", SuppliersTable)

	_, err := s.db.Exec(query, supplier.ID, supplier.Name, supplier.Logo)
	if err != nil {
		return 0, err
	}

	return supplier.ID, nil
}

func (s SupplierDBRepository) GetByID(id int) (*models.Supplier, error) {
	supplier := new(models.Supplier)

	query := fmt.Sprintf("select id, name, logo from %s where id = ?", SuppliersTable)
	err := s.db.QueryRow(query, id).Scan(&supplier.ID, &supplier.Name, &supplier.Logo)
	if err != nil {
		return nil, err
	}

	return supplier, nil
}

func (s SupplierDBRepository) GetAll() (*[]models.Supplier, error) {
	var suppliers []models.Supplier
	var supplier models.Supplier
	query := fmt.Sprintf("select id, name, logo from %s", SuppliersTable)
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
		err := rows.Scan(&supplier.ID, &supplier.Name, &supplier.Logo)
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

func (s SupplierDBRepository) Update(supplier *models.Supplier) error {
	query := fmt.Sprintf("update %s set name, logo where id = ?", SuppliersTable)
	_, err := s.db.Exec(query, &supplier.Name, &supplier.Logo, supplier.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s SupplierDBRepository) Delete(id int32) error {
	query := fmt.Sprintf("delete from %s where id = ?", SuppliersTable)
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s SupplierDBRepository) SearchByID(id int32) (bool, error) {
	query := fmt.Sprintf("select * from %s where id = ?", SuppliersTable)
	rows, err := s.db.Query(query, id)
	if err != nil {
		return false, err
	}

	if rows.Next() {
		return true, nil
	}
	return false, nil
}
