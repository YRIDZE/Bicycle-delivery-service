package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
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

	cases := []helpers.TestCaseMiddleware{
		{
			TestName:    "Ok",
			HeaderName:  "Authorization",
			HeaderValue: "Bearer ",
			Request: helpers.Request{
				Method: "GET",
				Url:    "/getUser",
				Token:  accessToken,
			},
			Want: helpers.ExpectedResponse{
				StatusCode: 200,
				BodyPart:   "",
			},
		},
		{
			TestName:    "Invalid Header Name",
			HeaderName:  "",
			HeaderValue: "Bearer ",
			Request: helpers.Request{
				Method: "GET",
				Url:    "/getUser",
				Token:  accessToken,
			},
			Want: helpers.ExpectedResponse{
				StatusCode: 401,
				BodyPart:   "bad token",
			},
		},
		{
			TestName:    "Invalid Header Value",
			HeaderName:  "Authorization",
			HeaderValue: "Bearr ",
			Request: helpers.Request{
				Method: "GET",
				Url:    "/getUser",
				Token:  accessToken,
			},
			Want: helpers.ExpectedResponse{
				StatusCode: 401,
				BodyPart:   "bad token",
			},
		},
		{
			TestName:    "Empty Token",
			HeaderName:  "Authorization",
			HeaderValue: "Bearer ",
			Request: helpers.Request{
				Method: "GET",
				Url:    "/getUser",
				Token:  "",
			},
			Want: helpers.ExpectedResponse{
				StatusCode: 401,
				BodyPart:   "bad token",
			},
		},
	}

	for _, test := range cases {
		t.Run(
			test.TestName, func(t *testing.T) {
				w := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "/getUser", strings.NewReader(""))

				req.Header.Set(test.HeaderName, test.HeaderValue+test.Request.Token)
				res := h.UserHandler.AuthMiddleware(http.HandlerFunc(userHandler.GetProfile))
				res.ServeHTTP(w, req)

				assert.Equal(t, test.Want.StatusCode, w.Code)
				assert.Contains(t, w.Body.String(), test.Want.BodyPart)
			},
		)
	}
}
