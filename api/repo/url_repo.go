package repo

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

type UrlRepo interface {
	CreateShortUrl(shortUrl, originalUrl string, expirationTimeHours time.Duration) (string, error)
	GetOriginalUrl(string) (string, error)
}

type UrlRepoImpl struct {
	db *redis.Client
}

func NewUrlRepoImpl(db *redis.Client) *UrlRepoImpl {
	return &UrlRepoImpl{
		db: db,
	}
}

func (u *UrlRepoImpl) CreateShortUrl(shortUrl, originalUrl string, expirationTimeHours time.Duration) (string, error) {
	isUrlExists := u.isOriginalUrlAlreadyExists(originalUrl)

	if isUrlExists {
		return shortUrl, errors.New("url already exists!")
	}

	err := u.db.HSet(dbCtx, "Urls", shortUrl, originalUrl).Err()
	if err != nil {
		return "", err
	}

	return shortUrl, nil
}

func (u *UrlRepoImpl) GetOriginalUrl(code string) (string, error) {
	urlsHash, _ := u.db.HGetAll(dbCtx, "Urls").Result()

	for k, v := range urlsHash {
		if k == code {
			return v, nil
		}
	}

	return "", errors.New("does not exist")
}

func (u *UrlRepoImpl) isOriginalUrlAlreadyExists(originalUrl string) bool {
	urlsHash, _ := u.db.HGetAll(dbCtx, "Urls").Result()

	for _, v := range urlsHash {
		if v == originalUrl {
			return true
		}
	}

	return false
}
