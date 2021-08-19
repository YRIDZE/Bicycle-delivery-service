package db_repository

import (
	"database/sql"
	"fmt"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserDBRepository struct {
	db *sql.DB
}

func NewUserDBRepository(db *sql.DB) *UserDBRepository {
	return &UserDBRepository{db: db}
}

func (u UserDBRepository) Create(user *models.User) (int32, error) {
	createUserQuery := fmt.Sprintf("insert into %s (firstname, lastname, email, password) value (?, ?, ?, ?)", usersTable)
	us, err := u.db.Prepare(createUserQuery)
	if err != nil {
		log.Fatal(err)
	}

	user.Password = generatePasswordHash(user.Password)

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

func (u UserDBRepository) GetByID(id int32) (*models.User, error) {
	user := new(models.User)
	query := fmt.Sprintf("select id, firstname, lastname, email, password from %s where id = ?", usersTable)

	err := u.db.QueryRow(query, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserDBRepository) GetByEmail(email string) (*models.User, error) {
	user := new(models.User)
	query := fmt.Sprintf("select id, firstname, lastname, email, password from %s where email = ?", usersTable)

	err := u.db.QueryRow(query, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserDBRepository) GetAll() (*[]models.User, error) {
	var users []models.User
	var user models.User
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

func (u UserDBRepository) Update(user *models.User) error {
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

func generatePasswordHash(password string) string {
	p, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(p)
}
