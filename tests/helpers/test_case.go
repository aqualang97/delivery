package helpers

import (
	"delivery/internal/models"
	"net/http"
)

type ExpectedResponse struct {
	StatusCode int
	//BodyPast   string
}

type TestCaseLogin struct {
	TestName    string
	User        models.User
	Password    string
	HandlerFunc func(w http.ResponseWriter, r *http.Request)
	Expected    ExpectedResponse
}
