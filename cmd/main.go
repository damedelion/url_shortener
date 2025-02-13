package main

import (
	"fmt"

	"github.com/damedelion/url_shortener/config"
	"github.com/damedelion/url_shortener/internal/server"
	"github.com/damedelion/url_shortener/pkg/db/postgres"
	"github.com/gorilla/mux"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		fmt.Printf("failed to read config: %v\n", err)
		return
	}

	var db any
	if cfg.DB.Type == "postgres" {
		dbConn, err := postgres.Connect(&cfg.DB)
		if err != nil {
			fmt.Printf("failed db connection: %n", err)
			return
		}
		defer dbConn.Close()

		//dbConn.Exec(migrations.CreateTableQuery)

		db = any(dbConn)
	}

	mux := mux.NewRouter()

	server := server.New(db, mux)
	server.Run(&cfg.Server)
}
