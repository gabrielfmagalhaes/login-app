package mapper

import (
	"login-app/internal/api/user/in"
	"login-app/internal/core/domain"
)

func ToDomain(r in.CreateUserRequest) *domain.User {
	return &domain.User{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
	}
}
