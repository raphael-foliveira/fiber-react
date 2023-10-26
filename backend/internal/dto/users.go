package dto

type CreateUser struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type UpdateUser struct {
	Password string `json:"password"`
}

type User struct {
	ID    int
	Email string
}

type UserWithTodos struct {
	ID    int
	Email string
	Todos []*Todo
}
