package db_repository

import (
	"database/sql"
	"fmt"

	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
)

type UserRepositoryI interface {
	Create(user *models.User) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	EmailExist(email string) (int, error)
	GetByID(id int32) (*models.User, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Create(user *models.User) (*models.User, error) {
	createUserQuery := fmt.Sprintf("insert into %s (firstname, lastname, email, password) value (?, ?, ?, ?)", UsersTable)

	res, err := u.db.Exec(createUserQuery, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.ID = int32(lastID)
	return user, nil
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
	query := fmt.Sprintf("select id, firstname, lastname, email, password from %s where email = ? and deleted is null", UsersTable)

	err := u.db.QueryRow(query, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) EmailExist(email string) (int, error) {
	var exist int
	query := fmt.Sprintf("select 1 from %s where email = ? and deleted is null", UsersTable)

	err := u.db.QueryRow(query, email).Scan(&exist)
	if err != nil {
		return 0, err
	}
	return exist, nil
}
