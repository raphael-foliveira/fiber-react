package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/raphael-foliveira/fiber-react/backend/internal/api/services"
	"github.com/raphael-foliveira/fiber-react/backend/internal/dto"
	"github.com/raphael-foliveira/fiber-react/backend/internal/errs"
)

type Users struct {
	service    *services.Users
	jwtService *services.Jwt
}

func NewUsers(service *services.Users, jwtService *services.Jwt) *Users {
	return &Users{service, jwtService}
}

func (u *Users) Find(c *fiber.Ctx) error {
	users, err := u.service.Find()
	if err != nil {
		return err
	}
	return c.Status(200).JSON(users)
}

func (u *Users) FindOneById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return errs.HTTPError{Code: 400, Message: "Invalid id"}
	}
	user, err := u.service.FindOne(id)
	if err != nil {
		var notFoundErr errs.NotFoundError
		if errors.As(err, &notFoundErr) {
			return errs.HTTPError{Code: 404, Message: "User not found"}
		}
		return err
	}
	return c.Status(200).JSON(user)
}

func (u *Users) Me(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	authorization, ok := headers["Authorization"]
	if !ok {
		return errs.HTTPError{Code: 401, Message: "Unauthorized"}
	}
	user, err := u.jwtService.ValidateToken(authorization[0], false)
	if err != nil {
		return errs.HTTPError{Code: 401, Message: "Unauthorized"}
	}
	userTodos, err := u.service.FindUserTodos(user.ID)
	return c.Status(200).JSON(userTodos)
}

func (u *Users) Create(c *fiber.Ctx) error {
	userDto := dto.CreateUser{}
	if err := c.BodyParser(&userDto); err != nil {
		return err
	}
	user, err := u.service.Create(&userDto)
	if err != nil {
		return err
	}
	return c.Status(201).JSON(user)
}

func (u *Users) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return errs.HTTPError{Code: 400, Message: "Invalid id"}
	}
	updateUser := dto.UpdateUser{}
	if err := c.BodyParser(&updateUser); err != nil {
		return err
	}
	user, err := u.service.Update(id, &updateUser)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(user)
}

func (u *Users) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return errs.HTTPError{Code: 400, Message: "Invalid id"}
	}
	err = u.service.Delete(id)
	if err != nil {
		return err
	}
	return c.Status(204).JSON(nil)
}
