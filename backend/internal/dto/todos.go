package dto

import (
	"database/sql"
	"time"

	"github.com/raphael-foliveira/fiber-react/backend/internal/models"
)

type CreateTodo struct {
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTodo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
}

func TodoFromModel(todoModel *models.Todo) *Todo {
	return &Todo{
		ID:          todoModel.ID,
		Title:       todoModel.Title,
		Description: todoModel.Description,
		Completed:   todoModel.Completed,
		CreatedAt:   todoModel.CreatedAt.Time,
		CompletedAt: todoModel.CompletedAt.Time,
	}
}

func TodosFromModels(todoModels []*models.Todo) []*Todo {
	todos := []*Todo{}
	for _, todoModel := range todoModels {
		todos = append(todos, &Todo{
			ID:          todoModel.ID,
			Title:       todoModel.Title,
			Description: todoModel.Description,
			Completed:   todoModel.Completed,
			CreatedAt:   todoModel.CreatedAt.Time,
			CompletedAt: todoModel.CompletedAt.Time,
		})
	}
	return todos
}

type TodoWithUser struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Completed   bool         `json:"completed"`
	User        User         `json:"user"`
	CreatedAt   sql.NullTime `json:"created_at"`
	CompletedAt sql.NullTime `json:"completed_at"`
}
