package service

import (
	"github.com/a-shdv/url-shortener/api/model"
	"github.com/a-shdv/url-shortener/api/repo"
	"github.com/google/uuid"
)

type RequestService interface {
	CreateShortUrl(*model.Request)
}

type RequestServiceImpl struct {
	repo repo.RequestRepo
}

func NewRequestService(repo repo.RequestRepo) *RequestServiceImpl {
	return &RequestServiceImpl{repo: repo}
}

func (r *RequestServiceImpl) CreateShortUrl(req *model.Request) {
	// generating a new short url address
	var shortUrl string
	if req.CustomShortUrl == "" {
		shortUrl = uuid.New().String()[:6]
	} else {
		shortUrl = req.CustomShortUrl
	}

	r.repo.CreateShortUrl(shortUrl, req.OriginalUrl, req.ExpirationTimeHours)
}
