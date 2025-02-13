package repository

import (
	"database/sql"

	"github.com/damedelion/url_shortener/internal/shortener"
)

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) shortener.Repository {
	return &repository{db: db}
}

func (r *repository) Create(shortURL, longURL string) error {
	return nil
}

func (r *repository) GetShort(longURL string) (string, error) {
	return "", nil
}

func (r *repository) GetLong(shortURL string) (string, error) {
	return "", nil
}
