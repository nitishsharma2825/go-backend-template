package repository

import "database/sql"

type Storage struct {
	Posts PostsRepository
	Users UsersRepository
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostgresPostsRepository{db: db},
		Users: &PostgresUsersRepository{db: db},
	}
}

func NewInMemoryStorage() Storage {
	return Storage{
		Posts: &InMemoryPostsRepository{},
		Users: &InMemoryUsersRepository{},
	}
}
