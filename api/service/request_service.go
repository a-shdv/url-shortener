package service

import (
	"github.com/a-shdv/url-shortener/api/model"
	"github.com/a-shdv/url-shortener/api/repo"
)

type RequestService struct {
	repo repo.RequestRepo
}

func NewRequestService(repo repo.RequestRepo) *RequestService {
	return &RequestService{repo: repo}
}

func (r *RequestService) CreateShortUrl(req model.Request) {
	r.repo.CreateShortUrl(req)
}
