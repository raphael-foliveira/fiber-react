package repositories

import (
	"database/sql"
	"errors"

	"github.com/raphael-foliveira/fiber-react/backend/internal/dto"
	"github.com/raphael-foliveira/fiber-react/backend/internal/errs"
	"github.com/raphael-foliveira/fiber-react/backend/internal/models"
)

type UsersRepository interface {
	Find() ([]*models.User, error)
	FindOne(id int) (*models.User, error)
	FindOneByEmail(email string) (*models.User, error)
	FindOneWithTodos(id int) (*dto.UserWithTodos, error)
	Create(todo *dto.CreateUser) (*models.User, error)
	Update(id int, todo *dto.UpdateUser) (*models.User, error)
	Delete(id int) error
}

type users struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) UsersRepository {
	return &users{db}
}

func (u *users) Find() ([]*models.User, error) {
	rows, err := u.db.Query(
		"SELECT id, email, username FROM users",
	)
	if err != nil {
		return nil, err
	}
	return scanUsers(rows)
}

func (u *users) FindOne(id int) (*models.User, error) {
	row := u.db.QueryRow(`
		SELECT id, email, username, password
		FROM
			users
		WHERE id = $1`,
		id,
	)
	return scanUser(row)
}

func (u *users) FindOneByEmail(email string) (*models.User, error) {
	row := u.db.QueryRow(`
		SELECT id, email, username, password
		FROM
			users
		WHERE email = $1`,
		email,
	)
	return scanUser(row)
}

func (u *users) FindOneWithTodos(id int) (*dto.UserWithTodos, error) {
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
			(email, username, password) 
		VALUES 
			($1, $2, $3) 
		RETURNING id, email, username, password`,
		user.Email,
		user.Username,
		user.Password,
	)
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
		id,
	)
	return scanUser(row)
}

func (u *users) Delete(id int) error {
	_, err := u.db.Exec(`
		DELETE FROM users 
		WHERE id = $1`,
		id,
	)
	return err
}

func scanUsers(rows *sql.Rows) ([]*models.User, error) {
	users := []*models.User{}
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Email, &user.Username); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func scanUser(row *sql.Row) (*models.User, error) {
	var user models.User
	if err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFoundError{Message: "user not found"}
		}
		return nil, err
	}
	return &user, nil
}

func scanUserWithTodos(rows *sql.Rows) (*dto.UserWithTodos, error) {
	var user dto.UserWithTodos
	ok := rows.Next()
	if !ok {
		return nil, errs.NotFoundError{Message: "user not found"}
	}
	if err := rows.Scan(&user.ID, &user.Email, &user.Username); err != nil {
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
