package sqlstore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elmira-aliyeva/go-rest-api/internal/model"
	"github.com/elmira-aliyeva/go-rest-api/internal/store"
	"github.com/elmira-aliyeva/go-rest-api/internal/store/sqlstore"
)

func TestUserRepository_Create(t *testing.T) {
	db, tearDown := sqlstore.TestDB(t, databaseURL)
	defer tearDown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, tearDown := sqlstore.TestDB(t, databaseURL)
	defer tearDown("users")

	s := sqlstore.New(db)
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNoFound.Error())

	u := model.TestUser(t)
	u.Email = email

	s.User().Create(u)

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
