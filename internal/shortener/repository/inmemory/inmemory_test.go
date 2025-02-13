package inmemory_test

import (
	"sync"
	"testing"

	"github.com/damedelion/url_shortener/internal/shortener/repository/inmemory"
	"github.com/stretchr/testify/assert"
)

func TestRepository_CreateAndGet(t *testing.T) {
	storageLongToShort := make(map[string]string)
	storageShortToLong := make(map[string]string)
	repo := inmemory.New(storageLongToShort, storageShortToLong)

	t.Run("Successful create and get", func(t *testing.T) {
		shortURL := "short123"
		longURL := "https://example.com"

		err := repo.Create(shortURL, longURL)
		assert.NoError(t, err)

		gotShort, err := repo.GetShort(longURL)
		assert.NoError(t, err)
		assert.Equal(t, shortURL, gotShort)

		gotLong, err := repo.GetLong(shortURL)
		assert.NoError(t, err)
		assert.Equal(t, longURL, gotLong)
	})

	t.Run("GetShort returns error if not found", func(t *testing.T) {
		_, err := repo.GetShort("https://notfound.com")
		assert.Error(t, err)
		assert.Equal(t, "not found", err.Error())
	})

	t.Run("GetLong returns error if not found", func(t *testing.T) {
		_, err := repo.GetLong("unknown")
		assert.Error(t, err)
		assert.Equal(t, "not found", err.Error())
	})
}

func TestRepository_ConcurrentAccess(t *testing.T) {
	storageLongToShort := make(map[string]string)
	storageShortToLong := make(map[string]string)
	repo := inmemory.New(storageLongToShort, storageShortToLong)

	var wg sync.WaitGroup
	shortURL := "short123"
	longURL := "https://example.com"

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			repo.Create(shortURL, longURL)
			wg.Done()
		}()
	}
	wg.Wait()

	gotShort, err := repo.GetShort(longURL)
	assert.NoError(t, err)
	assert.Equal(t, shortURL, gotShort)
}
