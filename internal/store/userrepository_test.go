package store_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elmira-aliyeva/go-rest-api/internal/model"
	"github.com/elmira-aliyeva/go-rest-api/internal/store"
)

func TestUserRepository_Create(t *testing.T) {
	s, tearDown := store.TestStore(t, databaseURL)
	defer tearDown("users")

	u, err := s.User().Create(&model.User{
		Email: "user@example.org",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, tearDown := store.TestStore(t, databaseURL)
	defer tearDown("users")

	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	u, err := s.User().Create(&model.User{
		Email: "user@example.org",
	})

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
