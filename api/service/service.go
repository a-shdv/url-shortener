package service

import (
	"github.com/a-shdv/url-shortener/api/model"
	"github.com/a-shdv/url-shortener/api/repo"
)

type Service struct {
	Request
	Response
}

func NewService(repo *repo.Repo) *Service {
	return &Service{
		Request:  NewRequestService(*repo.RequestRepo),
		Response: NewResponseService(*repo.ResponseRepo),
	}
}

type Request interface {
	CreateShortUrl(request model.Request)
}

type Response interface {
	GetOriginalUrl(url string)
}
