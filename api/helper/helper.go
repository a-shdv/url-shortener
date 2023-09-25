package helper

import (
	"log"
	"math/rand"
	"os"
	"strings"
)

func ParseUrlAddr(url string) string {
	var domain string

	domain = strings.Replace(url, "http://", "", 1)     // remove 'http://' from url-address
	domain = strings.Replace(domain, "https://", "", 1) // remove 'https://' from url-address
	domain = strings.Replace(domain, "www.", "", 1)     // remove 'www' from url-address
	domain = strings.Split(domain, "/")[0]

	err := isReqUrlServerAddr(domain)
	if err {
		log.Fatalf("forbidden to use this url address!")
	}

	return domain
}

func isReqUrlServerAddr(reqUrl string) bool {
	serverAddr := os.Getenv("SERVER_ADDR")

	if reqUrl == serverAddr {
		return true
	}
	return false
}

func GenerateRandomChar() string {
	charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	res := ""
	for i := 0; i < 8; i++ {
		res += string(charSet[rand.Intn(len(charSet))])
	}
	return res
}
