package shortener

type Usecase interface {
	Create(longURL string) (string, error)
	Get(shortURL string) (string, error)
}
