package server

import (
	"database/sql"

	"github.com/damedelion/url_shortener/internal/shortener"
	"github.com/damedelion/url_shortener/internal/shortener/delivery/http"
	"github.com/damedelion/url_shortener/internal/shortener/repository/inmemory"
	"github.com/damedelion/url_shortener/internal/shortener/repository/postgres"
	"github.com/damedelion/url_shortener/internal/shortener/usecase"
)

func (s *Server) handlersRegister() {
	var repository shortener.Repository

	switch db := s.db.(type) {
	case *sql.DB:
		repository = postgres.New(db)
	default:
		storageLongToShort := make(map[string]string)
		storageShortToLong := make(map[string]string)
		repository = inmemory.New(storageLongToShort, storageShortToLong)
	}

	usecase := usecase.New(repository)
	handlers := http.New(usecase)

	s.mux.HandleFunc("/", handlers.Create).Methods("POST")
	s.mux.HandleFunc("/{short_url}", handlers.Get).Methods("GET")
}
