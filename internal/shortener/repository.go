package shortener

type Repository interface {
	Create(shortURL, longURL string) error
	GetShort(longURL string) (string, error)
	GetLong(shortURL string) (string, error)
}
