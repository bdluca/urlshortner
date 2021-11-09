package shorturl

import (
	"fmt"
	"github.com/brunodeluca/gophercises/urlshort/internal/platform/database"
	"github.com/brunodeluca/gophercises/urlshort/internal/platform/sequence"
)

type MockRepo struct {
	DB *database.LocalDB
}

func (r *MockRepo) Save(url string) (ShortURL, error) {
	id := sequence.Generate()
	r.DB.SaveItem(id, url)
	return ShortURL{
		ID:  id,
		URL: url,
	}, nil
}

func (r *MockRepo) Get(shortID string) (ShortURL, error) {
	url := r.DB.GetItem(shortID)
	if url == "" {
		return ShortURL{}, fmt.Errorf("short url not found")
	}

	return ShortURL{
		ID:  shortID,
		URL: url,
	}, nil
}
