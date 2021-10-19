package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/YRIDZE/Bicycle-delivery-service/conf"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/handlers"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/services"
	"github.com/YRIDZE/Bicycle-delivery-service/tests/helpers"
	log "github.com/YRIDZE/yolo-log"
	"github.com/stretchr/testify/assert"
)

const userID = 1

func TestServer_Run(t *testing.T) {
	logger, _ := log.NewLogger(log.LoggerParams{})
	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		fmt.Printf("an error '%s' was not expected when opening a stub database connection", err)
		os.Exit(1)
	}

	cfg := &conf.Config{
		ConfigServer: &conf.ConfigServer{Port: "8083"},
		ConfigToken: &conf.ConfigToken{
			AccessSecret:          "access",
			AccessLifetimeMinutes: 1,
		},
	}

	tokenService := services.NewTokenService(cfg.ConfigToken, logger, db_repository.NewTokensRepositoryMock(db))
	_, accessToken, _ := tokenService.GenerateAccessToken(userID)
	userHandler := handlers.NewUserHandlerMock(
		cfg.ConfigToken,
		logger,
		db_repository.NewUserRepositoryMock(db),
		db_repository.NewTokensRepositoryMock(db),
	)
	h := handlers.NewAppHandlers((*handlers.UserHandler)(userHandler))

	testSrv := InitServer(cfg, h)
	go func() {
		if err := testSrv.ListenAndServe(); err != nil {
			return
		}
	}()

	cases := []helpers.TestCaseHandler{
		{
			TestName: "Successfully get user profile",
			Request: helpers.Request{
				Token:  accessToken,
				Method: "GET",
				Url:    "/getUser",
			},
			Want: helpers.ExpectedResponse{
				StatusCode: 200,
			},
		},
	}

	for _, test := range cases {
		t.Run(
			test.TestName, func(t *testing.T) {
				w := httptest.NewRecorder()

				req := httptest.NewRequest(test.Request.Method, fmt.Sprintf("http://localhost%s", testSrv.Addr)+test.Request.Url, nil)
				req.Header = map[string][]string{"Authorization": {fmt.Sprintf("Bearer %s", test.Request.Token)}}

				res := h.UserHandler.AuthMiddleware(http.HandlerFunc(userHandler.GetProfile))
				res.ServeHTTP(w, req)

				if assert.NoError(t, err) {
					assert.Equal(t, test.Want.StatusCode, w.Code)
				}
			},
		)
	}
}
