package repositories

import (
	"database/sql"

	"github.com/raphael-foliveira/fiber-react/backend/internal/dto"
	"github.com/raphael-foliveira/fiber-react/backend/internal/models"
)

type UsersRepository interface {
	Find() ([]*models.User, error)
	FindOneById(id int) (*dto.UserWithTodos, error)
	Create(todo *dto.CreateUser) (*models.User, error)
	Update(id int, todo *dto.UpdateUser) (*models.User, error)
	Delete(id int) error
}

type users struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) *users {
	return &users{db}
}

func (u *users) Find() ([]*models.User, error) {
	rows, err := u.db.Query(
		"SELECT id, email FROM users",
	)
	if err != nil {
		return nil, err
	}
	return scanUsers(rows)
}

func (u *users) FindOneById(id int) (*dto.UserWithTodos, error) {
	rows, err := u.db.Query(`
		SELECT u.id, u.email, t.id, t.title, t.description, t.completed
		FROM 
			users u
		LEFT JOIN 
			todos t ON t.user_id = u.id
		WHERE u.id = $1`,
		id,
	)
	if err != nil {
		return nil, err
	}
	return scanUserWithTodos(rows)
}

func (u *users) Create(user *dto.CreateUser) (*models.User, error) {
	row := u.db.QueryRow(`
	INSERT INTO users 
		(email) 
	VALUES 
		($1) 
	RETURNING id, email`,
		user.Email)
	return scanUser(row)
}

func (u *users) Update(id int, user *dto.UpdateUser) (*models.User, error) {
	row := u.db.QueryRow(`
	UPDATE users 
	SET 
		password = $1
	WHERE id = $2
	RETURNING id, email`,
		user.Password,
		id)
	return scanUser(row)
}

func (u *users) Delete(id int) error {
	_, err := u.db.Exec(`
	DELETE FROM users 
	WHERE id = $1`,
		id)
	return err
}

func scanUsers(rows *sql.Rows) ([]*models.User, error) {
	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func scanUser(row *sql.Row) (*models.User, error) {
	var user models.User
	if err := row.Scan(&user.ID, &user.Email); err != nil {
		return nil, err
	}
	return &user, nil
}

func scanUserWithTodos(rows *sql.Rows) (*dto.UserWithTodos, error) {
	var user dto.UserWithTodos
	rows.Next()
	if err := rows.Scan(&user.ID, &user.Email); err != nil {
		return nil, err
	}
	for rows.Next() {
		var todo dto.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed); err != nil {
			return nil, err
		}
		user.Todos = append(user.Todos, &todo)
	}
	return &user, nil
}
