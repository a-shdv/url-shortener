package repo

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

var dbCtx = context.Background()

func NewRedisDb(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB:       dbNo,
	})
	return rdb
}
