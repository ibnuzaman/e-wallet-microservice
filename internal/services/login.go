package services

import (
	"context"
	"ewallet-framework/internal/interfaces"
	"ewallet-framework/internal/models"
)

type LoginService struct {
	UserRepo interfaces.IUserRepository
}

func (s *LoginService) GetUserByUsername(ctx context.Context, req models.LoginRequest) (models.User, error) {
	user, err := s.UserRepo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
