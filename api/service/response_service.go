package service

import "github.com/a-shdv/url-shortener/api/repo"

type ResponseService interface {
	GetOriginalUrl(string)
}

type ResponseServiceImpl struct {
	repo repo.ResponseRepo
}

func NewResponseService(repo repo.ResponseRepo) *ResponseServiceImpl {
	return &ResponseServiceImpl{
		repo: repo,
	}
}

func (r *ResponseServiceImpl) GetOriginalUrl(url string) {
	//TODO implement me
	panic("implement me")
}
