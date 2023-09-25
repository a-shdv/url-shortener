package repo

import "github.com/go-redis/redis/v8"

// Repo структура.
type Repo struct {
	UrlRepo
}

// NewRepository конструктор.
func NewRepository(db *redis.Client) *Repo {
	return &Repo{
		UrlRepo: NewUrlRepoImpl(db),
	}
}
