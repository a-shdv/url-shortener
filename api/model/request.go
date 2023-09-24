package model

import "time"

type Request struct {
	OriginalUrl         string        `json:"originalUrl"`
	CustomShortUrl      string        `json:"customShortUrl"`
	ExpirationTimeHours time.Duration `json:"expirationTimeHours"`
}
