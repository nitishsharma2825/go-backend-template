package repository

import (
	"context"
	"database/sql"
)

// Model for User
type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

// Repository Interface for talking to users
type UsersRepository interface {
	Create(context.Context, *User) error
}

// Implement a postgres users repository
type PostgresUsersRepository struct {
	db *sql.DB
}

func (s *PostgresUsersRepository) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (username, password, email, created_at)
		VALUES ($1, $2, $3, NOW()) RETURNING id, created_at
		`
	err := s.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Password,
		user.Email,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

// Implement an in-memory users repository
type InMemoryUsersRepository struct {
	users []User // Example field to store users in memory
}

func (s *InMemoryUsersRepository) Create(ctx context.Context, user *User) error {
	// Implement the logic to create a user in memory
	s.users = append(s.users, *user)
	return nil
}
