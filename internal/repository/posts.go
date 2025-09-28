package repository

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

// Model for Post
type Post struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	UserID    int64  `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Repository interface for talking to posts
type PostsRepository interface {
	Create(context.Context, *Post) error
}

// Implement a postgres posts repository
type PostgresPostsRepository struct {
	db *sql.DB
}

func (s *PostgresPostsRepository) Create(ctx context.Context, post *Post) error {
	query := `
		INSERT INTO posts (title, user_id, content, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id, created_at, updated_at
		`
	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Title,
		post.UserID,
		post.Content,
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

// Implement an in-memory posts repository
type InMemoryPostsRepository struct {
	posts []Post // Example field to store posts in memory
}

func (s *InMemoryPostsRepository) Create(ctx context.Context, post *Post) error {
	// Implement the logic to create a post in memory
	s.posts = append(s.posts, *post)
	return nil
}
