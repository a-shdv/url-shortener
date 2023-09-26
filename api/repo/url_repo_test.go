package repo

import (
	"github.com/a-shdv/url-shortener/api/helper"
	"github.com/go-redis/redis/v8"
	"os"
	"sync"
	"testing"
	"time"
)

// TestCreateShortUrl тестирование создания записей в хэш-таблицы redis.
func TestCreateShortUrl(t *testing.T) {
	// создание клиента redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB:       0,
	})
	domain := "example.com/"
	expected := 10 * time.Second // ожидаемый результат
	wg := &sync.WaitGroup{}

	batchSize := 100 // размер пакета
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		// ожидание выполнения всех горутин текущей пакетной операции
		// обращение к БД идёт в виде пакетов, чтобы уменить количество запросов к БД
		if i%batchSize == 0 {
			wg.Wait()
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			k := helper.GenerateRandomChar()
			v := domain + helper.GenerateRandomChar()
			rdb.HSet(dbCtx, "UrlsTest2", k, v)
		}()
	}
	end := time.Since(start)

	if end >= expected {
		t.Errorf("Creation operation takes too long. Expected %d, got %d", expected, end)
	}
	t.Logf("Creation operation time: %d", end)
}

// TestGetOriginalUrl тестирование операции получения иходного url-адреса.
func TestGetOriginalUrl(t *testing.T) {
	// создание клиента redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB:       0,
	})
	expected := 5 * time.Second

	// использование метода HGetAll, который оптимизирован для получения данных из redis
	urlsHash, _ := rdb.HGetAll(dbCtx, "UrlsTest2").Result()

	var test string
	start := time.Now()
	go func() {
		test = getOriginalUrl(urlsHash, "XJcDZz4dTIeZ40WRPj3f")
	}()
	end := time.Since(start)

	if end >= expected {
		t.Errorf("Reading operation takes too long. Expected %d, got %d", expected, end)
	}
	t.Logf("Reading operation time: %d", end)
	t.Logf(test)
}

// TestGetShortUrl тестирование операции получения укороченной версии url-адреса.
func TestGetShortUrl(t *testing.T) {
	// создание клиента redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB:       0,
	})
	expected := 5 * time.Second

	// использование метода HGetAll, который оптимизирован для получения данных из redis
	urlsHash, _ := rdb.HGetAll(dbCtx, "UrlsTest2").Result()

	var test string
	start := time.Now()
	go func() {
		test = getShortUrl(urlsHash, "example.com/B9rnAoLg3dwFuj1cMufT")
	}()
	end := time.Since(start)

	if end >= expected {
		t.Errorf("Reading operation takes too long. Expected %d, got %d", expected, end)
	}
	t.Logf("Reading operation time: %d", end)
	t.Logf("short url: %s", test)

}
