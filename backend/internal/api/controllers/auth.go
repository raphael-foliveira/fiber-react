package controllers

import (
	"fmt"

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
	loginResponse, err := a.service.Signup(&user)
	if err != nil {
		return err
	}
	return c.Status(201).JSON(loginResponse)
}

func (a *Auth) Logout(c *fiber.Ctx) error {
	return nil
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

func (a *Auth) Authenticate(c *fiber.Ctx) error {
	token, err := parseAuthHeader(c)
	if err != nil {
		return err
	}
	user, err := a.service.Authenticate(token)
	if err != nil {
		return err
	}
	c.Context().SetUserValue("user", user)
	fmt.Println("user", user.Username)
	return c.Next()
}
