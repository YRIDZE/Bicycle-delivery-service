package handlers

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/YRIDZE/Bicycle-delivery-service/conf"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/services"
	"github.com/YRIDZE/Bicycle-delivery-service/tests/helpers"
	log "github.com/YRIDZE/yolo-log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const userID = 1

type UserHandlerTestSuite struct {
	suite.Suite
	accessToken string
	userHandler *UserHandler
	h           *AppHandlers
}

func (suite *UserHandlerTestSuite) SetupSuite() {
	logger, _ := log.NewLogger(log.LoggerParams{})
	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		fmt.Printf("an error '%s' was not expected when opening a stub database connection", err)
		os.Exit(1)
	}

	cfg := &conf.ConfigToken{
		AccessSecret:          "access",
		AccessLifetimeMinutes: 1,
	}

	tokenService := services.NewTokenService(cfg, logger, db_repository.NewTokensRepositoryMock(db))
	suite.userHandler = NewUserHandler(cfg, logger, db_repository.NewUserRepositoryMock(db), db_repository.NewTokensRepositoryMock(db))
	_, suite.accessToken, _ = tokenService.GenerateAccessToken(userID)

	suite.h = NewAppHandlers(suite.userHandler)
}

func TestUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

func (suite *UserHandlerTestSuite) TestUserHandler_GetProfile() {
	t := suite.T()
	handlerFunc := suite.userHandler.GetProfile
	cases := []helpers.TestCaseHandler{
		{
			TestName: "Successfully get user profile",
			Request: helpers.Request{
				Method: http.MethodGet,
				Url:    "/getUser",
				Token:  suite.accessToken,
			},
			HandlerFunc: handlerFunc,
			Want: helpers.ExpectedResponse{
				StatusCode: 200,
				BodyPart:   "",
			},
		},
		{
			TestName: "Unauthorized getting user profile",
			Request: helpers.Request{
				Method: http.MethodGet,
				Url:    "/getUser",
				Token:  "",
			},
			HandlerFunc: handlerFunc,
			Want: helpers.ExpectedResponse{
				StatusCode: 401,
				BodyPart:   "bad token",
			},
		},
	}

	for _, test := range cases {
		t.Run(
			test.TestName, func(t *testing.T) {
				request, recorder := helpers.PrepareHandlerTestCase(test)
				handler := suite.h.UserHandler.AuthMiddleware(http.HandlerFunc(test.HandlerFunc))
				handler.ServeHTTP(recorder, request)

				assert.Contains(t, recorder.Body.String(), test.Want.BodyPart)
				if assert.Equal(t, test.Want.StatusCode, recorder.Code) {
					if recorder.Code == http.StatusOK {
						helpers.AssertUserProfileResponse(t, recorder)
					}
				}
			},
		)
	}
}
