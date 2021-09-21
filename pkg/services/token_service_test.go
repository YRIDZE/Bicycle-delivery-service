package services

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/YRIDZE/Bicycle-delivery-service/conf"
	"github.com/YRIDZE/Bicycle-delivery-service/helper"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/YRIDZE/Bicycle-delivery-service/tests/helpers"
	log "github.com/YRIDZE/yolo-log"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const userID = 1

type TokenServiceTestSuite struct {
	suite.Suite
	mock         sqlmock.Sqlmock
	cfg          *conf.ConfigToken
	tokenService *TokenService
	token        *db_repository.TokensRepository
}

func (suite *TokenServiceTestSuite) SetupSuite() {
	logger, _ := log.NewLogger(log.LoggerParams{})
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		fmt.Printf("an error '%s' was not expected when opening a stub database connection", err)
		os.Exit(1)
	}
	suite.mock = mock

	suite.cfg = &conf.ConfigToken{
		AccessSecret:           "access",
		AccessLifetimeMinutes:  1,
		RefreshSecret:          "refresh",
		RefreshLifetimeMinutes: 1,
	}

	suite.token = db_repository.NewTokensRepository(db)
	suite.tokenService = NewTokenService(suite.cfg, logger, suite.token)
}

func TestTokenServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TokenServiceTestSuite))
}

func (suite *TokenServiceTestSuite) TestGetTokenFromBearerString() {
	testCases := []helpers.TestCaseGetBearerToken{
		{
			TestName:     "Get token successfully",
			BearerString: "Bearer test_token",
			Want:         "test_token",
		},
		{
			TestName:     "Return empty token if Bearer prefix is absent",
			BearerString: "Beare test_token",
			Want:         "",
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(
			testCase.TestName, func(t *testing.T) {
				got, _ := suite.tokenService.GetTokenFromBearerString(testCase.BearerString)
				assert.Equal(suite.T(), testCase.Want, got)
			},
		)
	}
}

func (suite *TokenServiceTestSuite) TestGenerateAccessToken() {
	_, tokenString, err := suite.tokenService.GenerateAccessToken(userID)

	assert.NoError(suite.T(), err)

	token, err := jwt.ParseWithClaims(
		tokenString, &helper.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(suite.cfg.AccessSecret), nil
		},
	)
	assert.NoError(suite.T(), err)

	claims, ok := token.Claims.(*helper.JwtCustomClaims)
	assert.True(suite.T(), ok)
	assert.True(suite.T(), token.Valid)

	got := claims.ID
	assert.Equal(suite.T(), int32(userID), got)

	expireTime := time.Unix(claims.ExpiresAt, 0)
	assert.WithinDuration(
		suite.T(), time.Now().Add(time.Minute*time.Duration(suite.cfg.AccessLifetimeMinutes)), expireTime, time.Second,
	)
}

func (suite *TokenServiceTestSuite) TestValidateAccessToken() {
	_, tokenString, _ := suite.tokenService.GenerateAccessToken(userID)
	_, refreshTokenString, _ := suite.tokenService.GenerateRefreshToken(userID)
	invalidTokenString := tokenString + "a"

	suite.cfg.AccessLifetimeMinutes = 0
	_, expiredTokenString, _ := suite.tokenService.GenerateAccessToken(userID)

	testCases := []helpers.TestCaseValidateToken{
		{
			TestName:     "Valid token",
			AccessToken:  tokenString,
			WantError:    false,
			WantErrorMsg: "",
			WantID:       userID,
		},
		{
			TestName:     "Invalid token",
			AccessToken:  invalidTokenString,
			WantError:    true,
			WantErrorMsg: "signature is invalid",
			WantID:       0,
		},
		{
			TestName:     "Token with non-expected signature",
			AccessToken:  refreshTokenString,
			WantError:    true,
			WantErrorMsg: "signature is invalid",
			WantID:       0,
		},
		{
			TestName:     "Expired token",
			AccessToken:  expiredTokenString,
			WantError:    true,
			WantErrorMsg: "token is expired",
			WantID:       0,
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(
			testCase.TestName, func(t *testing.T) {
				time.Sleep(500 * time.Millisecond)
				gotClaims, err := suite.tokenService.ValidateAccessToken(testCase.AccessToken)

				helpers.AssertTokenResponse(suite.T(), testCase, gotClaims, err)
			},
		)
	}
}
