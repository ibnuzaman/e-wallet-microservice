package repository

import (
	"context"
	"ewallet-framework/internal/models"

	"gorm.io/gorm"
)

type RegisterRepository struct {
	DB *gorm.DB
}

func (r *RegisterRepository) InsertNewUser(ctx context.Context, user *models.User) error {
	return r.DB.Create(user).Error
}
