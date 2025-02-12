package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	db  *sql.DB
	mux *mux.Router
}

func New(db *sql.DB, mux *mux.Router) *Server {
	return &Server{db: db, mux: mux}
}

func (s *Server) Run() {
	s.handlersRegister()

	server := http.Server{
		Addr:    ":3000",
		Handler: s.mux,
	}

	fmt.Println("Server is listening on", server.Addr)
	server.ListenAndServe()
}
