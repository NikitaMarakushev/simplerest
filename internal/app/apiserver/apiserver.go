package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"simplerest/internal/app/store/sqlstore"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *sqlstore.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (server *APIServer) Start() error {

	if err := server.configureLogger(); err != nil {
		return err
	}

	server.configureRouter()

	if err := server.configureStore(); err != nil {
		return err
	}

	server.logger.Info("API server started successfully!")

	return http.ListenAndServe(server.config.BindAddr, server.router)
}

func (server *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(server.config.LogLevel)

	if err != nil {
		return err
	}

	server.logger.SetLevel(level)

	return nil
}

func (server *APIServer) configureRouter() {
	server.router.HandleFunc("/hello", server.handleHello())
}

func (server *APIServer) handleHello() http.HandlerFunc {

	return func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "hello")
	}
}

func (server *APIServer) configureStore() error {
	storeInstance := sqlstore.New(server.config.Store)

	if err := storeInstance.Open(); err != nil {
		return err
	}

	server.store = storeInstance

	return nil
}
