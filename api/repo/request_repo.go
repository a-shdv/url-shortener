package repo

import (
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type RequestRepo interface {
	CreateShortUrl(shortUrl, originalUrl string, expirationTimeHours time.Duration)
}

type RequestRepoImpl struct {
	db *redis.Client
}

func NewRequestRepoImpl(db *redis.Client) *RequestRepoImpl {
	return &RequestRepoImpl{
		db: db,
	}
}

func (r *RequestRepoImpl) CreateShortUrl(shortUrl, originalUrl string, expirationTimeHours time.Duration) {
	// check if the user provided short is already in use
	val, _ := r.db.Get(dbCtx, shortUrl).Result()
	if val != "" {
		log.Fatalf("short url is already in use!")
	}

	err := r.db.Set(dbCtx, shortUrl, originalUrl, expirationTimeHours*3600*time.Second).Err()
	if err != nil {
		log.Fatal(err.Error())
	}
}
