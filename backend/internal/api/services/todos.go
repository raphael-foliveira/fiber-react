package services

import "github.com/raphael-foliveira/fiber-react/backend/internal/persistence/repositories"

type Todos struct {
	repository repositories.TodosRepository
}

func NewTodos(repository repositories.TodosRepository) *Todos {
	return &Todos{repository}
}
