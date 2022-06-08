package common

import (
	"login-app/api/out"
	"net/http"
)

func NewCreatedResponse(data interface{}) (int, out.ControllerResponse) {
	return http.StatusCreated, out.ControllerResponse{
		Code:    "created",
		Message: "Created",
	}
}
