package app

import (
	"github.com/brunodeluca/gophercises/urlshort/cmd/api/app/handlers"
	"github.com/brunodeluca/gophercises/urlshort/internal/shortener"
	"net/http"
)

func Build(dep *Dependencies) *http.ServeMux {
	// use cases
	shorterUseCase := shortener.UseCaseURLShortener{
		ShortURLRepo: dep.ShortURLRepo,
	}

	// handlers
	shortURLHandler := handlers.ShortURLHandler{
		ShorterUseCase: shorterUseCase,
	}

	sm := http.NewServeMux()
	sm.HandleFunc("/register", shortURLHandler.HandleRegistry())
	sm.HandleFunc("/", shortURLHandler.HandleRedirect())

	return sm
}
