package repo

import (
	"github.com/a-shdv/url-shortener/api/model"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"log"
	"time"
)

type RequestRepo struct {
	db *redis.Client
}

func NewRequestRepo(db *redis.Client) *RequestRepo {
	return &RequestRepo{
		db: db,
	}
}

func (r *RequestRepo) CreateShortUrl(req model.Request) {
	var shortUrl string
	if req.CustomShortUrl == "" {
		shortUrl = uuid.New().String()[:6]
	} else {
		shortUrl = req.CustomShortUrl
	}

	err := r.db.Set(dbCtx, shortUrl, req.OriginalUrl, req.ExpirationTimeHours*3600*time.Second).Err()
	if err != nil {
		log.Fatal(err.Error())
	}
}
