package services

import (
	"fmt"
	"login-app/internal/api/web/user/dto"
	"login-app/internal/core/domain"
	"login-app/internal/core/repository"
	"login-app/platform/logger"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	logger     logger.Logger
	repository repository.Repository
}

func NewService(logger logger.Logger, repository repository.Repository) *Service {
	return &Service{logger: logger, repository: repository}
}

func (s *Service) Create(request *dto.CreateUserRequest) (string, error) {
	findUser, err := s.repository.FindByEmail(request.Email)

	if findUser != (domain.User{}) {
		return "", fmt.Errorf("an user already exists with the same email %s", findUser.Email)
	}

	if err != nil {
		s.logger.Warnf("Error while trying to find user", err)

		return "", fmt.Errorf("failed to create user %s", request.Email)
	}

	user := domain.ToDomain(request)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		s.logger.Warnf("Error while trying hashify password", err)

		return "", fmt.Errorf("failed to create user %s", user.Email)
	}

	user.Password = string(hashedPassword)

	id, err := s.repository.Insert(user)

	if err != nil {
		s.logger.Warnf("Error while trying to create user", err)

		return "", fmt.Errorf("failed to create user %s", user.Email)
	}

	s.logger.Infof("User created", id)

	return id, nil
}
