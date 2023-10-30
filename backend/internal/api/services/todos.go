package services

import (
	"errors"

	"github.com/raphael-foliveira/fiber-react/backend/internal/dto"
	"github.com/raphael-foliveira/fiber-react/backend/internal/errs"
	"github.com/raphael-foliveira/fiber-react/backend/internal/models"
	"github.com/raphael-foliveira/fiber-react/backend/internal/persistence/repositories"
)

type Todos struct {
	repository repositories.TodosRepository
}

func NewTodos(repository repositories.TodosRepository) *Todos {
	return &Todos{repository}
}

func (t *Todos) Find() ([]*dto.Todo, error) {
	todos, err := t.repository.Find()
	if err != nil {
		return nil, err
	}
	return dto.TodosFromModels(todos), nil
}

func (t *Todos) FindOneById(id int) (*dto.TodoWithUser, error) {
	todo, err := t.repository.FindOneById(id)
	if err != nil {
		if errors.As(err, &errs.NotFoundError{}) {
			return nil, errs.HTTPError{Code: 404, Message: "Todo not found"}
		}
		return nil, err
	}
	return todo, nil
}

func (t *Todos) FindByUserId(userId int) ([]*dto.Todo, error) {
	todos, err := t.repository.FindByUserId(userId)
	if err != nil {
		return nil, err
	}
	return dto.TodosFromModels(todos), nil
}

func (t *Todos) Create(todo *dto.CreateTodo) (*models.Todo, error) {
	return t.repository.Create(todo)
}

func (t *Todos) Update(id int, todo *dto.UpdateTodo, userId int) (*models.Todo, error) {
	err := t.checkOwner(userId, id)
	if err != nil {
		return nil, err
	}
	return t.repository.Update(id, todo)
}

func (t *Todos) Delete(id, userId int) error {
	err := t.checkOwner(userId, id)
	if err != nil {
		return err
	}
	return t.repository.Delete(id)
}

func (t *Todos) checkOwner(userId, todoId int) error {
	todo, err := t.FindOneById(todoId)
	if err != nil {
		return errs.HTTPError{Code: 404, Message: "Todo not found"}
	}
	if todo.User.ID != userId {
		return errs.HTTPError{Code: 403, Message: "You are not the owner of this todo"}
	}
	return nil
}
