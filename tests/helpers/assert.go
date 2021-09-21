package helpers

import (
	"testing"

	"github.com/YRIDZE/Bicycle-delivery-service/helper"
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
