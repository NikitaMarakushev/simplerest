package store

import (
	"simplerest/internal/app/model"
)

type UserRepository struct {
	store *Store
}

func (repository *UserRepository) Create(user *model.User) (*model.User, error) {
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
	return nil, nil
}
