package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/YRIDZE/Bicycle-delivery-service/conf"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/requests"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/services"
	"github.com/YRIDZE/Bicycle-delivery-service/tests/helpers"
	log "github.com/YRIDZE/yolo-log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TokenHandlerTestSuite struct {
	suite.Suite
	userHandler  *UserHandler
	testSrv      *httptest.Server
	h            *AppHandlers
	tokenService *services.TokenService
}

func (suite *TokenHandlerTestSuite) SetupSuite() {
	logger, _ := log.NewLogger(log.LoggerParams{})
	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		fmt.Printf("an error '%s' was not expected when opening a stub database connection", err)
		os.Exit(1)
	}

	cfg := &conf.ConfigToken{
		AccessSecret:           "access",
		AccessLifetimeMinutes:  1,
		RefreshSecret:          "refresh",
		RefreshLifetimeMinutes: 5,
	}

	suite.tokenService = services.NewTokenService(cfg, logger, db_repository.NewTokensRepository(db))

	userHandler := NewUserHandler(cfg, logger, db_repository.NewUserRepositoryMock(db), db_repository.NewTokensRepositoryMock(db))
	suite.h = NewAppHandlers(userHandler)
	suite.testSrv = httptest.NewServer(suite.h.InitRoutes())
}

func TestTokenHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(TokenHandlerTestSuite))
}

func (suite *TokenHandlerTestSuite) TestTokenHandler_Login() {
	t := suite.T()
	handlerFunc := suite.h.userHandler.Login
	cases := []helpers.TestCaseTokenHandler{
		{
			TestName: "Successfully logged in",
			Request: helpers.Request{
				Method: http.MethodPost,
				Url:    "/login",
			},
			Body: &requests.LoginRequest{
				Email:    "29@gmail.com",
				Password: "password",
			},
			HandlerFunc: handlerFunc,
			Want: helpers.ExpectedResponse{
				StatusCode: 200,
				BodyPart:   "access_token",
			},
		},
		{
			TestName: "Unauthorized",
			Request: helpers.Request{
				Method: http.MethodPost,
				Url:    "/login",
			},
			Body: &requests.LoginRequest{
				Email:    "29@gmail.com",
				Password: "passwor",
			},
			HandlerFunc: handlerFunc,
			Want: helpers.ExpectedResponse{
				StatusCode: 401,
				BodyPart:   "invalid credentials",
			},
		},
	}

	for _, test := range cases {
		t.Run(
			test.TestName, func(t *testing.T) {
				j, _ := json.Marshal(test.Body)
				request := httptest.NewRequest(test.Request.Method, test.Request.Url, bytes.NewBuffer(j))
				request.Header.Set("Content-Type", "application/json")

				recorder := httptest.NewRecorder()

				test.HandlerFunc(recorder, request)

				assert.Contains(t, recorder.Body.String(), test.Want.BodyPart)
				if assert.Equal(t, test.Want.StatusCode, recorder.Code) {
					if recorder.Code == http.StatusOK {
						var response models.LoginResponse
						err := json.Unmarshal([]byte(recorder.Body.String()), &response)

						assert.NoError(t, err)
					}
				}
			},
		)
	}
}

func (suite *TokenHandlerTestSuite) TestTokenHandler_Refresh() {
	t := suite.T()
	handlerFunc := suite.h.userHandler.Refresh

	_, refreshToken, _ := suite.tokenService.GenerateRefreshToken(userID)

	cases := []helpers.TestCaseTokenHandler{
		{
			TestName: "Successfully refreshed",
			Request: helpers.Request{
				Method: http.MethodPost,
				Url:    "/refresh",
				Token:  refreshToken,
			},
			HandlerFunc: handlerFunc,
			Want: helpers.ExpectedResponse{
				StatusCode: 200,
				BodyPart:   "",
			},
		},
		{
			TestName: "Not refreshed",
			Request: helpers.Request{
				Method: http.MethodPost,
				Url:    "/refresh",
				Token:  refreshToken + "ff",
			},
			HandlerFunc: handlerFunc,
			Want: helpers.ExpectedResponse{
				StatusCode: 401,
				BodyPart:   "something went wrong",
			},
		},
	}

	for _, test := range cases {
		t.Run(
			test.TestName, func(t *testing.T) {

				request := httptest.NewRequest(test.Request.Method, test.Request.Url, nil)
				request.Header.Set("Cookie", fmt.Sprintf("refresh-token=%s", test.Request.Token))
				recorder := httptest.NewRecorder()
				test.HandlerFunc(recorder, request)

				assert.Contains(t, recorder.Body.String(), test.Want.BodyPart)
				if assert.Equal(t, test.Want.StatusCode, recorder.Code) {
					if recorder.Code == http.StatusOK {
						var response models.LoginResponse
						err := json.Unmarshal([]byte(recorder.Body.String()), &response)

						assert.NoError(t, err)
					}
				}
			},
		)
	}
}
