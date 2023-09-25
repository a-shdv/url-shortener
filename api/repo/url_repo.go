package repo

import (
	"github.com/a-shdv/url-shortener/api/helper"
	"github.com/go-redis/redis/v8"
	"log"
)

type UrlRepo interface {
	CreateShortUrl(shortUrl, originalUrl string) (string, error)
	GetOriginalUrlByCode(string) string
}

type UrlRepoImpl struct {
	db *redis.Client
}

func NewUrlRepoImpl(db *redis.Client) *UrlRepoImpl {
	return &UrlRepoImpl{
		db: db,
	}
}

func (u *UrlRepoImpl) CreateShortUrl(shortUrl, originalUrl string) (string, error) {
	// get existing data from db
	urlsHash, err := u.db.HGetAll(dbCtx, "Urls").Result()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// shortUrl already exists in db
	shortUrlDb := getShortUrl(urlsHash, originalUrl)
	if shortUrlDb != "" {
		return shortUrlDb, nil
	}

	// encrypt short url
	shortUrl = helper.GenerateRandomChar()

	// add key-value pair to 'Urls' hashtable
	err = u.db.HSet(dbCtx, "Urls", shortUrl, originalUrl).Err()
	if err != nil {
		return "", err
	}

	return shortUrl, nil
}

func (u *UrlRepoImpl) GetOriginalUrlByCode(code string) string {
	// get existing data from db
	urlsHash, err := u.db.HGetAll(dbCtx, "Urls").Result()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// getting original url by code (short url), provided in user request
	originalUrl := getOriginalUrl(urlsHash, code)

	return originalUrl
}

func getOriginalUrl(urlsHash map[string]string, shortUrl string) string {
	for k, v := range urlsHash {
		if k == shortUrl {
			return v
		}
	}
	return ""
}

func getShortUrl(urlsHash map[string]string, originalUrl string) string {
	for k, v := range urlsHash {
		if v == originalUrl {
			return k
		}
	}
	return ""
}
