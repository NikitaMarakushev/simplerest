package teststore

import (
	"simplerest/internal/app/store"
)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
}

func (store *Store) User() store.UserRepository {
	if store.userRepository != nil {
		return store.userRepository
	}

	store.userRepository = &UserRepository{
		store: store,
	}

	return store.userRepository
}
