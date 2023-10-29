package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raphael-foliveira/fiber-react/backend/internal/api/services"
	"github.com/raphael-foliveira/fiber-react/backend/internal/dto"
	"github.com/raphael-foliveira/fiber-react/backend/internal/errs"
)

type Todos struct {
	service *services.Todos
}

func NewTodos(service *services.Todos) *Todos {
	return &Todos{service}
}

func (t *Todos) Find(c *fiber.Ctx) error {
	todos, err := t.service.Find()
	if err != nil {
		return err
	}
	return c.Status(200).JSON(todos)
}

func (t *Todos) FindOneById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return errs.HTTPError{Code: 400, Message: "Invalid id"}
	}
	todo, err := t.service.FindOneById(id)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(todo)
}

func (t *Todos) Create(c *fiber.Ctx) error {
	createTodo := dto.CreateTodo{}
	if err := c.BodyParser(&createTodo); err != nil {
		return err
	}
	user := c.Locals("user").(*dto.User)
	createTodo.UserID = user.ID
	todo, err := t.service.Create(&createTodo)
	if err != nil {
		return err
	}
	return c.Status(201).JSON(todo)
}

func (t *Todos) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return errs.HTTPError{Code: 400, Message: "Invalid id"}
	}
	updateTodo := dto.UpdateTodo{}
	if err := c.BodyParser(&updateTodo); err != nil {
		return err
	}
	user := c.Locals("user").(*dto.User)
	todo, err := t.service.Update(id, &updateTodo, user.ID)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(todo)
}

func (t *Todos) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return errs.HTTPError{Code: 400, Message: "Invalid id"}
	}
	user := c.Locals("user").(*dto.User)
	if err := t.service.Delete(id, user.ID); err != nil {
		return err
	}
	return c.Status(204).JSON(nil)
}

func (t *Todos) checkOwner(c *fiber.Ctx, todoId int) error {
	authenticatedUser := c.Locals("user").(*dto.User)
	todo, err := t.service.FindOneById(todoId)
	if err != nil {
		return errs.HTTPError{Code: 404, Message: "Todo not found"}
	}
	if todo.User.ID != authenticatedUser.ID {
		return errs.HTTPError{Code: 403, Message: "You are not the owner of this todo"}
	}
	return nil
}
