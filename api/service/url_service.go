package service

import (
	"github.com/a-shdv/url-shortener/api/model"
	"github.com/a-shdv/url-shortener/api/repo"
)

// UrlService интерфейс.
type UrlService interface {
	CreateShortUrl(*model.Url) (string, error)
	GetOriginalUrlByCode(string) string
}

// UrlServiceImpl структура.
type UrlServiceImpl struct {
	repo repo.UrlRepo
}

// NewUrlService конструктор.
func NewUrlService(repo repo.UrlRepo) *UrlServiceImpl {
	return &UrlServiceImpl{repo: repo}
}

// CreateShortUrl сервис, отвечающий за создание укороченной версии url-адреса.
func (u *UrlServiceImpl) CreateShortUrl(req *model.Url) (string, error) {
	var shortUrl string

	// проверка на то, было ли указано желаемое значение укороченной версии url-адреса пользователем
	// во время отправки запроса
	if req.CustomShortUrl != "" {
		shortUrl = req.CustomShortUrl[:8] // ограничение на 8 символов
	}

	// создание новой укороченной версии url-адреса
	res, err := u.repo.CreateShortUrl(shortUrl, req.OriginalUrl)
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetOriginalUrlByCode получение исходного url-адреса по его укороченной версии.
func (u *UrlServiceImpl) GetOriginalUrlByCode(code string) string {
	url := u.repo.GetOriginalUrlByCode(code)
	return url
}
