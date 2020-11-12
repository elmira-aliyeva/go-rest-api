package store

import (
	"github.com/elmira-aliyeva/go-rest-api/internal/model"
)

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
