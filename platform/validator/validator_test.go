package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValidator(t *testing.T) {
	err := GetValidator()

	assert.NotNil(t, err)
}
