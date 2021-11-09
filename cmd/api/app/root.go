package app

import (
	"github.com/brunodeluca/gophercises/urlshort/cmd/api/app/handlers"
	"github.com/brunodeluca/gophercises/urlshort/internal/shortener"
	"net/http"
)

func Build(dep *Dependencies) *http.Server {
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

	return &http.Server{
		Addr:    ":8080",
		Handler: sm,
	}
}
