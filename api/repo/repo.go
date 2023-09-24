package repo

import "github.com/go-redis/redis/v8"

type Repo struct {
	*RequestRepo
	*ResponseRepo
}

func NewRepository(db *redis.Client) *Repo {
	return &Repo{
		RequestRepo:  NewRequestRepo(db),
		ResponseRepo: NewResponseRepo(db),
	}
}