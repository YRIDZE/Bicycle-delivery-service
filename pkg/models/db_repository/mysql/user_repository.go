package mysql

import (
	"database/sql"
	"fmt"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
)

type UserDBRepository struct {
	db *sql.DB
}

func NewUserDBRepository(db *sql.DB) *UserDBRepository {
	return &UserDBRepository{db: db}
}

func (u UserDBRepository) CreateUser(user *models.User) (int32, error) {
	fmt.Println(user.Password)
	createUserQuery := fmt.Sprintf("insert into %s (firstname, lastname, email, password) value (?, ?, ?, ?)", UsersTable)
	us, err := u.db.Prepare(createUserQuery)
	if err != nil {
		return 0, err
	}

	res, err := us.Exec(user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int32(lastId), nil
}

func (u UserDBRepository) GetUserByID(id int32) (*models.User, error) {
	user := new(models.User)
	query := fmt.Sprintf("select id, firstname, lastname, email, password from %s where id = ?", UsersTable)

	err := u.db.QueryRow(query, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserDBRepository) GetUserByEmail(email string) (*models.User, error) {
	user := new(models.User)
	query := fmt.Sprintf("select id, firstname, lastname, email, password from %s where email = ?", UsersTable)

	err := u.db.QueryRow(query, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserDBRepository) GetAllUsers() (*[]models.User, error) {
	var users []models.User
	var user models.User
	query := fmt.Sprintf("select id, firstname, lastname, email, password from %s", UsersTable)

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

func (u UserDBRepository) UpdateUser(user *models.User) error {
	query := fmt.Sprintf("update %s set firstname = ?, lastname = ?, email = ?, password = ? where id = ?", UsersTable)
	_, err := u.db.Exec(query, user.FirstName, user.LastName, user.Email, user.Password, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u UserDBRepository) DeleteUser(id int32) error {
	query := fmt.Sprintf("delete from %s where id = ?", UsersTable)
	_, err := u.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
