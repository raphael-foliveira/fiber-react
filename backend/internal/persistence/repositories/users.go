package repositories

import (
	"github.com/raphael-foliveira/fiber-react/backend/internal/dto"
	"github.com/raphael-foliveira/fiber-react/backend/internal/models"
)

type UsersRepository interface {
	FindAll() ([]*models.User, error)
	FindById(id int) (*models.User, error)
	Create(todo *dto.CreateUser) (*models.User, error)
	Update(id int, todo *dto.UpdateUser) (*models.User, error)
	Delete(id int) error
}
