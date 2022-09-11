package common

import (
	"net/http"
)

func NewCreatedResponse(data interface{}) (int, ControllerResponse) {
	return http.StatusCreated, ControllerResponse{
		Code:    "created",
		Message: "Created",
		Data:    data,
	}
}
