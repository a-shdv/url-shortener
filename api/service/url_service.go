package service

import (
	"github.com/a-shdv/url-shortener/api/model"
	"github.com/a-shdv/url-shortener/api/repo"
)

type UrlService interface {
	CreateShortUrl(*model.Url) (string, error)
	GetOriginalUrlByCode(string) string
}

type UrlServiceImpl struct {
	repo repo.UrlRepo
}

func NewUrlService(repo repo.UrlRepo) *UrlServiceImpl {
	return &UrlServiceImpl{repo: repo}
}

func (u *UrlServiceImpl) CreateShortUrl(req *model.Url) (string, error) {
	var shortUrl string

	// short url has already been initialized during request
	if req.CustomShortUrl != "" {
		shortUrl = req.CustomShortUrl[:8] // accept only 8 characters
	}

	// creating new short url
	res, err := u.repo.CreateShortUrl(shortUrl, req.OriginalUrl)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (u *UrlServiceImpl) GetOriginalUrlByCode(code string) string {
	// getting existing origin url from db
	url := u.repo.GetOriginalUrlByCode(code)
	return url
}
