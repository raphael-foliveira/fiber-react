package repositories

import (
	"database/sql"

	"github.com/raphael-foliveira/fiber-react/backend/internal/dto"
	"github.com/raphael-foliveira/fiber-react/backend/internal/models"
)

type TodosRepository interface {
	Find() ([]*models.Todo, error)
	FindOneById(id int) (*dto.TodoWithUser, error)
	Create(todo *dto.CreateTodo) (*models.Todo, error)
	Update(id int, todo *dto.UpdateTodo) (*models.Todo, error)
	Delete(id int) error
}

type todos struct {
	db *sql.DB
}

func NewTodos(db *sql.DB) TodosRepository {
	return &todos{db}
}

func (t *todos) Find() ([]*models.Todo, error) {
	rows, err := t.db.Query(
		"SELECT id, title, description, completed, user_id FROM todos",
	)
	if err != nil {
		return nil, err
	}
	return scanTodos(rows)
}

func (t *todos) FindOneById(id int) (*dto.TodoWithUser, error) {
	row := t.db.QueryRow(`
		SELECT t.id, t.title, t.description, t.completed, u.id, u.email
		FROM 
			todos t
		JOIN 
			users u ON t.user_id = u.id
		WHERE id = $1`,
		id,
	)
	return scanTodoWithUser(row)
}

func (t *todos) Create(todo *dto.CreateTodo) (*models.Todo, error) {
	row := t.db.QueryRow(`
		INSERT INTO todos 
			(title, description, completed, user_id) 
		VALUES 
			($1, $2, $3, $4) 
		RETURNING id, title, description, completed, user_id`,
		todo.Title,
		todo.Description,
		false,
		todo.UserID,
	)
	return scanTodo(row)
}

func (t *todos) Update(id int, todo *dto.UpdateTodo) (*models.Todo, error) {
	row := t.db.QueryRow(`
		UPDATE todos 
		SET 
			title = $1, 
			description = $2, 
			completed = $3 
		WHERE 
			id = $4 
		RETURNING id, title, description, completed, user_id`,
		todo.Title,
		todo.Description,
		todo.Completed,
		id,
	)
	return scanTodo(row)
}

func (t *todos) Delete(id int) error {
	_, err := t.db.Exec("DELETE FROM todos WHERE id = $1", id)
	return err
}

func scanTodos(rows *sql.Rows) ([]*models.Todo, error) {
	todos := []*models.Todo{}
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.UserID); err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}
	return todos, nil
}

func scanTodo(row *sql.Row) (*models.Todo, error) {
	var todo models.Todo
	if err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.UserID); err != nil {
		return nil, err
	}
	return &todo, nil
}

func scanTodoWithUser(row *sql.Row) (*dto.TodoWithUser, error) {
	var todo dto.TodoWithUser
	if err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.User.ID, &todo.User.Email); err != nil {
		return nil, err
	}
	return &todo, nil
}
