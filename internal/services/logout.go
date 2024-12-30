package services

import (
	"context"

	"ewallet-framework/internal/interfaces"

	"github.com/pkg/errors"
)

type LogoutService struct {
	UserRepo interfaces.IUserRepository
}

func (s *LogoutService) Logout(ctx context.Context, token string) error {
	err := s.UserRepo.DeleteUserSession(ctx, token)
	if err != nil {
		return errors.Wrap(err, "error while deleting user session")
	}
	return nil
}
