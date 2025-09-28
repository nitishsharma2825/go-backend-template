package repository

import (
	"context"
	"database/sql"
)

type UsersRepository interface {
	Create(context.Context) error
}

// Implement a postgres users repository
type PostgresUsersRepository struct {
	db *sql.DB
}

func (s *PostgresUsersRepository) Create(ctx context.Context) error {
	// Implement the logic to create a user in the database
	return nil
}

// Implement an in-memory users repository
type InMemoryUsersRepository struct {
	users []string // Example field to store users in memory
}

func (s *InMemoryUsersRepository) Create(ctx context.Context) error {
	// Implement the logic to create a user in memory
	s.users = append(s.users, "new user")
	return nil
}
