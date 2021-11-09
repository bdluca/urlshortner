package shortener

import (
	"fmt"
	"net/url"

	"github.com/brunodeluca/gophercises/urlshort/internal/shortener/shorturl"
)

type URLShorterUseCase interface {
	RegisterURL(url string) (*shorturl.ShortURL, error)
	GetURLFromID(shortID string) (string, error)
}

type UseCaseURLShortener struct {
	ShortURLRepo shorturl.Repository
}

func (u UseCaseURLShortener) RegisterURL(url string) (*shorturl.ShortURL, error) {
	if ok := validateURL(url); !ok {
		return nil, fmt.Errorf("invalid url")
	}

	shortURL, err := u.ShortURLRepo.Save(url)
	if err != nil {
		return nil, err
	}

	return &shortURL, nil
}

func validateURL(urlString string) bool {
	u, err := url.ParseRequestURI(urlString)
	if err != nil {
		return false
	}

	return u.Host != ""
}

func (u UseCaseURLShortener) GetURLFromID(shortID string) (string, error) {
	obj, err := u.ShortURLRepo.Get(shortID)
	if err != nil {
		return "", err
	}

	return obj.URL, nil
}
