package common

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatedResponse(t *testing.T) {
	expect := struct {
		statusCode int
		response   ControllerResponse
	}{
		http.StatusCreated,
		ControllerResponse{
			"created",
			"Created",
			map[string]interface{}{},
		},
	}

	statusCode, response := NewCreatedResponse()

	assert.Equal(t, expect.statusCode, statusCode)

	if assert.NotNil(t, response) {
		assert.Equal(t, expect.response.Code, response.Code)
		assert.Equal(t, expect.response.Data, response.Data)
		assert.Equal(t, expect.response.Message, response.Message)
	}
}
