package sqlstore

import (
	"database/sql"

	"github.com/elmira-aliyeva/go-rest-api/internal/store"

	"github.com/elmira-aliyeva/go-rest-api/internal/model"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create inserts into users table email and encrypted password
func (r *UserRepository) Create(u *model.User) error {

	if err := u.Validate(); err != nil {
		return err
	}

	// BeforeCreate encrypts password, sets EncryptedPassword
	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)

}

// FindByEmail looks for a user in the users table by email and returns the user
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow("SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNoFound
		}

		return nil, err
	}

	return u, nil
}
