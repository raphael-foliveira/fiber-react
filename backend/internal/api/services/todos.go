package services

import (
	"github.com/raphael-foliveira/fiber-react/backend/internal/dto"
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
	return t.repository.FindOneById(id)
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

func (t *Todos) Update(id int, todo *dto.UpdateTodo) (*models.Todo, error) {
	return t.repository.Update(id, todo)
}

func (t *Todos) Delete(id int) error {
	return t.repository.Delete(id)
}
