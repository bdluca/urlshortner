package shorturl

import (
	"context"
	"fmt"
	"github.com/brunodeluca/gophercises/urlshort/internal/platform/sequence"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type RemoteRepo struct {
	DB    *mongo.Database
	Cache *redis.Client
}

func (r *RemoteRepo) Save(url string) (ShortURL, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	id := sequence.Generate()
	_, err := r.DB.Collection("shorturl").InsertOne(ctx, bson.D{
		{"id", id},
		{"url", url},
	})

	if err != nil {
		return ShortURL{}, err
	}

	return ShortURL{ID: id, URL: url}, nil
}

func (r *RemoteRepo) Get(shortID string) (ShortURL, error) {
	shortURL, err := findInCache(shortID, r.Cache)
	if err == nil {
		log.Println("returning obj from cache")
		return shortURL, nil
	}

	shortURL, err = findInDB(shortID, r.DB)
	if err != nil {
		return ShortURL{}, err
	}

	if err := saveInCache(shortURL, r.Cache); err != nil {
		log.Printf("[WARNING] could not save obj into cache: %v\n", err)
	}

	log.Println("returning obj from db")
	return shortURL, nil
}

func findInCache(shortID string, rdb *redis.Client) (ShortURL, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	url, err := rdb.Get(ctx, shortID).Result()
	if err == redis.Nil {
		return ShortURL{}, fmt.Errorf("key not found")
	}

	if err != nil {
		return ShortURL{}, err
	}

	return ShortURL{ID: shortID, URL: url}, nil
}

func saveInCache(shortURL ShortURL, rdb *redis.Client) error {
	log.Printf("saving obj into cache: %v\n", shortURL)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := rdb.Set(ctx, shortURL.ID, shortURL.URL, 0).Err()
	return err
}

func findInDB(shortID string, db *mongo.Database) (ShortURL, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := db.Collection("shorturl").FindOne(ctx, bson.D{{"id", shortID}})
	if result.Err() != nil {
		return ShortURL{}, fmt.Errorf("error querying url: %v", result.Err())
	}

	var shortURL ShortURL
	if err := result.Decode(&shortURL); err != nil {
		return ShortURL{}, fmt.Errorf("error decoding query result: %v", result.Err())
	}

	return shortURL, nil
}
