package shorturl

type Repository interface {
	Save(url string) (ShortURL, error)
	Get(shortID string) (ShortURL, error)
}

type ShortURL struct {
	ID  string
	URL string
}
