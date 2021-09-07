package db_repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
)

type UserRepositoryI interface {
	Create(user *models.User) (int32, error)
	GetByEmail(email string) (*models.User, error)
	GetByID(id int32) (*models.User, error)
	GetAll() (*[]models.User, error)
	Update(user *models.User) error
	Delete(id int32) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Create(user *models.User) (int32, error) {
	createUserQuery := fmt.Sprintf("insert into %s (firstname, lastname, email, password) value (?, ?, ?, ?)", UsersTable)

	res, err := u.db.Exec(createUserQuery, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int32(lastId), nil
}

func (u *UserRepository) GetByID(id int32) (*models.User, error) {
	user := new(models.User)
	query := fmt.Sprintf("select id, firstname, lastname, email, password from %s where id = ? and deleted is null", UsersTable)

	err := u.db.QueryRow(query, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) GetByEmail(email string) (*models.User, error) {
	user := new(models.User)
	query := fmt.Sprintf(
		"select id, firstname, lastname, email, password from %s where email = ? and deleted is null", UsersTable,
	)

	err := u.db.QueryRow(query, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) GetAll() (*[]models.User, error) {
	var users []models.User
	var user models.User
	query := fmt.Sprintf("select id, firstname, lastname, email, password from %s where deleted is null", UsersTable)
	pr, err := u.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := pr.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &users, nil
}

func (u *UserRepository) Update(user *models.User) error {
	query := fmt.Sprintf("update %s set firstname = ?, lastname = ?, email = ?, password = ? where id = ?", UsersTable)
	_, err := u.db.Exec(query, user.FirstName, user.LastName, user.Email, user.Password, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) Delete(id int32) error {
	query := fmt.Sprintf("update %s set deleted = ? where id = ?", UsersTable)
	_, err := u.db.Exec(query, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}
