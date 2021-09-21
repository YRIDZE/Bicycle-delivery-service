package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/YRIDZE/Bicycle-delivery-service/conf"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/services"
	"github.com/YRIDZE/Bicycle-delivery-service/tests/helpers"
	log "github.com/YRIDZE/yolo-log"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler_AuthMiddleware(t *testing.T) {
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
	userHandler := NewUserHandlerMock(cfg, logger, db_repository.NewUserRepositoryMock(db), db_repository.NewTokensRepositoryMock(db))
	_, accessToken, _ := tokenService.GenerateAccessToken(userID)

	h := NewAppHandlers((*UserHandler)(userHandler))

	cases := []helpers.TestCaseHandler{
		{
			TestName: "Successfully get user profile",
			Request: helpers.Request{
				Token: accessToken,
				Url:   "/getUser",
			},
			Want: helpers.ExpectedResponse{
				StatusCode: 200,
			},
		},
		{
			TestName: "Unauthorized getting user profile",
			Request: helpers.Request{
				Token: "",
				Url:   "/getUser",
			},
			Want: helpers.ExpectedResponse{
				StatusCode: 401,
			},
		},
		{
			TestName: "Unauthorized getting user profile",
			Request: helpers.Request{
				Token: fmt.Sprintf("%s.%s", strings.Split(accessToken, ".")[0], strings.Split(accessToken, ".")[1]),
				Url:   "/getUser",
			},
			Want: helpers.ExpectedResponse{
				StatusCode: 401,
			},
		},
	}

	for _, test := range cases {
		t.Run(
			test.TestName, func(t *testing.T) {
				request, recorder := helpers.PrepareHandlerTestCase(test)
				handler := h.UserHandler.AuthMiddleware(http.HandlerFunc(userHandler.GetProfile))
				handler.ServeHTTP(recorder, request)

				assert.Equal(t, test.Want.StatusCode, recorder.Code)
			},
		)
	}
}
