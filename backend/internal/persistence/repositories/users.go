package repositories

import (
	"database/sql"

	"github.com/raphael-foliveira/fiber-react/backend/internal/dto"
	"github.com/raphael-foliveira/fiber-react/backend/internal/models"
)

type UsersRepository interface {
	Find() ([]*models.User, error)
	FindOneById(id int) (*dto.UserWithTodos, error)
	Create(todo *dto.CreateUser) (*models.User, error)
	Update(id int, todo *dto.UpdateUser) (*models.User, error)
	Delete(id int) error
}

type users struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) *users {
	return &users{db}
}
