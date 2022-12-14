package sqlstore

import (
	"database/sql"
	"simplerest/internal/app/model"
	"simplerest/internal/app/store/errors"
)

type UserRepository struct {
	store *Store
}

func (repository *UserRepository) Create(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.BeforeCreate(); err != nil {
		return err
	}

	return repository.store.database.QueryRow(
		"INSERT INTO users (email, enctypted_password) VALUES ($1, $2) RETURNING id",
		user.Email,
		user.EncryptedPassword,
	).Scan(&user.ID)
}

func (repository *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := repository.store.database.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrorRecordNotFound
		}

		return nil, err
	}

	return u, nil
}

func (repository *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}
	if err := repository.store.database.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE id = $1",
		id,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrorRecordNotFound
		}

		return nil, err
	}

	return u, nil
}
