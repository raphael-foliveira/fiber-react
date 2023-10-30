package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/raphael-foliveira/fiber-react/backend/internal/api/services"
	"github.com/raphael-foliveira/fiber-react/backend/internal/dto"
	"github.com/raphael-foliveira/fiber-react/backend/internal/errs"
)

type Users struct {
	service     *services.Users
	authService *services.Auth
}

func NewUsers(service *services.Users, authService *services.Auth) *Users {
	return &Users{service, authService}
}

// Find godoc
// @Summary Find users
// @Description Find users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} dto.User
// @Failure 400 {object} errs.HTTPError
// @Router /users [get]
func (u *Users) Find(c *fiber.Ctx) error {
	users, err := u.service.Find()
	if err != nil {
		return err
	}
	return c.Status(200).JSON(users)
}

// FindOneById godoc
// @Summary Find user by id
// @Description Find user by id
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} dto.User
// @Failure 400 {object} errs.HTTPError
// @Failure 404 {object} errs.HTTPError
// @Router /users/{id} [get]
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

// FindUserTodos godoc
// @Summary Find user todos
// @Description Find user todos
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {array} dto.Todo
// @Failure 400 {object} errs.HTTPError
// @Failure 404 {object} errs.HTTPError
// @Router /users/{id}/todos [get]
func (u *Users) FindUserTodos(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return errs.HTTPError{Code: 400, Message: "Invalid id"}
	}
	userTodos, err := u.service.FindUserTodos(id)
	if err != nil {
		var notFoundErr errs.NotFoundError
		if errors.As(err, &notFoundErr) {
			return errs.HTTPError{Code: 404, Message: "User not found"}
		}
		return err
	}
	return c.Status(200).JSON(userTodos)
}

// Create godoc
// @Summary Create user
// @Description Create user
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.CreateUser true "User"
// @Success 201 {object} dto.User
// @Failure 400 {object} errs.HTTPError
// @Router /users [post]
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

// Update godoc
// @Summary Update user
// @Description Update user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body dto.UpdateUser true "User"
// @Success 200 {object} dto.User
// @Failure 400 {object} errs.HTTPError
// @Router /users/{id} [put]
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

// Delete godoc
// @Summary Delete user
// @Description Delete user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 204
// @Failure 400 {object} errs.HTTPError
// @Router /users/{id} [delete]
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
