package dto

type CreateTodo struct {
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTodo struct {
	ID int `json:"id"`
}
