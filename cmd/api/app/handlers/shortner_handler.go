package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/brunodeluca/gophercises/urlshort/internal/shortener"
	"net/http"
	"strings"
)

type ShortURLHandler struct {
	ShorterUseCase shortener.URLShorterUseCase
}

func (h ShortURLHandler) HandleRegistry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortURL, err := h.ShorterUseCase.RegisterURL(r.URL.Query().Get("url"))
		if err != nil {
			fmt.Printf("unexpected error registering url: %v\n", err)
			http.Error(w, "unexpected error registering url", http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(shortURL)
		if err != nil {
			fmt.Printf("unexpected error encoding response json: %v\n", err)
			http.Error(w, "unexpected error encoding response json", http.StatusInternalServerError)
		}
	}
}

func (h ShortURLHandler) HandleRedirect() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.ReplaceAll(r.URL.Path, "/", "")
		url, err := h.ShorterUseCase.GetURLFromID(id)
		if err != nil {
			fmt.Printf("unexpected error getting url: %v\n", err)
			http.NotFound(w, r)
			return
		}

		http.Redirect(w, r, url, http.StatusFound)
	}
}
