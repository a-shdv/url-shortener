package helper

import (
	"os"
	"strings"
)

// TODO
func parseUrlDomain(url string) string {
	var domain string

	domain = strings.Replace(url, "http://", "", 1)  // remove 'http://' from url-address
	domain = strings.Replace(url, "https://", "", 1) // remove 'https://' from url-address
	domain = strings.Replace(url, "www.", "", 1)     // remove 'www' from url-address
	domain = strings.Split(url, "/")[0]

	return domain
}

// IsReqUrlServerAddr TODO
func IsReqUrlServerAddr(reqUrl string) bool {
	reqUrl = parseUrlDomain(reqUrl)
	serverAddr := os.Getenv("SERVER_ADDR")

	if reqUrl == serverAddr {
		return true
	}
	return false
}

// ReplaceHttpsWithHttp TODO
func ReplaceHttpsWithHttp(url string) string {
	// make every url https
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}
