package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/sessions"

	"github.com/elmira-aliyeva/go-rest-api/internal/store/sqlstore"
)

// Start checks the db, starts listening on the given port
func Start(config *Config) error {
	// opens db and checks the connection
	db, err := newDB(config.DataBaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	// returns Store instance with db set to given db
	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))

	// returns server instance with gorilla/mux router, store set to given store,
	// sessionStore set to the one given, configures the router
	srv := newServer(store, sessionStore)

	return http.ListenAndServe(config.BindAddr, srv)
}

// opens db and checks the connection
func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
