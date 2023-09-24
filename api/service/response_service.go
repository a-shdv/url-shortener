package service

import "github.com/a-shdv/url-shortener/api/repo"

type ResponseService struct {
	repo repo.ResponseRepo
}

func NewResponseService(repo repo.ResponseRepo) *ResponseService {
	return &ResponseService{
		repo: repo,
	}
}

func (r *ResponseService) GetOriginalUrl(url string) {
	//TODO implement me
	panic("implement me")
}
