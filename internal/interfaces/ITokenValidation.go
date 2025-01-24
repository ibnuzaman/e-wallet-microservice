package interfaces

import (
	"context"
	"ewallet-framework/cmd/proto/tokenvalidation"
	"ewallet-framework/helpers"
)

type ITokenValidationHandler interface {
	TokenValidationHandler(ctx context.Context, req *tokenvalidation.TokenRequest) (*tokenvalidation.TokenResponse, error)
}

type ITokenValidationService interface {
	TokenValidation(ctx context.Context, token string) (*helpers.ClaimToken, error)
}
