package usecase

import (
	"github.com/damedelion/url_shortener/internal/shortener"
)

type usecase struct {
	repository shortener.Repository
}

func New(repository shortener.Repository) shortener.Usecase {
	return &usecase{repository: repository}
}

func (u *usecase) Create(longURL string) (string, error) {
	return "", nil
}

func (u *usecase) Get(shortURL string) (string, error) {
	return "", nil
}
