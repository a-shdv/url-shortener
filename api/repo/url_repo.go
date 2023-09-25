package repo

import (
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type UrlRepo interface {
	CreateShortUrl(shortUrl, originalUrl string, expirationTimeHours time.Duration)
}

type UrlRepoImpl struct {
	db *redis.Client
}

func NewUrlRepoImpl(db *redis.Client) *UrlRepoImpl {
	return &UrlRepoImpl{
		db: db,
	}
}

func (r *UrlRepoImpl) CreateShortUrl(shortUrl, originalUrl string, expirationTimeHours time.Duration) {
	// check if the user provided short is already in use
	shortUrlDb, _ := r.db.Get(dbCtx, shortUrl).Result()
	if shortUrlDb != "" {
		log.Fatalf("short url is already in use!")
	}

	err := r.db.Set(dbCtx, shortUrl, originalUrl, expirationTimeHours*3600*time.Second).Err()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
