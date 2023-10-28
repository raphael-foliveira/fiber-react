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
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type UserWithTodos struct {
	ID    int     `json:"id"`
	Email string  `json:"email"`
	Todos []*Todo `json:"todos"`
}
