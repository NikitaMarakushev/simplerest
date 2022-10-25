package teststore

import (
	"simplerest/internal/app/model"
)

type UserRepository struct {
	store *teststore.Store
	user map[struct]*model.User
}

func Create(u *model.User)