package services

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/raphael-foliveira/fiber-react/backend/internal/dto"
	"github.com/raphael-foliveira/fiber-react/backend/internal/errs"
)

var accessJwtSecret = []byte(os.Getenv("ACCESS_JWT_SECRET"))
var refreshJwtSecret = []byte(os.Getenv("REFRESH_JWT_SECRET"))

type JwtClaims struct {
	Sub   int    `json:"sub"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type Jwt struct{}

func NewJwt() *Jwt {
	return &Jwt{}
}

func (j *Jwt) GenerateTokens(user *dto.User) (*dto.LoginResponse, error) {
	accessToken, err := j.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}
	refreshToken, err := j.GenerateRefreshToken(user)
	if err != nil {
		return nil, err
	}
	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (j *Jwt) GenerateAccessToken(user *dto.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtClaims{
		user.ID,
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "fiber-react",
		},
	})
	return token.SignedString(accessJwtSecret)
}

func (j *Jwt) GenerateRefreshToken(user *dto.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtClaims{
		user.ID,
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(720 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "fiber-react",
		},
	})
	return token.SignedString(refreshJwtSecret)
}

func (j *Jwt) ValidateToken(token string, isRefreshToken bool) (*dto.User, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if isRefreshToken {
			return refreshJwtSecret, nil
		}
		return accessJwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !parsedToken.Valid {
		return nil, errs.HTTPError{Code: 401, Message: "invalid token"}
	}
	claims, ok := parsedToken.Claims.(*JwtClaims)
	if !ok {
		return nil, errs.HTTPError{Code: 401, Message: "invalid token"}
	}
	return &dto.User{
		ID:    claims.Sub,
		Email: claims.Email,
	}, nil
}
