package helpers

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/requests"
)

type TestCaseGetBearerToken struct {
	TestName     string
	BearerString string
	Want         string
}

type TestCaseValidateToken struct {
	TestName     string
	AccessToken  string
	WantError    bool
	WantErrorMsg string
	WantID       int32
}

type TestCaseUserCreate struct {
	TestName     string
	User         *models.User
	Want         *models.User
	WantErr      bool
	WantErrorMsg string
}

type TestCaseUserGetByID struct {
	TestName     string
	UserID       int32
	Want         *models.User
	WantErr      bool
	WantErrorMsg string
}

type Request struct {
	Method string
	Url    string
	Token  string
}

type ExpectedResponse struct {
	StatusCode int
	BodyPart   string
}

type TestCaseHandler struct {
	TestName    string
	Request     Request
	Body        string
	HandlerFunc func(w http.ResponseWriter, r *http.Request)
	Want        ExpectedResponse
}

type TestCaseTokenHandler struct {
	TestName    string
	Request     Request
	Body        *requests.LoginRequest
	HandlerFunc func(w http.ResponseWriter, r *http.Request)
	Want        ExpectedResponse
}

type TestCaseMiddleware struct {
	TestName    string
	Request     Request
	HeaderName  string
	HeaderValue string
	Want        ExpectedResponse
}

func PrepareHandlerTestCase(test TestCaseHandler) (request *http.Request, recorder *httptest.ResponseRecorder) {
	request = httptest.NewRequest(test.Request.Method, test.Request.Url, strings.NewReader(""))

	if test.Request.Token != "" {
		request.Header.Set("Authorization", "Bearer "+test.Request.Token)
	}

	return request, httptest.NewRecorder()
}
