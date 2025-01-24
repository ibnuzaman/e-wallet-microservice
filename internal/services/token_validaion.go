package services

import (
	"context"
	"ewallet-framework/helpers"
	"ewallet-framework/internal/interfaces"
)

type TokenValidationService struct {
	UserRepo interfaces.IUserRepository
}

func (s *TokenValidationService) TokenValidation(ctx context.Context, token string) (*helpers.ClaimToken, error) {
	var (
		claimToken *helpers.ClaimToken
		err        error
	)

	claimToken, err = helpers.ValidateToken(ctx, token)
	if err != nil {
		return nil, err
	}

	_, err = s.UserRepo.GetUserSessionByToken(ctx, token)
	if err != nil {
		return nil, err
	}

	return claimToken, nil
}
