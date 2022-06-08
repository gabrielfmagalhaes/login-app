package common

import (
	"login-app/api/out"
	"net/http"
)

func NewBadRequestResponse(data interface{}) (int, out.ControllerResponse) {
	return http.StatusBadRequest, out.ControllerResponse{
		Code:    "bad_request",
		Message: "Bad Request",
		Data:    data,
	}
}
