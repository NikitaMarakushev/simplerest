package apiserver

import (
	"database/sql"
	"net/http"
	"simplerest/internal/app/store/sqlstore"
)

const driverName = "postgres"

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)

	if err != nil {
		return err
	}

	defer db.Close()
	store := sqlstore.New(db)
	srv := newServer(store)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open(driverName, databaseURL)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
