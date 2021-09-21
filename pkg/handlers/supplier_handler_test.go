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
	"github.com/YRIDZE/Bicycle-delivery-service/tests/helpers"
	log "github.com/YRIDZE/yolo-log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SupplierHandlerTestSuite struct {
	suite.Suite
	supplierHandler *SupplierHandler
	testSrv         *httptest.Server
}

func (suite *SupplierHandlerTestSuite) SetupSuite() {
	logger, _ := log.NewLogger(log.LoggerParams{})
	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		fmt.Printf("an error '%s' was not expected when opening a stub database connection", err)
		os.Exit(1)
	}

	userHandler := NewUserHandler(&conf.ConfigToken{}, logger, db_repository.NewUserRepositoryMock(db), db_repository.NewTokensRepository(db))

	supplierRepo := db_repository.NewSupplierRepositoryMock(db)
	suite.supplierHandler = NewSupplierHandler(logger, supplierRepo)
	h := NewAppHandlers(userHandler, suite.supplierHandler)

	suite.testSrv = httptest.NewServer(h.InitRoutes())
}

func TestSupplierHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(SupplierHandlerTestSuite))
}

func (suite *SupplierHandlerTestSuite) TestUserHandler_GetByID() {
	t := suite.T()
	handlerFunc := suite.supplierHandler.GetByID
	cases := []helpers.TestCaseHandler{
		{
			TestName: "Successfully get supplier by id",
			Request: helpers.Request{
				Method: http.MethodGet,
				Url:    "/getSupplierById/?id=1",
			},
			HandlerFunc: handlerFunc,
			Want: helpers.ExpectedResponse{
				StatusCode: 200,
				BodyPart:   "",
			},
		},
		{
			TestName: "Error getting supplier by id",
			Request: helpers.Request{
				Method: http.MethodGet,
				Url:    "/getSupplierById/?id=!",
			},
			HandlerFunc: handlerFunc,
			Want: helpers.ExpectedResponse{
				StatusCode: 404,
				BodyPart:   "invalid id parameter",
			},
		},
	}

	for _, test := range cases {
		t.Run(
			test.TestName, func(t *testing.T) {
				request := httptest.NewRequest(test.Request.Method, test.Request.Url, strings.NewReader(""))
				recorder := httptest.NewRecorder()
				test.HandlerFunc(recorder, request)

				assert.Contains(t, recorder.Body.String(), test.Want.BodyPart)
				if assert.Equal(t, test.Want.StatusCode, recorder.Code) {
					if recorder.Code == http.StatusOK {
						helpers.AssertGetSupplierResponse(t, recorder)
					}
				}
			},
		)
	}
}
