package helper

import (
	"log"
	"math/rand"
	"os"
	"strings"
)

// ParseUrlAddr парсинг url-адреса, чтобы привести его к более общему виду, например, "example.com/something".
func ParseUrlAddr(url string) string {
	var res string

	res = strings.Replace(url, "http://", "", 1)  // убрать 'http://' из url-адреса.
	res = strings.Replace(res, "https://", "", 1) // убрать 'https://' из url-адреса.
	res = strings.Replace(res, "www.", "", 1)     // убрать 'www' из url-адреса.

	// проверка на то, вдруг пользователь ввёл тот же адрес, что и у сервера.
	err := isReqUrlServerAddr(res)
	if err {
		log.Fatalf("forbidden to use this url address!")
	}

	return res
}

// isReqUrlServerAddr проверка на то, вдруг пользователь ввёл тот же адрес, что и у сервера.
func isReqUrlServerAddr(reqUrl string) bool {
	serverAddr := os.Getenv("SERVER_ADDR")

	if reqUrl == serverAddr {
		return true
	}
	return false
}

// GenerateRandomChar генерирует символы из множества: /^[A-z0-9]{8}$/
func GenerateRandomChar() string {
	charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	res := ""
	for i := 0; i < 8; i++ {
		res += string(charSet[rand.Intn(len(charSet))])
	}
	return res
}
