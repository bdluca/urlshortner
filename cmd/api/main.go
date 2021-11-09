package main

import (
	"github.com/brunodeluca/gophercises/urlshort/cmd/api/app"
	"github.com/brunodeluca/gophercises/urlshort/internal/platform/environment"
	"log"
	"os"
)

func main() {
	env := environment.GetStringFrom(os.Getenv("GO_ENVIRONMENT"))
	dependencies := app.BuildDependencies(env)

	log.Println("listening on port 8080")
	log.Fatal(app.Build(dependencies).ListenAndServe())
}
