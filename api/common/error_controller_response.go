package common

import "net/http"

type responseCode string

const (
	errBadRequest responseCode = "bad_request"
)

type controllerResponse struct {
	Code    responseCode `json:"code"`
	Message string       `json:"message"`
	Data    interface{}  `json:"data"`
}

func NewBadRequestResponse() (int, controllerResponse) {
	return http.StatusBadRequest, controllerResponse{
		errBadRequest,
		"Bad Request",
		map[string]interface{}{},
	}
}
