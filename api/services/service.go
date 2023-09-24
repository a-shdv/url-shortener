package services

import "github.com/a-shdv/url-shortener/api/repositories"

type Service struct {
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{}
}
