package dto

import "github.com/raphael-foliveira/fiber-react/backend/internal/models"

type CreateTodo struct {
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTodo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func TodosFromModels(todoModels []*models.Todo) []*Todo {
	todos := []*Todo{}
	for _, todoModel := range todoModels {
		todos = append(todos, &Todo{
			ID:          todoModel.ID,
			Title:       todoModel.Title,
			Description: todoModel.Description,
			Completed:   todoModel.Completed,
		})
	}
	return todos
}

type TodoWithUser struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	User        User   `json:"user"`
}
