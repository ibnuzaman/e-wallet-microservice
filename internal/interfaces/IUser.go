package interfaces

import (
	"context"
	"ewallet-framework/internal/models"
)

type IUserRepository interface {
	InsertNewUser(ctx context.Context, user *models.User) error
	GetUserByUsername(ctx context.Context, username string) (models.User, error)
	NewInsertNewUser(ctx context.Context, session *models.UserSession) error
}
