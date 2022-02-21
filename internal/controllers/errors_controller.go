package controllers

import (
	"fmt"
	"net/http"
)

type ErrorsController struct {
}

func (e ErrorsController) ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "Error\n Page not found")
	}

}
