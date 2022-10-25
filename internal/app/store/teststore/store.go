package teststore

import (
	"simplerest/internal/app/model"
	"simplerest/internal/app/store"
)

type Store struct {
	userRepository *UserRepository
}

func (store *Store) User() store.UserRepository {
	if store.userRepository != nil {
		return store.userRepository
	}

	store.userRepository = &UserRepository{
		store: store,
		users: make(map[string]*model.User),
	}

	return store.userRepository
}
