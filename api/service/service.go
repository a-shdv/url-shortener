package service

import (
	"github.com/a-shdv/url-shortener/api/model"
	"github.com/a-shdv/url-shortener/api/repo"
)

type Service struct {
	RequestService
	ResponseService
}

func NewService(repo *repo.Repo) *Service {
	return &Service{
		RequestService:  NewRequestService(repo.RequestRepo),
		ResponseService: NewResponseService(repo.ResponseRepo),
	}
}

type Request interface {
	CreateShortUrl(request model.Request)
}

type Response interface {
	GetOriginalUrl(url string)
}
