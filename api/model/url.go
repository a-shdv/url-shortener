package model

import "time"

type Url struct {
	OriginalUrl         string        `json:"originalUrl"`
	CustomShortUrl      string        `json:"customShortUrl"`
	ExpirationTimeHours time.Duration `json:"expirationTimeHours"`
}
