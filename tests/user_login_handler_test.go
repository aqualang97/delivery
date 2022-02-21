package tests

import (
	config "delivery/configs"
	"delivery/internal/auth/services"
	"delivery/internal/controllers"
	"delivery/internal/models"
	"delivery/internal/repositories/my_mocks"
	myLog "delivery/logs"
	"delivery/tests/helpers"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http/httptest"
	"testing"
)

type AuthHandlerTestSuite struct {
	suite.Suite
	user         *models.User
	tokenService *services.TokenService
	authCtr      *controllers.AuthController
	testSRV      *httptest.Server
	cfg          *config.Config
}

func (suite *AuthHandlerTestSuite) SetupSuite() {
	cfg := &config.Config{
		AccessSecret:           "access",
		RefreshSecret:          "refresh",
		AccessLifetimeMinutes:  1,
		RefreshLifetimeMinutes: 1,
		Logger:                 myLog.LogInit(),
	}
	p, err := bcrypt.GenerateFromPassword([]byte("11111111"), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	l := myLog.LogInit()
	tokenService := services.NewTokenService(cfg)
	suite.tokenService = tokenService

	authController := controllers.NewAuthController(cfg, l, my_mocks.NewUserRepoMock(), my_mocks.NewUserAccessRepoMock(cfg), my_mocks.NewUserRefreshRepoMock(cfg))
	//newGoMock := gomock.NewController(suite.T())
	//authController := controllers.NewAuthController(cfg, l,
	//	mock.NewMockUserRepositoryInterface(newGoMock),
	//	mock.NewMockUserAccessTokenRepositoryInterface(newGoMock),
	//	mock.NewMockUserRefreshTokenRepositoryInterface(newGoMock))
	suite.cfg = cfg
	suite.authCtr = authController
	suite.user = &models.User{
		ID:           1,
		Email:        "alex-test@example.com",
		Login:        "Alex",
		PasswordHash: string(p),
		CreatedAt:    nil,
		UpdatedAt:    nil,
	}
	//suite.authCtr = authController
}

func (suite *AuthHandlerTestSuite) SetupTest() {

}
func (suite *AuthHandlerTestSuite) TearDownSuite() {

}
func (suite *AuthHandlerTestSuite) TearDownTest() {

}

func TestAuthHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(AuthHandlerTestSuite))
}

func (suite AuthHandlerTestSuite) TestLogin() {

	t := suite.T()
	handFunc := suite.authCtr.Login
	//p1, _ := bcrypt.GenerateFromPassword([]byte("11111111"), bcrypt.DefaultCost)

	cases := []helpers.TestCaseLogin{
		{
			TestName:    "Success",
			HandlerFunc: handFunc,
			//User: models.User{
			//	ID:           1,
			//	Email:        "alex-test@example.com",
			//	Login:        "Alex",
			//	PasswordHash: string(p1),
			//	CreatedAt:    nil,
			//	UpdatedAt:    nil,
			//},
			Password: "11111111",
			Expected: helpers.ExpectedResponse{
				StatusCode: 200,
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			user, err := suite.authCtr.UserRepository.GetUserByEmail("alex-test@example.com")
			assert.NoError(t, err)
			err = bcrypt.CompareHashAndPassword([]byte(suite.user.PasswordHash), []byte(test.Password))
			assert.NoError(t, err)
			fmt.Println(user.ID)
			err = suite.authCtr.UserAccessTokenRepository.ExpiredAccessToken(test.User.ID)
			assert.NoError(t, err)
			err = suite.authCtr.UserRefreshTokenRepository.ExpiredRefreshToken(test.User.ID)
			assert.NoError(t, err)

			accessString, err := services.GenerateToken(user.ID, suite.cfg.AccessLifetimeMinutes, suite.cfg.AccessSecret)
			assert.NoError(t, err)

			accessHash, err := services.GetHashOfToken(accessString)
			assert.NoError(t, err)

			refreshString, err := services.GenerateToken(user.ID, suite.cfg.RefreshLifetimeMinutes, suite.cfg.RefreshSecret)

			refreshHash, _ := services.GetHashOfToken(refreshString)
			respA := models.UserAccessToken{
				UserID:      user.ID,
				AccessToken: accessHash,
				CreatedAt:   nil,
				ExpiredAt:   nil,
				Expired:     "false",
			}
			respR := models.UserRefreshToken{
				UserID:       user.ID,
				RefreshToken: refreshHash,
				CreatedAt:    nil,
				ExpiredAt:    nil,
				Expired:      "false",
			}
			err = suite.authCtr.UserAccessTokenRepository.InsertAccessToken(respA)
			err = suite.authCtr.UserRefreshTokenRepository.InsertRefreshToken(respR)

		})
	}
}

//type MiddlewareTestSuite struct {
//	suite.Suite
//
//	middleware middlware.Middleware
//	testSRV    *httptest.Server
//}
