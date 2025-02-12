package http

import (
	"net/http"

	"github.com/damedelion/url_shortener/internal/shortener"
)

type delivery struct {
	usecase shortener.Usecase
}

func New(usecase shortener.Usecase) shortener.Delivery {
	return &delivery{usecase: usecase}
}

func (d *delivery) Create(w http.ResponseWriter, r *http.Request) {

}

func (d *delivery) Get(w http.ResponseWriter, r *http.Request) {

}
