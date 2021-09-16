package db_repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
)

type SupplierRepositoryI interface {
	Create(supplier *models.Supplier) (*models.Supplier, error)
	GetByID(id int) (*models.Supplier, error)
	GetAll() (*[]models.Supplier, error)
	GetByName(name string) (int32, error)
	Update(supplier *models.Supplier) (*models.Supplier, error)
	DeleteUnnecessary(period int) error
}

type SupplierRepository struct {
	db *sql.DB
}

func NewSupplierRepository(db *sql.DB) *SupplierRepository {
	return &SupplierRepository{db: db}
}

func (s SupplierRepository) Create(supplier *models.Supplier) (*models.Supplier, error) {
	query := fmt.Sprintf("insert into %s (name, type, image, opening, closing) value (?, ?, ?, ?, ?)", SuppliersTable)

	res, err := s.db.Exec(
		query, supplier.Name, supplier.Type, supplier.Image, supplier.WorkHours.Opening, supplier.WorkHours.Closing,
	)
	if err != nil {
		return nil, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	supplier.ID = int32(lastID)

	return supplier, nil
}

func (s SupplierRepository) GetByID(id int) (*models.Supplier, error) {
	var deletedV sql.NullString
	supplier := new(models.Supplier)

	query := fmt.Sprintf("select id, name, type, image, opening, closing, deleted from %s where id = ?", SuppliersTable)
	err := s.db.QueryRow(query, id).Scan(
		&supplier.ID, &supplier.Name, &supplier.Type, &supplier.Image, &supplier.WorkHours.Opening, &supplier.WorkHours.Closing,
		&deletedV,
	)
	if err != nil {
		return nil, err
	}

	supplier.Deleted = ""
	if deletedV.Valid {
		supplier.Deleted = deletedV.String
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
	query := fmt.Sprintf("select id, name, type, image, opening, closing from %s where deleted is null", SuppliersTable)
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
		err := rows.Scan(
			&supplier.ID, &supplier.Name, &supplier.Type, &supplier.Image, &supplier.WorkHours.Opening,
			&supplier.WorkHours.Closing,
		)
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

func (s SupplierRepository) Update(supplier *models.Supplier) (*models.Supplier, error) {
	query := fmt.Sprintf("update %s set name = ?, type = ?, image = ?, opening = ?, closing = ? where id = ?", SuppliersTable)
	_, err := s.db.Exec(
		query, supplier.Name, supplier.Type, supplier.Image, supplier.WorkHours.Opening, supplier.WorkHours.Closing, supplier.ID,
	)
	if err != nil {
		return nil, err
	}
	return supplier, nil
}

func (s SupplierRepository) DeleteUnnecessary(period int) error {
	query := fmt.Sprintf(
		"update %s set deleted = ? where (current_timestamp - updated_at) > ? and deleted is null", SuppliersTable,
	)
	_, err := s.db.Exec(query, (time.Now().UTC()).Format("2006-01-02 15:04:05.999999"), period)
	if err != nil {
		return err
	}
	return nil
}
