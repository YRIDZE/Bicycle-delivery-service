package db_repository

import (
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/tests/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	mock     sqlmock.Sqlmock
	userRepo *UserRepository
}

func (suite *UserRepositoryTestSuite) SetupSuite() {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		fmt.Printf("an error '%s' was not expected when opening a stub database connection", err)
		os.Exit(1)
	}
	suite.userRepo = NewUserRepository(db)
	suite.mock = mock
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

func (suite *UserRepositoryTestSuite) TestUserRepository_Create() {
	testCases := []helpers.TestCaseUserCreate{
		{
			TestName: "Success",
			User: &models.User{
				FirstName: "Iryna",
				LastName:  "L",
				Email:     "iryna@gmail.com",
				Password:  "pswd",
			},
			Want: &models.User{
				ID:        1,
				FirstName: "Iryna",
				LastName:  "L",
				Email:     "iryna@gmail.com",
				Password:  "pswd",
			},
			WantErr:      false,
			WantErrorMsg: "",
		},
	}

	for _, testCases := range testCases {
		suite.T().Run(
			testCases.TestName, func(t *testing.T) {
				query := fmt.Sprintf("insert into %s (firstname, lastname, email, password) value (?, ?, ?, ?)", UsersTable)
				suite.mock.ExpectExec(query).WithArgs(
					testCases.User.FirstName, testCases.User.LastName, testCases.User.Email, testCases.User.Password,
				).WillReturnResult(sqlmock.NewResult(1, 1))

				res, err := suite.userRepo.Create(testCases.User)
				if err != nil {
					t.Errorf("error was not expected while updating stats: %s", err)
				}

				if err := suite.mock.ExpectationsWereMet(); err != nil {
					t.Errorf("there were unfulfilled expectations: %s", err)
				}
				if !testCases.WantErr {
					assert.NoError(t, err)
					assert.Equal(t, res, testCases.Want)
				} else {
					assert.Error(t, err)
					assert.Contains(t, err.Error(), testCases.WantErrorMsg)
				}
			},
		)
	}
}

func (suite *UserRepositoryTestSuite) TestUserRepository_GetByID() {
	testCases := []helpers.TestCaseUserGetByID{
		{
			TestName: "Success",
			UserID:   1,
			Want: &models.User{
				ID:        1,
				FirstName: "Iryna",
				LastName:  "L",
				Email:     "iryna@gmail.com",
				Password:  "pswd",
			},
			WantErr:      false,
			WantErrorMsg: "",
		},
	}

	for _, testCases := range testCases {
		suite.T().Run(
			testCases.TestName, func(t *testing.T) {
				query := fmt.Sprintf(
					"select id, firstname, lastname, email, password from %s where id = ? and deleted is null", UsersTable,
				)

				expectedRows := sqlmock.NewRows([]string{"id", "firstname", "lastname", "email", "password"}).
					AddRow(
						testCases.Want.ID, testCases.Want.FirstName, testCases.Want.LastName, testCases.Want.Email,
						testCases.Want.Password,
					)

				suite.mock.ExpectQuery(query).WithArgs(testCases.UserID).WillReturnRows(expectedRows)

				res, err := suite.userRepo.GetByID(testCases.UserID)
				if err != nil {
					t.Errorf("error was not expected while updating stats: %s", err)
				}

				if err := suite.mock.ExpectationsWereMet(); err != nil {
					t.Errorf("there were unfulfilled expectations: %s", err)
				}

				if !testCases.WantErr {
					assert.NoError(t, err)
					assert.Equal(t, testCases.Want, res)
				} else {
					assert.Error(t, err)
					assert.Contains(t, err.Error(), testCases.WantErrorMsg)
				}
			},
		)
	}
}
