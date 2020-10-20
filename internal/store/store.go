package store

import (
	"database/sql"

	_ "github.com/lib/pq" // postgres driver
)

// Store ...
type Store struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository
}

// New returns Store instance with config set to given config
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open opens database given in config, checks the conncetion to db, sets store db to opened db
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DataBaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// Close ...
func (s *Store) Close() {
	s.db.Close()
}

// User ...
func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

// store.User().Create()
