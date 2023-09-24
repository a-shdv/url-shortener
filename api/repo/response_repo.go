package repo

import "github.com/go-redis/redis/v8"

type ResponseRepo struct {
	db *redis.Client
}

func NewResponseRepo(db *redis.Client) *ResponseRepo {
	return &ResponseRepo{
		db: db,
	}
}
