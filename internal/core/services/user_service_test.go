package services

import (
	"fmt"
	"login-app/internal/api/web/user/dto"
	"login-app/internal/core/repository"
	"login-app/platform/logger"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenExistingUser_whenCreate_thenReturnError(t *testing.T) {
	s := NewService(logger.LoggerMock{}, repository.MockUserRepository{})

	id, err := s.Create(&dto.CreateUserRequest{
		Name:                 "John Doe",
		Email:                "johndoe@email.com",
		Password:             "12345678",
		PasswordConfirmation: "12345678",
	})

	assert.Equal(t, "", id)

	if assert.Error(t, err) {
		assert.Equal(t, fmt.Errorf("an user already exists with the same email johndoe@email.com"), err)
	}
}

func TestGivenUserRequest_whenCreate_thenReturnObjectId(t *testing.T) {
	s := NewService(logger.LoggerMock{}, repository.MockUserRepository{})

	id, err := s.Create(&dto.CreateUserRequest{
		Name:                 "Jane Doe",
		Email:                "janedoe@email.com",
		Password:             "12345678",
		PasswordConfirmation: "12345678",
	})

	assert.Equal(t, "id", id)
	assert.Nil(t, err)
}
