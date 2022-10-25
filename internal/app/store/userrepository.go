package store

import (
	"simplerest/internal/app/model"
)

type UserRepository struct {
	store *Store
}

func (repository *UserRepository) Create(user *model.User) (*model.User, error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.BeforeCreate(); err != nil {
		return nil, err
	}

	if err := repository.store.database.QueryRow(
		"INSERT INTO users (email, enctypted_password) VALUES ($1, $2) RETURNING id",
		user.Email,
		user.EncryptedPassword,
	).Scan(&user.ID); err != nil {
		return nil, err
	}

	return user, nil
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
		return nil, err
	}

	return u, nil
}
