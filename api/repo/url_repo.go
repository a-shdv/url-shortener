package repo

import (
	"github.com/a-shdv/url-shortener/api/helper"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
)

// UrlRepo интерфейс.
type UrlRepo interface {
	CreateShortUrl(shortUrl, originalUrl string) (string, error)
	GetOriginalUrlByCode(string) string
}

// UrlRepoImpl структура.
type UrlRepoImpl struct {
	db *redis.Client
}

// NewUrlRepoImpl конструктор.
func NewUrlRepoImpl(db *redis.Client) *UrlRepoImpl {
	return &UrlRepoImpl{
		db: db,
	}
}

// CreateShortUrl метод, отвечающий за создание короткого url-адреса и сохранение в БД.
func (u *UrlRepoImpl) CreateShortUrl(shortUrl, originalUrl string) (string, error) {
	// получить уже существующие данные из бд.
	urlsHash, err := u.db.HGetAll(dbCtx, "Urls").Result()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// проверить, существует ли укороченная версия url-адреса в бд.
	shortUrlDb := getShortUrl(urlsHash, originalUrl)
	if shortUrlDb != "" {
		return shortUrlDb, nil
	}

	// закодировать короткий url-адрес символами /^[A-z0-9]{8}$/
	if shortUrl == "" {
		shortUrl = helper.GenerateRandomChar()
	}

	// добавить пару ключ-значение в хэш-таблицу'Urls'.
	wg := &sync.WaitGroup{}
	go func() {
		wg.Add(1)
		err = u.db.HSet(dbCtx, "Urls", shortUrl, originalUrl).Err()
		wg.Done()
	}()
	wg.Wait()

	if err != nil {
		return "", err
	}

	return shortUrl, nil
}

// GetOriginalUrlByCode метод, отвечающий за получение исходного url-адреса по его укороченной версии (коду).
func (u *UrlRepoImpl) GetOriginalUrlByCode(code string) string {
	// получить уже существующие данные из бд.
	urlsHash, err := u.db.HGetAll(dbCtx, "Urls").Result()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// получение исходного url-адреса по его укороченной версии (коду).
	originalUrl := getOriginalUrl(urlsHash, code)

	return originalUrl
}

// getOriginalUrl метод-хелпер для получения исходного url-адреса.
func getOriginalUrl(urlsHash map[string]string, shortUrl string) string {
	for k, v := range urlsHash {
		if k == shortUrl {
			return v
		}
	}

	return ""
}

// getShortUrl метод-хелпер для получения укороченной версии url-адреса.
func getShortUrl(urlsHash map[string]string, originalUrl string) string {
	for k, v := range urlsHash {
		if v == originalUrl {
			return k
		}
	}
	return ""
}
