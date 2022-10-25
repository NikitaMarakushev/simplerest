package store_test

import (
	"github.com/stretchr/testify/assert"
	"simplerest/internal/app/model"
	"simplerest/internal/app/store"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	user, err := s.User().Create(model.TestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	testStore, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "user@example.org"
	_, err := testStore.User().FindByEmail(email)
	assert.Error(t, err)

	user := model.TestUser(t)
	user.Email = email
	testStore.User().Create(user)
	user, err = testStore.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}
