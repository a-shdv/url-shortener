package repo

import "github.com/go-redis/redis/v8"

type Repo struct {
	UrlRepo
}

func NewRepository(db *redis.Client) *Repo {
	return &Repo{
		UrlRepo: NewUrlRepoImpl(db),
	}
}
