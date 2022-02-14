package tests

import (
	config "delivery/configs"
	"delivery/internal/controllers"
	myLog "delivery/logs"
	"github.com/stretchr/testify/suite"
)

type AuthHandlerTestSuite struct {
	suite.Suite
	cfg            *config.Config
	authController controllers.AuthController
	//testSRV        *httptest.Server
}

func (suite *AuthHandlerTestSuite) SetupSuite() {
	cfg := &config.Config{
		AccessSecret:           "access",
		RefreshSecret:          "refresh",
		AccessLifetimeMinutes:  1,
		RefreshLifetimeMinutes: 1,
		Logger:                 myLog.LogInit(),
	}

	suite.cfg = cfg

}

func (suite *AuthHandlerTestSuite) SetupTest() {

}
func (suite *AuthHandlerTestSuite) TearDownSuite() {

}
func (suite *AuthHandlerTestSuite) TearDownTest() {

}

//type MiddlewareTestSuite struct {
//	suite.Suite
//
//	middleware middlware.Middleware
//	testSRV    *httptest.Server
//}
