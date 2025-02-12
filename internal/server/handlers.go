package server

import (
	"sync"

	"github.com/damedelion/url_shortener/internal/shortener/delivery/http"
	"github.com/damedelion/url_shortener/internal/shortener/repository/inmemory"
	"github.com/damedelion/url_shortener/internal/shortener/usecase"
)

func (s *Server) handlersRegister() {
	storage := make(map[string]string)
	mutex := &sync.Mutex{}
	repository := inmemory.New(storage, mutex)
	usecase := usecase.New(repository)
	handlers := http.New(usecase)

	s.mux.HandleFunc("/", handlers.Create).Methods("POST")
	s.mux.HandleFunc("/", handlers.Get).Methods("GET")
}
