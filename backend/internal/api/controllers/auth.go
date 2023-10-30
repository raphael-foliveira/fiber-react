package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raphael-foliveira/fiber-react/backend/internal/api/services"
	"github.com/raphael-foliveira/fiber-react/backend/internal/dto"
)

type Auth struct {
	service *services.Auth
}

func NewAuth(service *services.Auth) *Auth {
	return &Auth{service}
}

// Login godoc
// @Summary Login
// @Description Login
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body dto.Login true "Credentials"
// @Success 200 {object} dto.LoginResponse
// @Failure 401 {object} errs.HTTPError
// @Router /auth/login [post]
func (a *Auth) Login(c *fiber.Ctx) error {
	credentials := dto.Login{}
	if err := c.BodyParser(&credentials); err != nil {
		return err
	}
	loginResponse, err := a.service.Login(&credentials)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(loginResponse)
}

// Signup godoc
// @Summary Signup
// @Description Signup
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dto.CreateUser true "User"
// @Success 201 {object} dto.LoginResponse
// @Failure 400 {object} errs.HTTPError
// @Router /auth/signup [post]
func (a *Auth) Signup(c *fiber.Ctx) error {
	user := dto.CreateUser{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	loginResponse, err := a.service.Signup(&user)
	if err != nil {
		return err
	}
	return c.Status(201).JSON(loginResponse)
}

// Logout godoc
// @Summary Logout
// @Description Logout
// @Tags auth
// @Accept json
// @Produce json
// @Param refreshToken body dto.RefreshToken true "Refresh Token"
// @Success 204
// @Failure 401 {object} errs.HTTPError
func (a *Auth) Logout(c *fiber.Ctx) error {
	refreshToken := dto.RefreshToken{}
	if err := c.BodyParser(&refreshToken); err != nil {
		return err
	}
	user := c.Locals("user").(*dto.User)
	if err := a.service.Logout(refreshToken.Token, user.ID); err != nil {
		return err
	}
	return c.SendStatus(204)
}

// RefreshToken godoc
// @Summary Refresh Token
// @Description Refresh Token
// @Tags auth
// @Accept json
// @Produce json
// @Param refreshToken body dto.RefreshToken true "Refresh Token"
// @Success 200 {object} dto.RefreshTokenResponse
// @Failure 401 {object} errs.HTTPError
// @Router /auth/refresh-token [post]
func (a *Auth) RefreshToken(c *fiber.Ctx) error {
	refreshToken := dto.RefreshToken{}
	if err := c.BodyParser(&refreshToken); err != nil {
		return err
	}
	token, err := a.service.RefreshToken(&refreshToken)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(token)
}

func (a *Auth) Authenticate(c *fiber.Ctx) error {
	token, err := parseAuthHeader(c)
	if err != nil {
		return err
	}
	user, err := a.service.Authenticate(token)
	if err != nil {
		return err
	}
	c.Locals("user", user)
	return c.Next()
}
