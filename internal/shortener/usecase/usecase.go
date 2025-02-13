package usecase

import (
	"fmt"
	"math/rand"

	"github.com/damedelion/url_shortener/internal/shortener"
	"github.com/damedelion/url_shortener/pkg/base63"
	"github.com/damedelion/url_shortener/pkg/math"
)

type usecase struct {
	repository shortener.Repository
}

func New(repository shortener.Repository) shortener.Usecase {
	return &usecase{repository: repository}
}

func (u *usecase) Create(longURL string) (string, error) {
	shortURL, err := u.repository.GetShort(longURL)
	if err == nil {
		return shortURL, nil
	}

	maxCount := math.PowInt64(int64(63), int64(10))
	id := rand.Int63n(maxCount)
	res, _ := base63.ToBase63(id, 10)
	shortURL = string(res)

	err = u.repository.Create(shortURL, longURL)
	if err != nil {
		return "", fmt.Errorf("repository error: %v", err)
	}

	return shortURL, nil
}

func (u *usecase) Get(shortURL string) (string, error) {
	longURL, err := u.repository.GetLong(shortURL)
	if err != nil {
		return "", fmt.Errorf("%s not found", shortURL)
	}

	return longURL, nil
}
