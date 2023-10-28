package services

import (
	"github.com/raphael-foliveira/fiber-react/backend/internal/apperror"
	"github.com/raphael-foliveira/fiber-react/backend/internal/dto"
	"github.com/raphael-foliveira/fiber-react/backend/internal/persistence/repositories"
)

type Auth struct {
	tokensRepository repositories.RefreshTokensRepository
	usersService     Users
}

func NewAuth(tokensRepository repositories.RefreshTokensRepository, usersService Users) *Auth {
	return &Auth{tokensRepository, usersService}
}

func (a *Auth) Login(email, password string) (*dto.LoginResponse, error) {
	user, err := a.usersService.FindOneByEmail(email)
	if err != nil {
		return nil, err
	}
	refreshToken, err := a.tokensRepository.Create("1", user.ID)
	return &dto.LoginResponse{
		RefreshToken: refreshToken.Token,
		AccessToken:  "",
	}, nil
}

func (a *Auth) Signup(user *dto.CreateUser) (*dto.LoginResponse, error) {
	if user.Password != user.ConfirmPassword {
		return nil, apperror.HTTPError{Code: 400, Message: "passwords do not match"}
	}
	newUser, err := a.usersService.Create(user)
	if err != nil {
		return nil, err
	}
	return a.Login(newUser.Email, newUser.Password)
}
