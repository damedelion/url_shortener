package inmemory

import (
	"fmt"
	"sync"

	"github.com/damedelion/url_shortener/internal/shortener"
)

type repository struct {
	storageLongToShort map[string]string
	storageShortToLong map[string]string
	mutex              sync.Mutex
}

func New(st1, st2 map[string]string) shortener.Repository {
	return &repository{storageLongToShort: st1, storageShortToLong: st2, mutex: sync.Mutex{}}
}

func (r *repository) Create(shortURL, longURL string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.storageLongToShort[longURL] = shortURL
	r.storageShortToLong[shortURL] = longURL
	return nil
}

func (r *repository) GetShort(longURL string) (string, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	res, ok := r.storageLongToShort[longURL]

	if !ok {
		return "", fmt.Errorf("not found")
	}
	return res, nil
}

func (r *repository) GetLong(shortURL string) (string, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	res, ok := r.storageShortToLong[shortURL]

	if !ok {
		return "", fmt.Errorf("not found")
	}
	return res, nil
}
