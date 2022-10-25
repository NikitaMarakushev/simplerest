package teststore

import (
	"simplerest/internal/app/model"
	"simplerest/internal/app/store/erro"
)

type UserRepository struct {
	store *Store
	users map[string]*model.User
}

func (r *UserRepository) Create(u *model.User) erros {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.users[u.Email] = u
	u.ID = len(r.users)

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, erros) {
	u, ok := r.users[email]

	if !ok {
		return nil, erros.
	}

	return u, nil
}
