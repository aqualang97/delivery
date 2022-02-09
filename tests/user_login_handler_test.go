package tests

import (
	handProv "delivery/internal/auth/handle_provide"
	"delivery/internal/auth/middlware"
	"github.com/stretchr/testify/suite"
	"net/http/httptest"
)

type AuthLoginHandlerTestSuite struct {
	suite.Suite

	loginHandler *handProv.HandlerProvider
	testSRV      *httptest.Server
}
type MiddlewareTestSuite struct {
	suite.Suite

	middleware middlware.Middleware
	testSRV    *httptest.Server
}
