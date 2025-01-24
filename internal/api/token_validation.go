package api

import (
	"context"
	"ewallet-framework/cmd/proto/tokenvalidation"
	"ewallet-framework/helpers"
	"ewallet-framework/internal/constants"
	"ewallet-framework/internal/interfaces"
	"fmt"
)

type TokenValidationHandler struct {
	TokenValidationService interfaces.ITokenValidationService
	tokenvalidation.UnimplementedTokenValidationServer
}

func (s *TokenValidationHandler) ValidateToken(ctx context.Context, req *tokenvalidation.TokenRequest) (*tokenvalidation.TokenResponse, error) {
	var (
		token = req.Token
		log   = helpers.Logger
	)

	if token == "" {
		err := fmt.Errorf("token is empty")
		log.Error(err)
		return &tokenvalidation.TokenResponse{
			Message: err.Error(),
		}, nil
	}

	claimToken, err := s.TokenValidationService.TokenValidation(ctx, token)
	if err != nil {
		return &tokenvalidation.TokenResponse{
			Message: err.Error(),
		}, nil
	}

	return &tokenvalidation.TokenResponse{
		Message: constants.SuccessMessage,
		Data: &tokenvalidation.UserData{
			UserId:   int64(claimToken.UserID),
			Username: claimToken.Username,
			FullName: claimToken.Fullname,
			Email:    claimToken.Email,
		},
	}, nil

}
