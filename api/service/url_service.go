package service

import (
	"github.com/a-shdv/url-shortener/api/helper"
	"github.com/a-shdv/url-shortener/api/model"
	"github.com/a-shdv/url-shortener/api/repo"
)

type UrlService interface {
	CreateShortUrl(*model.Url) string
}

type UrlServiceImpl struct {
	repo repo.UrlRepo
}

func NewUrlService(repo repo.UrlRepo) *UrlServiceImpl {
	return &UrlServiceImpl{repo: repo}
}

func (r *UrlServiceImpl) CreateShortUrl(req *model.Url) string {
	// generating a new short url address
	var shortUrl string
	if req.CustomShortUrl == "" {
		shortUrl = helper.GenerateRandomChar()
	} else {
		shortUrl = req.CustomShortUrl
	}

	r.repo.CreateShortUrl(shortUrl, req.OriginalUrl, req.ExpirationTimeHours)

	return shortUrl
}
