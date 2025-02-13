package server

import (
	"fmt"
	"net/http"

	"github.com/damedelion/url_shortener/config"
	"github.com/gorilla/mux"
)

type Server struct {
	db  any
	mux *mux.Router
}

func New(db any, mux *mux.Router) *Server {
	return &Server{db: db, mux: mux}
}

func (s *Server) Run(config *config.Server) {
	s.handlersRegister()

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: s.mux,
	}

	fmt.Println("Server is listening on", server.Addr)
	server.ListenAndServe()
}
