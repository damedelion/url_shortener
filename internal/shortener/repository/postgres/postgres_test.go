package postgres

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestRepositoryCreate(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)

	shortURL := "abcd12345"
	longURL := "https://example.com"
	columns := []string{"id", "short_url", "long_url", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns).AddRow(
		uuid.New().String(),
		shortURL,
		longURL,
		time.Now(),
		time.Now(),
	)

	mock.ExpectQuery(CreateQuery).WithArgs(shortURL, longURL).WillReturnRows(rows)

	err = repo.Create(shortURL, longURL)
	require.NoError(t, err)
}

func TestRepositoryGetShort(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)

	longURL := "https://example.com"
	shortURL := "abcd12345"
	columns := []string{"id", "short_url", "long_url", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns).AddRow(
		uuid.New().String(),
		shortURL,
		longURL,
		time.Now(),
		time.Now(),
	)

	mock.ExpectQuery(GetShortQuery).WithArgs(longURL).WillReturnRows(rows)

	result, err := repo.GetShort(longURL)
	require.NoError(t, err)
	require.Equal(t, shortURL, result)
}

func TestRepositoryGetLong(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)

	shortURL := "abcd12345"
	longURL := "https://example.com"
	columns := []string{"id", "short_url", "long_url", "created_at", "updated_at"}
	rows := sqlmock.NewRows(columns).AddRow(
		uuid.New().String(),
		shortURL,
		longURL,
		time.Now(),
		time.Now(),
	)

	mock.ExpectQuery(GetLongQuery).WithArgs(shortURL).WillReturnRows(rows)

	result, err := repo.GetLong(shortURL)
	require.NoError(t, err)
	require.Equal(t, longURL, result)
}
