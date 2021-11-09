package app

import (
	"github.com/brunodeluca/gophercises/urlshort/internal/platform/database"
	"github.com/brunodeluca/gophercises/urlshort/internal/platform/environment"
	"github.com/brunodeluca/gophercises/urlshort/internal/shortener/shorturl"
	"log"
	"os"
)

type Dependencies struct {
	ShortURLRepo shorturl.Repository
}

func BuildDependencies(env environment.Environment) *Dependencies {
	log.Printf("running as %s.\n", env.String())

	switch env {
	case environment.Production:
		mongoUser := os.Getenv("MONGO_USER")
		mongoPwd := os.Getenv("MONGO_PWD")
		mongoHost := os.Getenv("MONGO_HOST")
		mongoPort := os.Getenv("MONGO_PORT")

		redisAddr := os.Getenv("REDIS_ADDR")
		redisPwd := os.Getenv("REDIS_PWD")

		mongoClient := database.NewMongoClient(mongoUser, mongoPwd, mongoHost, mongoPort)
		redisClient := database.NewRedisClient(redisAddr, redisPwd, 0)

		shortURLRepo := shorturl.RemoteRepo{
			DB:    mongoClient.Database("test"),
			Cache: redisClient,
		}

		return &Dependencies{
			ShortURLRepo: &shortURLRepo,
		}
	case environment.Local:
		localDB := database.NewMockDB()

		mockShortURLRepo := shorturl.MockRepo{
			DB: localDB,
		}

		return &Dependencies{
			ShortURLRepo: &mockShortURLRepo,
		}
	}

	return nil
}
