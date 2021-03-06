package common

import (
	"login-app/internal/api/out"
	"net/http"
)

func NewCreatedResponse(data interface{}) (int, out.ControllerResponse) {
	return http.StatusCreated, out.ControllerResponse{
		Code:    "created",
		Message: "Created",
		Data:    data,
	}
}
