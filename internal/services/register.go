package services

import (
	"context"
	"ewallet-framework/external"
	"ewallet-framework/internal/interfaces"
	"ewallet-framework/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	RegisterRepo   interfaces.IUserRepository
	ExternalWallet interfaces.IWallet
}

func (s *RegisterService) Register(ctx context.Context, request models.User) (interface{}, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	request.Password = string(hashPassword)

	err = s.RegisterRepo.InsertNewUser(ctx, &request)
	if err != nil {
		return nil, err
	}

	e := external.ExWallet{}
	e.CreateWallet(ctx, request.ID)
	_, err = s.ExternalWallet.CreateWallet(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	resp := request
	resp.Password = ""
	return resp, nil
}
