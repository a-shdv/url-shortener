package repo

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

type UrlRepo interface {
	CreateShortUrl(shortUrl, originalUrl string, expirationTimeHours time.Duration) (string, error)
	//GetOriginalUrl(string) (string, error)
}

type UrlRepoImpl struct {
	db *redis.Client
}

func NewUrlRepoImpl(db *redis.Client) *UrlRepoImpl {
	return &UrlRepoImpl{
		db: db,
	}
}

func (u *UrlRepoImpl) CreateShortUrl(originalUrl, shortUrl string, expirationTimeHours time.Duration) (string, error) {
	shortUrlDb := u.getShortUrlByOriginalUrl(originalUrl)
	if shortUrlDb != "" {
		return shortUrlDb, errors.New("url is already in database")
	}
	err := u.db.Set(dbCtx, originalUrl, shortUrl, expirationTimeHours).Err()
	if err != nil {
		return "", err
	}
	return shortUrl, nil
}

func (u *UrlRepoImpl) getShortUrlByOriginalUrl(originalUrl string) string {
	shortUrlDb, _ := u.db.Get(dbCtx, originalUrl).Result()
	return shortUrlDb
}
