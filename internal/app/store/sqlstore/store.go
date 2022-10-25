package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	database       *sql.DB
	userRepository *UserRepository
}

func New(database *sql.DB) *Store {
	return &Store{
		database: database,
	}
}

func (store *Store) User() *UserRepository {
	if store.userRepository != nil {
		return store.userRepository
	}

	store.userRepository = &UserRepository{
		store: store,
	}

	return store.userRepository
}
