package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	config   *Config
	database *sql.DB
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}
func (store *Store) Open() error {
	database, err := sql.Open("postgres", store.config.DatabaseURL)

	if err != nil {
		return err
	}

	if err := database.Ping(); err != nil {
		return err
	}

	store.database = database

	return nil
}

func (store *Store) Close() {
	store.database.Close()
}
