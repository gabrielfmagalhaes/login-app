package adapter

import (
	"login-app/internal/api/web/user/dto"
	"login-app/internal/core/domain"
)

func ToDomain(dto *dto.CreateUserRequest) *domain.User {
	return &domain.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
}
