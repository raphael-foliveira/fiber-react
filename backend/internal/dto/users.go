package dto

import "github.com/raphael-foliveira/fiber-react/backend/internal/models"

type CreateUser struct {
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type UpdateUser struct {
	Password string `json:"password"`
}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func UsersFromModels(userModels []*models.User) []*User {
	users := []*User{}
	for _, userModel := range userModels {
		users = append(users, &User{
			ID:       userModel.ID,
			Email:    userModel.Email,
			Username: userModel.Username,
		})
	}
	return users
}

type UserWithTodos struct {
	ID       int     `json:"id"`
	Email    string  `json:"email"`
	Username string  `json:"username"`
	Todos    []*Todo `json:"todos"`
}
