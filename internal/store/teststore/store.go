package teststore

import (
	"github.com/elmira-aliyeva/go-rest-api/internal/model"
	"github.com/elmira-aliyeva/go-rest-api/internal/store"
)

// Store ...
type Store struct {
	userRepository *UserRepository
}

// New returns Store instance with config set to given config
func New() *Store {
	return &Store{}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}

// store.User().Create()
