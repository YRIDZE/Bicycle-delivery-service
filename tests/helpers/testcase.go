package helpers

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
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
	User         models.User
	Want         models.User
	WantErr      bool
	WantErrorMsg string
}
