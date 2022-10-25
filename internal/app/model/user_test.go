package model_test

import (
	"github.com/stretchr/testify/assert"
	"simplerest/internal/app/model"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		user    func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			user: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "with encrypted password",
			user: func() *model.User {
				user := model.TestUser(t)
				user.Password = ""
				user.EncryptedPassword = "encryptedPassword"

				return user
			},
			isValid: true,
		},
		{
			name: "empty email",
			user: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""

				return u
			},
			isValid: false,
		},
		{
			name: "invalid email",
			user: func() *model.User {
				u := model.TestUser(t)
				u.Email = "Invalid"

				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			user: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""

				return u
			},
			isValid: false,
		},
		{
			name: "short password",
			user: func() *model.User {
				u := model.TestUser(t)
				u.Password = "1"

				return u
			},
			isValid: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.isValid {
				assert.NoError(t, testCase.user().Validate())
			} else {
				assert.Error(t, testCase.user().Validate())
			}
		})
	}
}

func TestUser_BeforeCreate(t *testing.T) {
	user := model.TestUser(t)
	assert.NoError(t, user.BeforeCreate())
	assert.NotEmpty(t, user.EncryptedPassword)
}
