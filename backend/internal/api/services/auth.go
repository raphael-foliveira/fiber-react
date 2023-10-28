package services

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/raphael-foliveira/fiber-react/backend/internal/dto"
	"github.com/raphael-foliveira/fiber-react/backend/internal/errs"
	"github.com/raphael-foliveira/fiber-react/backend/internal/persistence/repositories"
)

type Auth struct {
	tokensRepository repositories.RefreshTokensRepository
	usersService     *Users
	jwtService       *Jwt
}

func NewAuth(tokensRepository repositories.RefreshTokensRepository, usersService *Users, jwtService *Jwt) *Auth {
	return &Auth{tokensRepository, usersService, jwtService}
}

func (a *Auth) Login(credentials *dto.Login) (*dto.LoginResponse, error) {
	user, err := a.usersService.FindOneByEmail(credentials.Email)
	if err != nil {
		return nil, err
	}
	if user.Password != credentials.Password {
		return nil, errs.HTTPError{Code: 401, Message: "invalid credentials"}
	}
	tokens, err := a.jwtService.GenerateTokens(&dto.User{
		ID:    user.ID,
		Email: user.Email,
	})
	_, err = a.tokensRepository.Create(tokens.RefreshToken, user.ID)
	if err != nil {
		log.Error(err)
		return nil, errs.HTTPError{Code: 500, Message: "could not save refresh token"}
	}
	return &dto.LoginResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		User: &dto.User{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
		},
	}, nil
}

func (a *Auth) Signup(user *dto.CreateUser) (*dto.LoginResponse, error) {
	fmt.Println(user)
	if user.Password != user.ConfirmPassword {
		return nil, errs.HTTPError{Code: 400, Message: "passwords do not match"}
	}
	newUser, err := a.usersService.Create(user)
	if err != nil {
		return nil, err
	}
	return a.Login(&dto.Login{
		Email:    newUser.Email,
		Password: newUser.Password,
	})
}

func (a *Auth) RefreshToken(refreshToken *dto.RefreshToken) (string, error) {
	token, err := a.tokensRepository.FindOne(refreshToken.Token)
	if err != nil {
		return "", errs.HTTPError{Code: 401, Message: "invalid refresh token"}
	}
	user, err := a.usersService.FindOne(token.UserID)
	if err != nil {
		return "", err
	}
	return a.jwtService.GenerateAccessToken(&dto.User{
		ID:    user.ID,
		Email: user.Email,
	})
}
