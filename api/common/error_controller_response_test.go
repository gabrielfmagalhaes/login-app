package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBadRequestResponse(t *testing.T) {
	expect := struct {
		statusCode int
		response   controllerResponse
	}{
		400,
		controllerResponse{
			errBadRequest,
			"Bad Request",
			map[string]interface{}{},
		},
	}

	statusCode, response := NewBadRequestResponse()

	assert.Equal(t, expect.statusCode, statusCode)

	if assert.NotNil(t, response) {
		assert.Equal(t, expect.response.Code, response.Code)
		assert.Equal(t, expect.response.Data, response.Data)
		assert.Equal(t, expect.response.Message, response.Message)
	}
}
