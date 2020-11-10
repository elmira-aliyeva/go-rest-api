package store

import (
	"github.com/elmira-aliyeva/go-rest-api/internal/model"
)

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
