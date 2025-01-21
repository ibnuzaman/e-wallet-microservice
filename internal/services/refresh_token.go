package services

import (
	"context"
	"ewallet-framework/helpers"
	"ewallet-framework/internal/interfaces"
	"ewallet-framework/internal/models"
	"time"
)

type RefreshTokenService struct {
	UserRepo interfaces.IUserRepository
}

func (s *RefreshTokenService) RefreshToken(ctx context.Context, refreshToken string, tokenClaim helpers.ClaimToken) (models.RefreshTokenResponse, error) {
	resp := models.RefreshTokenResponse{}
	token, err := helpers.GenerateToken(ctx, tokenClaim.UserID, tokenClaim.Username, tokenClaim.Fullname, "token", tokenClaim.Email, time.Now())

	if err != nil {
		return resp, err
	}

	err = s.UserRepo.UpdateTokenWByRefreshToken(ctx, token, refreshToken)
	if err != nil {
		return resp, err
	}
	resp.Token = token
	return resp, nil
}
