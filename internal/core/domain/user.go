package domain

import "login-app/internal/api/web/user/dto"

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

func ToDomain(dto *dto.CreateUserRequest) *User {
	return &User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
}
