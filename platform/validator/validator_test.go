package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type AnyStructStub struct {
	Name string `json:"name" validate:"required"`
}

func TestGetValidator(t *testing.T) {
	stubStruct := AnyStructStub{}

	err := GetValidator()

	if assert.NotNil(t, err) {
		err := err.Struct(stubStruct)

		assert.NotNil(t, err)
	}
}
