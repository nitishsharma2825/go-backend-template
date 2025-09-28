package repository

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

// Model for Post
type Post struct {
	ID        int64    `json:"id"`
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	UserID    int64    `json:"user_id"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
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
		INSERT INTO posts (content, title, user_id, tags, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id, created_at, updated_at
		`
	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserID,
		pq.Array(post.Tags),
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
