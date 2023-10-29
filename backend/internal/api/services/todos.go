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

func (t *Todos) Find() ([]*models.Todo, error) {
	return t.repository.Find()
}

func (t *Todos) FindOneById(id int) (*dto.TodoWithUser, error) {
	return t.repository.FindOneById(id)
}

func (t *Todos) FindByUserId(userId int) ([]*models.Todo, error) {
	return t.repository.FindByUserId(userId)
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
