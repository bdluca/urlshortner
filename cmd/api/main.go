package main

import (
	"github.com/brunodeluca/gophercises/urlshort/cmd/api/app"
	"github.com/brunodeluca/gophercises/urlshort/internal/platform/environment"
	"log"
	"net/http"
	"os"
)

func main() {
	env := environment.GetStringFrom(os.Getenv("GO_ENVIRONMENT"))
	dependencies := app.BuildDependencies(env)

	log.Println("listening on port 8080")

	handlers := app.Build(dependencies)
	server := http.Server{
		Addr:    ":8080",
		Handler: handlers,
	}

	log.Fatal(server.ListenAndServe())
}
