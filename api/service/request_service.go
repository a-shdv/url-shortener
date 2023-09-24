package service

import "github.com/a-shdv/url-shortener/api/repo"

type RequestService struct {
	repo repo.RequestRepo
}

func NewRequestService(repo repo.RequestRepo) *RequestService {
	return &RequestService{repo: repo}
}

func (r RequestService) CreateShortUrl() {
	//TODO implement me
	panic("implement me")
}
