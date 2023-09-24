package repo

import "github.com/go-redis/redis/v8"

type RequestRepo struct {
	db *redis.Client
}

func NewRequestRepo(db *redis.Client) *RequestRepo {
	return &RequestRepo{
		db: db,
	}
}
