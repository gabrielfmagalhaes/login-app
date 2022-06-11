package common

import (
	"login-app/internal/api/out"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatedResponse(t *testing.T) {
	expect := struct {
		statusCode int
		response   out.ControllerResponse
	}{
		http.StatusCreated,
		out.ControllerResponse{
			Code:    "created",
			Message: "Created",
			Data:    "any-data",
		},
	}

	statusCode, response := NewCreatedResponse("any-data")

	assert.Equal(t, expect.statusCode, statusCode)

	if assert.NotNil(t, response) {
		assert.Equal(t, expect.response.Code, response.Code)
		assert.Equal(t, expect.response.Data, response.Data)
		assert.Equal(t, expect.response.Message, response.Message)
	}
}
