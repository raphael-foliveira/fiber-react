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

// FindUserTodos godoc
// @Summary Find user todos
// @Description Find user todos
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "User ID"
// @Success 200 {array} dto.Todo
// @Failure 400 {object} errs.HTTPError
// @Failure 404 {object} errs.HTTPError
// @Security BearerAuth
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

// Update godoc
// @Summary Update user
// @Description Update user
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "User ID"
// @Param user body dto.UpdateUser true "User"
// @Success 200 {object} dto.User
// @Failure 400 {object} errs.HTTPError
// @Security BearerAuth
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
// @Param Authorization header string true "Bearer token"
// @Param id path int true "User ID"
// @Success 204
// @Failure 400 {object} errs.HTTPError
// @Security BearerAuth
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
