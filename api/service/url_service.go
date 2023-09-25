package service

import (
	"github.com/a-shdv/url-shortener/api/helper"
	"github.com/a-shdv/url-shortener/api/model"
	"github.com/a-shdv/url-shortener/api/repo"
	"time"
)

type UrlService interface {
	CreateShortUrl(*model.Url) (string, error)
	//GetOriginalUrl(string) (string, error)
}

type UrlServiceImpl struct {
	repo repo.UrlRepo
}

func NewUrlService(repo repo.UrlRepo) *UrlServiceImpl {
	return &UrlServiceImpl{repo: repo}
}

func (u *UrlServiceImpl) CreateShortUrl(req *model.Url) (string, error) {
	var shortUrl string
	if req.CustomShortUrl != "" {
		shortUrl = req.CustomShortUrl
	} else {
		shortUrl = helper.GenerateRandomChar()
	}

	res, err := u.repo.CreateShortUrl(req.OriginalUrl, shortUrl, 12*time.Hour)
	if err != nil {
		return res, err
	}

	return res, nil
}

//func (u *UrlServiceImpl) GetOriginalUrl(reqUrl string) (string, error) {
//	urlDb, err := u.repo.GetOriginalUrl(reqUrl)
//	if err != nil {
//		return "", err
//	}
//	return urlDb, nil
//}
