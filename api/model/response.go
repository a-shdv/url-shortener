package model

import "time"

type Response struct {
	OriginalUrl         string        `json:"originalUrl"`
	ShortUrl            string        `json:"shortUrl"`
	ExpirationTimeHours time.Duration `json:"expirationTimeHours"`
}
