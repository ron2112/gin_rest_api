package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ron2112/gin_rest_api/internal/models"
)

func CreateTodo(pool *pgxpool.Pool, title string, completed bool) (*models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query string = `
		INSERT INTO todos (title, completed)
		VALUES ($1, $2)
		RETURNING id, title, completed, created_at, updated_at
	`

	var todo models.Todo

	err := pool.QueryRow(ctx, query, title, completed).Scan(
		&todo.Id,
		&todo.Title,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func GetAllTodos(pool *pgxpool.Pool) ([]models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT * FROM todos ORDER BY created_At DESC
	`

	rows, err := pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var todo models.Todo
		err = rows.Scan(
			&todo.Id,
			&todo.Title,
			&todo.Completed,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func GetTodo(pool *pgxpool.Pool, id int) (*models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT * FROM todos WHERE id = $1 LIMIT 1
	`

	var todo models.Todo
	err := pool.QueryRow(ctx, query, id).Scan(
		&todo.Id,
		&todo.Title,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func UpdateTodo(pool *pgxpool.Pool, id int, title string, completed bool) (*models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		UPDATE todos
		SET title = $2, completed = $3, updated_at = CURRENT_TIMESTAMP
		where id = $1
		RETURNING id, title, completed, created_at, updated_at
	`

	var todo models.Todo
	err := pool.QueryRow(ctx, query, id, title, completed).Scan(
		&todo.Id,
		&todo.Title,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func DeleteTodo(pool *pgxpool.Pool, id int) (*models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		DELETE FROM todos WHERE id = $1 RETURNING id, title, completed, created_at, updated_at
	`

	var todo models.Todo
	err := pool.QueryRow(ctx, query, id).Scan(
		&todo.Id,
		&todo.Title,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}