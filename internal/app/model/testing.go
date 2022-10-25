package model

import "testing"

func TestUser(testing *testing.T) *User {
	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}
