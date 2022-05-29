package common

import "net/http"

func NewCreatedResponse() (int, ControllerResponse) {
	return http.StatusCreated, ControllerResponse{
		"created",
		"Created",
		map[string]interface{}{},
	}
}
