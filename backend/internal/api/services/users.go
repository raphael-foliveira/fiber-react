package services

import (
	"github.com/raphael-foliveira/fiber-react/backend/internal/dto"
	"github.com/raphael-foliveira/fiber-react/backend/internal/models"
	"github.com/raphael-foliveira/fiber-react/backend/internal/persistence/repositories"
)

type Users struct {
	repository repositories.UsersRepository
}

func NewUsers(repository repositories.UsersRepository) *Users {
	return &Users{repository}
}

func (u *Users) Find() ([]*models.User, error) {
	return u.repository.Find()
}

func (u *Users) FindOneByEmail(email string) (*models.User, error) {
	return u.repository.FindOneByEmail(email)
}

func (u *Users) FindOneWithTodos(id int) (*dto.UserWithTodos, error) {
	return u.repository.FindOneWithTodos(id)
}

func (u *Users) Create(user *dto.CreateUser) (*models.User, error) {
	return u.repository.Create(user)
}

func (u *Users) Update(id int, user *dto.UpdateUser) (*models.User, error) {
	return u.repository.Update(id, user)
}

func (u *Users) Delete(id int) error {
	return u.repository.Delete(id)
}
