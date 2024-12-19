package services

import (
	"context"
	"time"

	"ewallet-framework/helpers"
	"ewallet-framework/internal/interfaces"
	"ewallet-framework/internal/models"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	UserRepo interfaces.IUserRepository
}

func (s *LoginService) Login(ctx context.Context, req models.LoginRequest) (models.LoginResponse, error) {
	var (
		resp models.LoginResponse
		now  = time.Now()
	)
	userDetail, err := s.UserRepo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return resp, errors.Wrap(err, "error while getting user by username")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(req.Password)); err != nil {
		return resp, errors.Wrap(err, "invalid password")
	}

	token, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, "jwt", now)
	if err != nil {
		return resp, errors.Wrap(err, "error while generating token")
	}

	refreshToken, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, "jwt", now)
	if err != nil {
		return resp, errors.Wrap(err, "error while generating token")
	}

	userSession := models.UserSession{
		UserID:              userDetail.ID,
		Token:               token,
		RefreshToken:        refreshToken,
		TokenExpired:        now.Add(helpers.MapTypeToken["token"]),
		RefreshTokenExpired: now.Add(helpers.MapTypeToken["refresh_token"]),
	}
	err = s.UserRepo.NewInsertNewUser(ctx, &userSession)
	if err != nil {
		return resp, errors.Wrap(err, "error while inserting new user")
	}

	resp.UserID = userDetail.ID
	resp.Username = userDetail.Username
	resp.FullName = userDetail.FullName
	resp.Email = userDetail.Email
	resp.Token = token
	resp.RefreshToken = refreshToken

	return resp, nil
}
