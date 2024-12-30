package repository

import (
	"context"
	"ewallet-framework/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) InsertNewUser(ctx context.Context, user *models.User) error {
	// if err := r.DB.Where("email = ?", user.Email).First(&models.User{}).Error; err == nil {
	// 	return constants.ErrEmailorUsernameAlreadyExist
	// }

	// if err := r.DB.Where("username = ?", user.Username).First(&models.User{}).Error; err == nil {
	// 	return constants.ErrUsernameAlreadyExist
	// }
	// var (
	// 	ErrEmailorUsernameAlreadyExist = errors.New("Email or username already exist")
	// )
	// if err := r.DB.Where("email = ?", user.Email).First(&models.User{}).Error; err == nil {
	// 	return ErrEmailorUsernameAlreadyExist
	// }

	return r.DB.Create(user).Error
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	var (
		user models.User
		err  error
	)
	err = r.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) NewInsertNewUser(ctx context.Context, session *models.UserSession) error {
	return r.DB.Create(session).Error
}

func (r *UserRepository) DeleteUserSession(ctx context.Context, token string) error {
	return r.DB.Exec("DELETE FROM user_sessions WHERE token = ?", token).Error
}

func (r *UserRepository) GetUserSessionByToken(ctx context.Context, token string) (models.UserSession, error) {
	var (
		session models.UserSession
		err     error
	)
	err = r.DB.Where("token = ?", token).First(&session).Error
	if err != nil {
		return session, err
	}
	return session, nil
}
