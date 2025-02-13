package postgres

import (
	"database/sql"
	"fmt"

	"github.com/damedelion/url_shortener/internal/entities"
	"github.com/damedelion/url_shortener/internal/shortener"
)

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) shortener.Repository {
	return &repository{db: db}
}

func (r *repository) Create(shortURL, longURL string) error {
	url := &entities.URL{}

	row := r.db.QueryRow(CreateQuery, shortURL, longURL)

	if err := row.Scan(
		&url.ID,
		&url.ShortURL,
		&url.LongURL,
		&url.CreatedAt,
		&url.UpdatedAt,
	); err != nil {
		return fmt.Errorf("failed to scan from Create query")
	}

	return nil
}

func (r *repository) GetShort(longURL string) (string, error) {
	url := &entities.URL{}

	row := r.db.QueryRow(GetShortQuery, longURL)

	if err := row.Scan(
		&url.ID,
		&url.ShortURL,
		&url.LongURL,
		&url.CreatedAt,
		&url.UpdatedAt,
	); err != nil {
		return "", fmt.Errorf("failed to scan from GetShort query")
	}

	return url.ShortURL, nil
}

func (r *repository) GetLong(shortURL string) (string, error) {
	url := &entities.URL{}

	row := r.db.QueryRow(GetLongQuery, shortURL)

	if err := row.Scan(
		&url.ID,
		&url.ShortURL,
		&url.LongURL,
		&url.CreatedAt,
		&url.UpdatedAt,
	); err != nil {
		return "", fmt.Errorf("failed to scan from GetLong query")
	}

	return url.LongURL, nil
}
