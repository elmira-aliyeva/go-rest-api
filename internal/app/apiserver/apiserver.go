package apiserver

import (
	"io"
	"net/http"

	"github.com/elmira-aliyeva/go-rest-api/internal/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// New returns an instance of APIServer struct with config set to given config,
// sets logrus logger as logger, sets gorilla/mux router as server router
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start configures logger, router, store and starts listening on the given port
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	// creates new Store instance with server's store config (dbURL), opens db, sets store db to this db, sets server store
	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("starting api server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}
}

// creates new Store instance with server's store config (dbURL), opens db, sets store db to this db, sets server's store
func (s *APIServer) configureStore() error {

	// Store - config (databaseURL), db, userRepository
	// New returns Store instance with config set to given config
	st := store.New(s.config.Store)

	// Open opens database given in config, checks the conncetion to db, sets store db to opened db
	if err := st.Open(); err != nil {
		return err
	}

	// set server store to this store
	s.store = st
	return nil
}
