package repository

import (
	"context"
	"database/sql"
)

type PostsRepository interface {
	Create(context.Context) error
}

// Implement a postgres posts repository
type PostgresPostsRepository struct {
	db *sql.DB
}

func (s *PostgresPostsRepository) Create(ctx context.Context) error {
	// Implement the logic to create a post in the database
	return nil
}

// Implement an in-memory posts repository
type InMemoryPostsRepository struct {
	posts []string // Example field to store posts in memory
}

func (s *InMemoryPostsRepository) Create(ctx context.Context) error {
	// Implement the logic to create a post in memory
	s.posts = append(s.posts, "new post")
	return nil
}
