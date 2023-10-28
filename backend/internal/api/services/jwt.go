package services

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/raphael-foliveira/fiber-react/backend/internal/apperror"
	"github.com/raphael-foliveira/fiber-react/backend/internal/dto"
)

var accessJwtSecret = []byte(os.Getenv("ACCESS_JWT_SECRET"))
var refreshJwtSecret = []byte(os.Getenv("REFRESH_JWT_SECRET"))

type JwtClaims struct {
	Sub   int    `json:"sub"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type JwtService struct{}

func NewJwtService() *JwtService {
	return &JwtService{}
}

func (j *JwtService) GenerateAccessToken(user *dto.User) (string, error) {
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

func (j *JwtService) GenerateRefreshToken(user *dto.User) (string, error) {
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

func (j *JwtService) ValidateToken(token string, isRefreshToken bool) (*dto.User, error) {
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
		return nil, apperror.HTTPError{Code: 401, Message: "invalid token"}
	}
	claims, ok := parsedToken.Claims.(*JwtClaims)
	if !ok {
		return nil, apperror.HTTPError{Code: 401, Message: "invalid token"}
	}
	return &dto.User{
		ID:    claims.Sub,
		Email: claims.Email,
	}, nil
}
