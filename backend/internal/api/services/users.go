package services

import (
	"github.com/raphael-foliveira/fiber-react/backend/internal/dto"
	"github.com/raphael-foliveira/fiber-react/backend/internal/errs"
	"github.com/raphael-foliveira/fiber-react/backend/internal/models"
	"github.com/raphael-foliveira/fiber-react/backend/internal/persistence/repositories"
)

type Users struct {
	repository   repositories.UsersRepository
	todosService *Todos
}

func NewUsers(repository repositories.UsersRepository, todosService *Todos) *Users {
	return &Users{repository, todosService}
}

func (u *Users) Find() ([]*dto.User, error) {
	users, err := u.repository.Find()
	if err != nil {
		return nil, err
	}
	return dto.UsersFromModels(users), nil
}

func (u *Users) FindOneByEmail(email string) (*models.User, error) {
	return u.repository.FindOneByEmail(email)
}

func (u *Users) FindOne(id int) (*dto.User, error) {
	userModel, err := u.repository.FindOne(id)
	if err != nil {
		return nil, err
	}
	return dto.UserFromModel(userModel), nil
}

func (u *Users) FindOneWithTodos(id int) (*dto.UserWithTodos, error) {
	user, err := u.FindOne(id)
	if err != nil {
		return nil, err
	}
	todos, err := u.FindUserTodos(id)
	if err != nil {
		return nil, err
	}
	return &dto.UserWithTodos{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Todos:    todos,
	}, err
}

func (u *Users) FindUserTodos(id int) ([]*dto.Todo, error) {
	_, err := u.FindOne(id)
	if err != nil {
		return nil, err
	}
	return u.todosService.FindByUserId(id)
}

func (u *Users) Create(user *dto.CreateUser) (*models.User, error) {
	err := u.checkConflicts(user)
	if err != nil {
		return nil, err
	}
	return u.repository.Create(user)
}

func (u *Users) checkConflicts(user *dto.CreateUser) error {
	_, err := u.repository.FindOneByEmail(user.Email)
	if err == nil {
		return &errs.ConflictError{
			Message: "Email already in use",
			Field:   "email",
		}
	}
	_, err = u.repository.FindOneByUsername(user.Username)
	if err == nil {
		return &errs.ConflictError{
			Message: "Username already in use",
			Field:   "username",
		}
	}
	return nil
}

func (u *Users) Update(id int, user *dto.UpdateUser) (*models.User, error) {
	return u.repository.Update(id, user)
}

func (u *Users) Delete(id int) error {
	return u.repository.Delete(id)
}
