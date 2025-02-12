package main

import (
	"database/sql"

	"github.com/damedelion/url_shortener/internal/server"
	"github.com/gorilla/mux"
)

func main() {
	db := &sql.DB{}
	mux := mux.NewRouter()

	server := server.New(db, mux)
	server.Run()
}
