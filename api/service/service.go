package service

import (
	"github.com/a-shdv/url-shortener/api/model"
	"github.com/a-shdv/url-shortener/api/repo"
)

type Service struct {
	UrlService
}

func NewService(repo *repo.Repo) *Service {
	return &Service{
		UrlService: NewUrlService(repo.UrlRepo),
	}
}

type Url interface {
	CreateShortUrl(request model.Url)
}
