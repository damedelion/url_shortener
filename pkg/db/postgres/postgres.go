package postgres

import (
	"database/sql"
	"fmt"

	"github.com/damedelion/url_shortener/config"
	_ "github.com/lib/pq"
)

func Connect(config *config.DB) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s timezone=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode, config.Timezone)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db, err: %v\n", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to verify db connection, err: %v\n", err)
	}

	return db, nil
}
