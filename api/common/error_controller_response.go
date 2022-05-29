package common

import "net/http"

func NewBadRequestResponse() (int, ControllerResponse) {
	return http.StatusBadRequest, ControllerResponse{
		"bad_request",
		"Bad Request",
		map[string]interface{}{},
	}
}
