package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Handlers *Handlers
}

func NewServer(handlers *Handlers) *Server {
	return &Server{
		Handlers: handlers,
	}
}

func (s Server) StartServer() error {
	router := mux.NewRouter()

	router.Path("/employees").Methods("POST").HandlerFunc(s.Handlers.HandleNewEmployee)
	router.Path("/employees").Methods("GET").HandlerFunc(s.Handlers.HandleAllEmployee)
	router.Path("/employees").Methods("DELETE").HandlerFunc(s.Handlers.HandelDelEmployee)

	if err := http.ListenAndServe(":9999", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Server closed")
			return nil
		}
		return err
	}
	return nil
}
