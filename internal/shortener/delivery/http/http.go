package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/damedelion/url_shortener/internal/dto"
	"github.com/damedelion/url_shortener/internal/shortener"
	"github.com/gorilla/mux"
)

type delivery struct {
	usecase shortener.Usecase
}

func New(usecase shortener.Usecase) shortener.Delivery {
	return &delivery{usecase: usecase}
}

func (d *delivery) Create(w http.ResponseWriter, r *http.Request) {
	var longURLDTO dto.LongURL
	err := json.NewDecoder(r.Body).Decode(&longURLDTO)
	if err != nil {
		http.Error(w, "failed to decode", http.StatusBadRequest)
		return
	}

	shortURL, err := d.usecase.Create(longURLDTO.URL)
	if err != nil {
		http.Error(w, fmt.Sprintf("usecase error: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	shortURLDTO := dto.ShortURL{URL: shortURL}
	err = json.NewEncoder(w).Encode(shortURLDTO)
	if err != nil {
		http.Error(w, "failed to encode", http.StatusInternalServerError)
		return
	}

}

func (d *delivery) Get(w http.ResponseWriter, r *http.Request) {
	shortURL := mux.Vars(r)["short_url"]
	shortURLDTO := dto.ShortURL{URL: shortURL}

	longURL, err := d.usecase.Get(shortURLDTO.URL)
	if err != nil {
		http.Error(w, fmt.Sprintf("usecase error: %v", err), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	longURLDTO := dto.LongURL{URL: longURL}
	err = json.NewEncoder(w).Encode(longURLDTO)
	if err != nil {
		http.Error(w, "failed to encode", http.StatusInternalServerError)
		return
	}
}
