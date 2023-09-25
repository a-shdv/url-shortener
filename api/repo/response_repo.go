package repo

import "github.com/go-redis/redis/v8"

type ResponseRepo interface {
}

type ResponseRepoImpl struct {
	db *redis.Client
}

func NewResponseRepoImpl(db *redis.Client) *ResponseRepoImpl {
	return &ResponseRepoImpl{
		db: db,
	}
}
