package repository

import "login-app/internal/core/domain"

type MockUserRepository struct {
}

func (r MockUserRepository) Insert(user *domain.User) (string, error) {
	return "id", nil
}

func (r MockUserRepository) FindByEmail(email string) (domain.User, error) {
	mockExistingEmail := "johndoe@email.com"

	if email == mockExistingEmail {
		return domain.User{
			ID:       "id",
			Name:     "john doe",
			Email:    "johndoe@email.com",
			Password: "12345678",
		}, nil
	}

	return domain.User{}, nil
}
