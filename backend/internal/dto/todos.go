package dto

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

type TodoWithUser struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	User        User   `json:"user"`
}
