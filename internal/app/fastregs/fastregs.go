package fastregs

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Fastregs struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *Fastregs {
	return &Fastregs{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *Fastregs) Start() error {
	s.logger.Info("starting fast registers server")

	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Fastregs) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *Fastregs) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *Fastregs) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
