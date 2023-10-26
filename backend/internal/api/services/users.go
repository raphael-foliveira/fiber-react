package services

import "github.com/raphael-foliveira/fiber-react/backend/internal/persistence/repositories"

type Users struct {
	repository repositories.UsersRepository
}

func NewUsers(repository repositories.UsersRepository) *Users {
	return &Users{repository}
}
