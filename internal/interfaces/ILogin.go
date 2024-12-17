package interfaces

import (
	"context"
	"ewallet-framework/internal/models"
)

type ILoginService interface {
	GetUserByUsername(ctx context.Context, req models.LoginRequest) (models.User, error)
}
