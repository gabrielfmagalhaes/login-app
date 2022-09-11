package common

import "net/http"

func NewBadRequestResponse(data interface{}) (int, ControllerResponse) {
	return http.StatusBadRequest, ControllerResponse{
		Code:    "bad_request",
		Message: "Bad Request",
		Data:    data,
	}
}
