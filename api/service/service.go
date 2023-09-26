package service

import (
	"github.com/a-shdv/url-shortener/api/model"
	"github.com/a-shdv/url-shortener/api/repo"
)

// Service структура.
type Service struct {
	UrlService
}

// NewService конструктор
func NewService(repo *repo.Repo) *Service {
	return &Service{
		UrlService: NewUrlService(repo.UrlRepo),
	}
}

// Url интерфейс
type Url interface {
	CreateShortUrl(request model.Url)
}
