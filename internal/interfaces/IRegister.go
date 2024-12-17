package interfaces

import (
	"context"
	"ewallet-framework/internal/models"
)

type IRegisService interface {
	Register(ctx context.Context, request models.User) (interface{}, error)
}
