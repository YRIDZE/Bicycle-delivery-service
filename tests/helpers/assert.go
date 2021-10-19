package helpers

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/YRIDZE/Bicycle-delivery-service/helper"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/stretchr/testify/assert"
)

func AssertTokenResponse(t *testing.T, testCase TestCaseValidateToken, gotClaims *helper.JwtCustomClaims, err error) {
	t.Helper()

	if testCase.WantError {
		assert.Error(t, err)
		assert.Contains(t, err.Error(), testCase.WantErrorMsg)
	} else {
		assert.NoError(t, err)
		assert.Equal(t, testCase.WantID, gotClaims.ID)
	}
}

func AssertUserProfileResponse(t *testing.T, recorder *httptest.ResponseRecorder) {
	t.Helper()

	var response models.UserResponse
	err := json.Unmarshal([]byte(recorder.Body.String()), &response)

	if assert.NoError(t, err) {
		assert.Equal(
			t, models.UserResponse{
				ID:        1,
				FirstName: "firstname",
				LastName:  "lastname",
				Email:     "email",
			}, response,
		)
	}
}

func AssertGetSupplierResponse(t *testing.T, recorder *httptest.ResponseRecorder) {
	t.Helper()

	var response models.SupplierResponse
	err := json.Unmarshal([]byte(recorder.Body.String()), &response)

	if assert.NoError(t, err) {
		assert.Equal(
			t, models.SupplierResponse{
				ID:    1,
				Name:  "name",
				Type:  "type",
				Image: "image",
				WorkHours: models.WorkingHours{
					Opening: "open",
					Closing: "close",
				},
				Deleted: "",
			}, response,
		)
	}
}
