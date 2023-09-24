package helpers

import (
	"os"
	"strings"
)

// EnforceHTTP ...
func EnforceHTTP(url string) string {
	// make every url https
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}

// RemoveDomainError ...
func RemoveDomainError(url string) bool {
	// basically this functions removes all the commonly found
	// prefixes from URL such as http, https, www
	// then checks of the remaining string is the DOMAIN itself
	if url == os.Getenv("URL_ADDR") {
		return false
	}

	newURL := strings.Replace(url, "http://", "", 1)    // remove 'http' from url-address
	newURL = strings.Replace(newURL, "https://", "", 1) // remove 'https' from url-address
	newURL = strings.Replace(newURL, "www.", "", 1)     // remove 'www' from url-addresss

	newURL = strings.Split(newURL, "/")[0] // remove everything that comes after '/'

	if newURL == os.Getenv("URL_ADDR") { // ensure that url from request is not the same as server url
		return false
	}
	return true
}
