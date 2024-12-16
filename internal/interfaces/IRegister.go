package interfaces

import (
	"context"
	"ewallet-framework/internal/models"
)

type IRegisterRepository interface {
	InsertNewUser(ctx context.Context, user *models.User) error
}

type IRegisterService interface {
	Register(ctx context.Context, request models.User) (interface{}, error)
}
