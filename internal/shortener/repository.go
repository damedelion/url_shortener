package shortener

type Repository interface {
	Create(shortURL, longURL string) error
	Get(shortURL string) (string, error)
}
