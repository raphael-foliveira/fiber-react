package repositories

import (
	"github.com/raphael-foliveira/fiber-react/backend/internal/dto"
	"github.com/raphael-foliveira/fiber-react/backend/internal/models"
)

type TodosRepository interface {
	FindAll() ([]*models.Todo, error)
	FindById(id int) (*models.Todo, error)
	Create(todo *dto.CreateTodo) (*models.Todo, error)
	Update(id int, todo *models.Todo) (*models.Todo, error)
	Delete(id int) error
}
