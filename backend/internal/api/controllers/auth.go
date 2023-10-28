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

func (a *Auth) Signup(c *fiber.Ctx) error {
	user := dto.CreateUser{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	return c.Status(201).JSON(&user)
}

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
