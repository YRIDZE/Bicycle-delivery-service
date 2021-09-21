package db_repository

import (
	"database/sql"

	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
)

type UserRepositoryMock struct {
	db *sql.DB
}

func NewUserRepositoryMock(db *sql.DB) *UserRepositoryMock {
	return &UserRepositoryMock{db: db}
}

func (u *UserRepositoryMock) Create(user *models.User) (*models.User, error) {
	return user, nil
}

func (u *UserRepositoryMock) GetByID(id int32) (*models.User, error) {
	user := &models.User{
		ID:        id,
		FirstName: "firstname",
		LastName:  "lastname",
		Email:     "email",
		Password:  "$2a$10$px.3fiUQFZPHiik82ARVVeSOgYN4YAudf2mcw9Xeh8Hzrdr7o2hqW",
	}
	return user, nil
}

func (u *UserRepositoryMock) GetByEmail(email string) (*models.User, error) {
	user := &models.User{
		ID:        1,
		FirstName: "firstname",
		LastName:  "lastname",
		Email:     email,
		Password:  "$2a$10$px.3fiUQFZPHiik82ARVVeSOgYN4YAudf2mcw9Xeh8Hzrdr7o2hqW",
	}
	return user, nil
}

func (u *UserRepositoryMock) GetAll() (*[]models.User, error) {
	user := []models.User{
		{
			ID:        1,
			FirstName: "firstname",
			LastName:  "lastname",
			Email:     "email",
			Password:  "$2a$10$px.3fiUQFZPHiik82ARVVeSOgYN4YAudf2mcw9Xeh8Hzrdr7o2hqW",
		},
	}
	return &user, nil
}

func (u *UserRepositoryMock) Update(user *models.User) (*models.User, error) {
	return user, nil
}

func (u *UserRepositoryMock) Delete(id int32) error {
	return nil
}
