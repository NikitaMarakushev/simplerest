package model_test

import (
	"github.com/stretchr/testify/assert"
	"simplerest/internal/app/model"
	"testing"
)

func TestUser_BeforeCreate(testing *testing.T) {
	user := model.TestUser(testing)
	assert.NoError(testing, user.BeforeCreate())
	assert.NotEmpty(testing, user.EncryptedPassword)
}
