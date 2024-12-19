package interfaces

import (
	"context"
	"ewallet-framework/internal/models"
)

type ILoginService interface {
	Login(ctx context.Context, req models.LoginRequest) (models.LoginResponse, error)
}
