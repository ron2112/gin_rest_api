package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ron2112/gin_rest_api/internal/models"
)

func CreateUser(pool *pgxpool.Pool, user *models.User) (*models.User, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO users (email, password)
		VALUES ($1, $2)
		RETURNING id, email, created_at, updated_at
	`

	var newUser models.User
	err := pool.QueryRow(ctx, query, user.Email, user.Password).Scan(
		&newUser.Id,
		&newUser.Email,
		&newUser.CreatedAt,
		&newUser.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func GetUserByEmail(pool *pgxpool.Pool, email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, email, created_at, updated_at
		FROM users
		WHERE email = $1
		LIMIT 1
	`

	var newUser models.User

	err := pool.QueryRow(ctx, query, email).Scan(
		&newUser.Id,
		&newUser.Email,
		&newUser.CreatedAt,
		&newUser.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func GetUserById(pool *pgxpool.Pool, id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, email, created_at, updated_at
		FROM users
		WHERE id = $1
		LIMIT 1
	`

	var newUser models.User

	err := pool.QueryRow(ctx, query, id).Scan(
		&newUser.Id,
		&newUser.Email,
		&newUser.CreatedAt,
		&newUser.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}