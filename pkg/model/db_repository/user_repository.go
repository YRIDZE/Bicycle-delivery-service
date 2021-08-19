package db_repository

import (
	"database/sql"
	"fmt"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/model"
	"log"
)

type UserDBRepository struct {
	db *sql.DB
}

func NewUserDBRepository(db *sql.DB) *UserDBRepository {
	return &UserDBRepository{db: db}
}

func (u UserDBRepository) Create(user *model.User) (int32, error) {
	createUserQuery := fmt.Sprintf("insert into %s (firstname, lastname, email, password) value (?, ?, ?, ?)", usersTable)
	us, err := u.db.Prepare(createUserQuery)
	if err != nil {
		log.Fatal(err)
	}
	res, err := us.Exec(user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return int32(lastId), nil
}

func (u UserDBRepository) GetByEmail(email *string) (*model.User, error) {
	user := new(model.User)
	query := fmt.Sprintf("select id, firstname, lastname, email, password from %s where email = ?", usersTable)

	err := u.db.QueryRow(query, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserDBRepository) GetAll() (*[]model.User, error) {
	var users []model.User
	var user model.User
	query := fmt.Sprintf("select id, firstname, lastname, email, password from %s", usersTable)

	rows, err := u.db.Query(query)
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
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (u UserDBRepository) Update(user *model.User) error {
	query := fmt.Sprintf("update %s set firstname = ?, lastname = ?, email = ?, password = ? where id = ?", usersTable)
	_, err := u.db.Exec(query, user.FirstName, user.LastName, user.Email, user.Password, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u UserDBRepository) Delete(id int) error {
	query := fmt.Sprintf("delete from %s where id = ?", usersTable)
	_, err := u.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
