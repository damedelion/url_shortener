package inmemory

import (
	"sync"

	"github.com/damedelion/url_shortener/internal/shortener"
)

type repository struct {
	storage map[string]string
	mutex   *sync.Mutex
}

func New(storage map[string]string, mutex *sync.Mutex) shortener.Repository {
	return &repository{storage: storage, mutex: mutex}
}

func (r *repository) Create(shortURL, longURL string) error {
	return nil
}

func (r *repository) Get(shortURL string) (string, error) {
	return "", nil
}
