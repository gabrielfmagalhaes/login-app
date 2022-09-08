package services

import (
	"fmt"
	"login-app/internal/api/web/user/dto"
	"login-app/internal/core/adapter"
	"login-app/platform/logger"
	"login-app/platform/mongodb"
)

type Service struct {
	logger     logger.Logger
	repository *mongodb.Repository
}

func NewService(logger logger.Logger, repository *mongodb.Repository) *Service {
	return &Service{logger: logger, repository: repository}
}

func (s *Service) Create(request *dto.CreateUserRequest) (string, error) {
	user := adapter.ToDomain(request)

	id, err := s.repository.Insert(user)

	if err != nil {
		s.logger.Warnf("Error while trying to create user", err)

		return "", fmt.Errorf("Failed to create user %s", user.Email)
	}

	s.logger.Infof("User created", id)

	return id, nil
}
