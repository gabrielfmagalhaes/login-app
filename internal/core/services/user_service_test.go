package services

import (
	"login-app/internal/api/web/user/dto"
	"login-app/internal/core/domain"
	"login-app/platform/logger"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockUserRepository struct {
}

func (r MockUserRepository) Insert(user *domain.User) (string, error) {
	return "id", nil
}

func TestGivenUserRequest_whenNewUser_thenReturnObjectId(t *testing.T) {
	s := NewService(logger.LoggerMock{}, MockUserRepository{})

	id, err := s.Create(&dto.CreateUserRequest{
		Name:                 "john doe",
		Email:                "johndoe@email.com",
		Password:             "12345678",
		PasswordConfirmation: "12345678",
	})

	assert.Equal(t, "id", id)
	assert.Nil(t, err)
}
